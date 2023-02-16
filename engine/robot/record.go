package robot

import (
	"fmt"
	"time"

	"github.com/yqchilde/pkgs/utils"

	"github.com/yqchilde/wxbot/engine/pkg/sqlite"
)

// MessageRecord 消息记录表
type MessageRecord struct {
	ID         uint      `gorm:"primarykey"`         // 主键
	Type       string    `gorm:"column:type"`        // 消息类型
	FromWxId   string    `gorm:"column:from_wxid"`   // 消息来源wxid，群消息为群wxid，私聊消息为发送者wxid
	FromNick   string    `gorm:"column:from_nick"`   // 消息来源昵称，群消息为群昵称，私聊消息为发送者昵称
	SenderWxId string    `gorm:"column:sender_wxid"` // 消息具体发送者wxid
	SenderNick string    `gorm:"column:sender_nick"` // 消息具体发送者昵称
	Content    string    `gorm:"column:content"`     // 消息内容
	CreatedAt  time.Time `gorm:"column:created_at"`  // 创建时间
}

var db sqlite.DB

func initMessageRecordDB() error {
	if db.Orm != nil {
		return nil
	}

	dbPath := "data/manager/plugins.db"
	if !utils.CheckPathExists(dbPath) {
		return fmt.Errorf("db file not found: %s", dbPath)
	}
	if err := sqlite.Open(dbPath, &db); err != nil {
		return err
	}
	return nil
}

// GetHistoryByWxId 根据wxId获取消息记录
func (ctx *Ctx) GetHistoryByWxId(wxId string) ([]MessageRecord, error) {
	if err := initMessageRecordDB(); err != nil {
		return nil, err
	}

	var msgRecord []MessageRecord
	if err := db.Orm.Table("__message").Where("from_wxid = ?", wxId).Find(&msgRecord).Error; err != nil {
		return nil, err
	}
	return msgRecord, nil
}

// GetHistoryByWxIdAndDate 根据wxId和日期获取消息记录
func (ctx *Ctx) GetHistoryByWxIdAndDate(wxId, date string) ([]MessageRecord, error) {
	if err := initMessageRecordDB(); err != nil {
		return nil, err
	}

	var msgRecord []MessageRecord
	if err := db.Orm.Table("__message").Where("from_wxid = ? AND STRFTIME('%Y-%m-%d', created_at, 'localtime') = ?", wxId, date).Find(&msgRecord).Error; err != nil {
		return nil, err
	}
	return msgRecord, nil
}