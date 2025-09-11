package notifications

import (
	"fmt"
	"github.com/sreio/gold/tools"
	"net/url"
	"sync"
)

type PushDeer struct {
	Url        string
	DeviceKeys []string
	Text       string
}

func (p *PushDeer) SetMsgData(msgData MsgData) MsgSendChannelInterface {
	return &PushDeer{
		Url:        p.Url,
		DeviceKeys: msgData.PushDeer.DeviceKeys,
		Text:       msgData.PushDeer.Text,
	}
}

func (p *PushDeer) SendMessage() bool {
	if len(p.DeviceKeys) == 0 {
		return true
	}
	var wg sync.WaitGroup
	errCh := make(chan error, len(p.DeviceKeys))
	escText := url.QueryEscape(p.Text)

	for _, key := range p.DeviceKeys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			u := fmt.Sprintf(p.Url, k, escText)
			_, code, err := tools.HTTPRequest("GET", u, nil, nil)
			if err != nil {
				errCh <- fmt.Errorf("key %s code:%d err:%v", k, code, err)
			}
		}(key)
	}
	wg.Wait()
	close(errCh)
	return len(errCh) == 0
}
