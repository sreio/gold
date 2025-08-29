package model

type Source struct {
	ID
	Type string `gorm:"column:type;type:string;size:256;not null;unique" json:"type"`
	Name string `gorm:"column:name;type:string;size:256;not null" json:"name"`
	Time
}
