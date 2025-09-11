package notifications

import (
	"fmt"
	"github.com/sreio/gold/tools"
	"net/url"
	"sync"
)

type Uid string
type SendKey string

type ServerJiang struct {
	Url         string
	UserSendMap map[Uid]SendKey
	Title       string
	Desp        string
	Tags        string
	Short       string
}

func (sj *ServerJiang) SetMsgData(msgData MsgData) MsgSendChannelInterface {
	return &ServerJiang{
		Url:         sj.Url,
		UserSendMap: msgData.ServerJiang.UserSendMap,
		Title:       msgData.ServerJiang.Title,
		Desp:        msgData.ServerJiang.Desp,
		Tags:        msgData.ServerJiang.Tags,
		Short:       msgData.ServerJiang.Short,
	}
}

func (sj *ServerJiang) SendMessage() bool {
	if len(sj.UserSendMap) == 0 {
		return true
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	postData := fmt.Sprintf("title=%s&desp=%s&tags=%s&short=%s",
		url.QueryEscape(sj.Title),
		url.QueryEscape(sj.Desp),
		url.QueryEscape(sj.Tags),
		url.QueryEscape(sj.Short),
	)

	var wg sync.WaitGroup
	errCh := make(chan error, len(sj.UserSendMap))

	for uid, sendKey := range sj.UserSendMap {
		uid, sendKey := uid, sendKey
		wg.Add(1)
		go func() {
			defer wg.Done()
			u := fmt.Sprintf(sj.Url, string(uid), string(sendKey))
			_, code, err := tools.HTTPRequest("POST", u, headers, []byte(postData))
			if err != nil {
				errCh <- fmt.Errorf("uid %s code:%d err:%v", uid, code, err)
			}
		}()
	}
	wg.Wait()
	close(errCh)
	return len(errCh) == 0
}
