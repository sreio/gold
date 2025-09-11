package notifications

import "github.com/sreio/gold/web/model"

type MsgData struct {
	PushDeer struct {
		DeviceKeys []string
		Text       string
	}

	ServerJiang struct {
		UserSendMap map[Uid]SendKey
		Title       string
		Desp        string
		Tags        string
		Short       string
	}
}

type MsgSendChannelInterface interface {
	SetMsgData(MsgData) MsgSendChannelInterface
	SendMessage() bool
}

var MsgMap = map[string]MsgSendChannelInterface{
	model.PUSHDEER: &PushDeer{
		Url: "https://api2.pushdeer.com/message/push?pushkey=%s&text=%s",
	},
	model.SERVERJIANG: &ServerJiang{
		Url: "https://%s.push.ft07.com/send/%s.send",
	},
}
