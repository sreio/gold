package model

type Chanel string

const (
	PUSHDEER    = "push_deer"
	SERVERJIANG = "server_jiang"
)

type ChanelList []Chanel

var ChanelDataList = ChanelList{PUSHDEER, SERVERJIANG}
