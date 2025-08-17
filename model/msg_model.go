package model

import (
	helper "go-micro/utils"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	MsgType   string    `json:"msg_type"`
	MsgBody   string    `json:"message_body"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type MessageRepo interface {
	Insert(model *Message) error
	Gets(model *Message) error
	Delete(model *Message) error
}

type messageRepo struct {
	db *gorm.DB
}

func NewMsgRepository(db *gorm.DB) MessageRepo {
	return &messageRepo{db}
}

func (d *messageRepo) Insert(model *Message) error {
	if model.ID == "" {
		model.ID = helper.NewUUID()
	}
	return d.db.Save(model).Error
}

func (d *messageRepo) Gets(model *Message) error {
	return d.db.Find(&model).Error
}

func (d *messageRepo) Delete(model *Message) error {
	return d.db.Delete(&model).Error
}
