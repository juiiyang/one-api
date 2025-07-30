package monitor

import (
	"fmt"

	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/common/message"
	"github.com/songquanpeng/one-api/model"
)

func notifyRootUser(subject string, content string) {
	if config.MessagePusherAddress != "" {
		err := message.SendMessage(subject, content, content)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("failed to send message: %s", err.Error()))
		} else {
			return
		}
	}
	if config.RootUserEmail == "" {
		config.RootUserEmail = model.GetRootUserEmail()
	}
	err := message.SendEmail(subject, config.RootUserEmail, content)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("failed to send email: %s", err.Error()))
	}
}

// DisableChannel disable & notify
func DisableChannel(channelId int, channelName string, reason string) {
	model.UpdateChannelStatusById(channelId, model.ChannelStatusAutoDisabled)
	logger.Logger.Info(fmt.Sprintf("channel #%d has been disabled: %s", channelId, reason))
	subject := fmt.Sprintf("Channel Status Change Reminder")
	content := message.EmailTemplate(
		subject,
		fmt.Sprintf(`
            <p>Hello!</p>
            <p>Channel “<strong>%s</strong>” (#%d) has been disabled.</p>
            <p>Reason for disabling:</p>
            <p style="background-color: #f8f8f8; padding: 10px; border-radius: 4px;">%s</p>
        `, channelName, channelId, reason),
	)
	notifyRootUser(subject, content)
}

func MetricDisableChannel(channelId int, successRate float64) {
	model.UpdateChannelStatusById(channelId, model.ChannelStatusAutoDisabled)
	logger.Logger.Info(fmt.Sprintf("channel #%d has been disabled due to low success rate: %.2f", channelId, successRate*100))
	subject := fmt.Sprintf("Channel Status Change Reminder")
	content := message.EmailTemplate(
		subject,
		fmt.Sprintf(`
            <p>Hello!</p>
            <p>Channel #%d has been automatically disabled by the system.</p>
            <p>Reason for disabling:</p>
            <p style="background-color: #f8f8f8; padding: 10px; border-radius: 4px;">In the last %d calls, the success rate of this channel was <strong>%.2f%%</strong>, which is below the system threshold of <strong>%.2f%%</strong>.</p>
        `, channelId, config.MetricQueueSize, successRate*100, config.MetricSuccessRateThreshold*100),
	)
	notifyRootUser(subject, content)
}

// EnableChannel enable & notify
func EnableChannel(channelId int, channelName string) {
	model.UpdateChannelStatusById(channelId, model.ChannelStatusEnabled)
	logger.Logger.Info(fmt.Sprintf("channel #%d has been enabled", channelId))
	subject := fmt.Sprintf("Channel Status Change Reminder")
	content := message.EmailTemplate(
		subject,
		fmt.Sprintf(`
            <p>Hello!</p>
            <p>Channel “<strong>%s</strong>” (#%d) has been re-enabled.</p>
            <p>You can now continue using this channel.</p>
        `, channelName, channelId),
	)
	notifyRootUser(subject, content)
}
