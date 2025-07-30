package openai

import (
	"github.com/Laisky/zap"

	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/relay/adaptor/openai_compatible"
	"github.com/songquanpeng/one-api/relay/model"
)

func ErrorWrapper(err error, code string, statusCode int) *model.ErrorWithStatusCode {
	logger.Logger.Error("API error",
		zap.String("code", code),
		zap.Error(err))

	Error := model.Error{
		Message: err.Error(),
		Type:    "one_api_error",
		Code:    code,
	}
	return &model.ErrorWithStatusCode{
		Error:      Error,
		StatusCode: statusCode,
	}
}

// NormalizeDataLine normalizes SSE data lines
// This function delegates to the shared implementation for consistency
func NormalizeDataLine(data string) string {
	return openai_compatible.NormalizeDataLine(data)
}
