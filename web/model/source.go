package model

type Source string

const (
	IcbcOne Source = "icbc_1"
	IcbcTwo Source = "icbc_2"
)

type SourceList []Source

var SourcesList = SourceList{IcbcOne, IcbcTwo}
