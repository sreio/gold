package model

type Gold struct {
	ComID
	SourceType string  `gorm:"column:source_type;type:string;size:256;not null;comment:来源类型;index" json:"source_type"`
	Price      float64 `gorm:"column:price;type:float;not null;comment:每克金价" json:"price"`
	OtherData  string  `gorm:"column:other_data;type:string;comment:其他数据json格式" json:"other_data"`
	ComTime
}
