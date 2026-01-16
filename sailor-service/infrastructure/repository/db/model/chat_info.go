package model

import "time"

type ChatHistory struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	UserId    string `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`
	SessionId string `gorm:"column:session_id;not null;comment:用户id" json:"session_id"` // 资源名称

	Title      string     `gorm:"column:title;comment:详细信息" json:"title"`   //详细信息
	Status     string     `gorm:"column:status;comment:详细信息" json:"status"` //详细信息
	FavoriteId string     `gorm:"column:favorite_id;not null;comment:收藏" json:"favorite_id"`
	FavoriteAt *time.Time `gorm:"column:favorite_at;not null;comment:收藏时间" json:"favorite_at"`                                          // 创建时间
	ChatAt     *time.Time `gorm:"column:chat_at;not null;comment:对话时间" json:"chat_at"`                                                  // 创建时间
	CreatedAt  *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
}

type ChatHistoryDetail struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	SessionId string `gorm:"column:session_id;not null;comment:用户id" json:"session_id"` // 资源名称

	QaId             string     `gorm:"column:qa_id;comment:qaid" json:"qa_id"` //详细信息
	Query            string     `gorm:"column:query;comment:问句" json:"query"`   //详细信息
	Answer           string     `gorm:"column:answer;not null;comment:答案" json:"answer"`
	Status           string     `gorm:"column:status;not null;comment:状态" json:"status"`
	FavoriteId       string     `gorm:"column:favorite_id;not null;comment:收藏id" json:"favorite_id"`
	Like             string     `gorm:"column:like_status;not null;comment:点赞状态" json:"like_status"`
	ResourceRequired string     `gorm:"column:resource_required;not null;comment:需要资源" json:"resource_required"`
	CreatedAt        *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt        *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
}

type ChatFavorite struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	FavoriteId string `gorm:"column:favorite_id;not null;comment:用户id" json:"favorite_id"` // 资源名称

	Title      string     `gorm:"column:title;comment:详细信息" json:"title"`                                                               //详细信息
	FavoriteAt *time.Time `gorm:"column:favorite_at;not null;comment:收藏时间" json:"favorite_at"`                                          // 创建时间
	CreatedAt  *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
}

type ChatFavoriteDetail struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	SessionId  string     `gorm:"column:session_id;not null;comment:用户id" json:"session_id"` // 资源名称
	FavoriteId string     `gorm:"column:favorite_id;not null;comment:收藏id" json:"favorite_id"`
	QaId       string     `gorm:"column:qa_id;comment:详细信息" json:"qa_id"` //详细信息
	Query      string     `gorm:"column:query;comment:详细信息" json:"query"` //详细信息
	Answer     string     `gorm:"column:answer;not null;comment:收藏" json:"answer"`
	Status     string     `gorm:"column:status;not null;comment:收藏" json:"status"`
	Like       string     `gorm:"column:like_status;not null;comment:收藏" json:"like_status"`
	CreatedAt  *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;default:current_timestamp();comment:更新时间" json:"updated_at"` // 更新时间
}

type AssistantConfig struct {
	//ID     int    `gorm:"primary_key;type:int(11) auto_increment;not null;comment:'ID';" json:"id,omitempty"` // 逻辑主键
	UserId    string     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"` // 资源名称
	TType     string     `gorm:"column:type;not null;comment:配置类型" json:"type"`
	Config    string     `gorm:"column:config;comment:配置" json:"config"`                                                               //详细信息
	CreatedAt *time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at;not null;autoUpdateTime;default:current_timestamp();comment:更新时间" json:"updated_at"` // 更新时间
}
