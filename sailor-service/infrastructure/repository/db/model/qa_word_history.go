package model

import (
	"time"
)

const TableNameQaWordHistory = "t_qa_word_history"

type QaWordHistory struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	UserId string `gorm:"column:user_id;not null;comment:用户id" json:"user_id"` // 资源名称

	QWordList *string    `gorm:"column:qword_list;comment:详细信息" json:"qword_list"`                                                     //详细信息
	CreatedAt *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
}
