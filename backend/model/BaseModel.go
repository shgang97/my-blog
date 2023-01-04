package model

import "time"

/*
@author: shg
@since: 2022/11/30
@desc: //TODO
*/

type BaseModel struct {
	Pk       uint      `gorm:"primaryKey" json:"pk"`
	Id       string    `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
