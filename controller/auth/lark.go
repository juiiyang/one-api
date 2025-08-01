package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Laisky/errors/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/controller"
	"github.com/songquanpeng/one-api/model"
)

type LarkOAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type LarkUser struct {
	Name   string `json:"name"`
	OpenID string `json:"open_id"`
	Email  string `json:"email"`
}

func getLarkUserInfoByCode(code string) (*LarkUser, error) {
	if code == "" {
		return nil, errors.New("Invalid parameter")
	}
	values := map[string]string{
		"client_id":     config.LarkClientId,
		"client_secret": config.LarkClientSecret,
		"code":          code,
		"grant_type":    "authorization_code",
		"redirect_uri":  fmt.Sprintf("%s/oauth/lark", config.ServerAddress),
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/authen/v2/oauth/token", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		logger.Logger.Info(err.Error())
		return nil, errors.New("Unable to connect to Lark server, please try again later!")
	}
	defer res.Body.Close()
	var oAuthResponse LarkOAuthResponse
	err = json.NewDecoder(res.Body).Decode(&oAuthResponse)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest("GET", "https://passport.feishu.cn/suite/passport/oauth/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oAuthResponse.AccessToken))
	res2, err := client.Do(req)
	if err != nil {
		logger.Logger.Info(err.Error())
		return nil, errors.New("Unable to connect to Lark server, please try again later!")
	}
	var larkUser LarkUser
	err = json.NewDecoder(res2.Body).Decode(&larkUser)
	if err != nil {
		return nil, err
	}
	return &larkUser, nil
}

func LarkOAuth(c *gin.Context) {
	ctx := c.Request.Context()
	session := sessions.Default(c)
	state := c.Query("state")
	if state == "" || session.Get("oauth_state") == nil || state != session.Get("oauth_state").(string) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "state is empty or not same",
		})
		return
	}
	username := session.Get("username")
	if username != nil {
		LarkBind(c)
		return
	}
	code := c.Query("code")
	larkUser, err := getLarkUserInfoByCode(code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user := model.User{
		LarkId: larkUser.OpenID,
	}
	if model.IsLarkIdAlreadyTaken(user.LarkId) {
		err := user.FillUserByLarkId()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	} else {
		if config.RegisterEnabled {
			parts := strings.Split(larkUser.Email, "@")
			if len(parts) > 1 {
				user.Username = parts[0]
			} else {
				user.Username = "lark_" + strconv.Itoa(model.GetMaxUserId()+1)
			}
			if larkUser.Name != "" {
				user.DisplayName = larkUser.Name
			} else {
				user.DisplayName = "Lark User"
			}
			user.Role = model.RoleCommonUser
			user.Status = model.UserStatusEnabled

			if err := user.Insert(ctx, 0); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "The administrator has turned off new user registration",
			})
			return
		}
	}

	if user.Status != model.UserStatusEnabled {
		c.JSON(http.StatusOK, gin.H{
			"message": "User has been banned",
			"success": false,
		})
		return
	}
	controller.SetupLogin(&user, c)
}

func LarkBind(c *gin.Context) {
	code := c.Query("code")
	larkUser, err := getLarkUserInfoByCode(code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user := model.User{
		LarkId: larkUser.OpenID,
	}
	if model.IsLarkIdAlreadyTaken(user.LarkId) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "This Lark account has already been bound",
		})
		return
	}
	session := sessions.Default(c)
	id := session.Get("id")
	// id := c.GetInt("id")  // critical bug!
	user.Id = id.(int)
	err = user.FillUserById()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user.LarkId = larkUser.OpenID
	err = user.Update(false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "bind",
	})
	return
}
