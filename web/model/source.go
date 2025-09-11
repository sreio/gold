package model

type Source string

const (
	ICBCSHOPPRICE     = "icbc_shop_price"
	ICBCBANKRUYIPRICE = "icbc_bank_ruyi_price"
)

type SourceList []Source

var SourcesList = SourceList{ICBCSHOPPRICE, ICBCBANKRUYIPRICE}
