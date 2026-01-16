package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// KnowledgeNetworkInfoRecord  记录加
type KnowledgeNetworkInfoRecord struct {
	ID        string                `gorm:"column:f_id;not null;comment:逻辑主键" json:"id"`                                                            // 逻辑主键
	Name      string                `gorm:"column:f_name;not null;comment:资源名称" json:"name"`                                                        // 资源名称
	Version   int32                 `gorm:"column:f_version;not null;comment:资源版本" json:"version"`                                                  // 资源版本
	Type      int32                 `gorm:"column:f_type;not null;comment:资源类型；1:知识网络；2:数据源；3:图谱；4:图分析服务" json:"type"`                              // 资源类型；1:知识网络；2:数据源；3:图谱；4:图分析服务
	ConfigID  string                `gorm:"column:f_config_id;not null;comment:资源对应的配置ID，配置文件中定义" json:"config_id"`                                 // 资源对应的配置ID，配置文件中定义
	RealID    string                `gorm:"column:f_real_id;not null;comment:资源在所属平台的ID，eg：知识网络在AD平台上的id" json:"real_id"`                           // 资源在所属平台的ID，eg：知识网络在AD平台上的id
	Detail    *string               `gorm:"column:f_detail;comment:详细信息" json:"detail"`                                                             //详细信息
	CreatedAt *time.Time            `gorm:"column:f_created_at;not null;comment:创建时间" json:"created_at"`                                            // 创建时间
	UpdatedAt *time.Time            `gorm:"column:f_updated_at;not null;autoUpdateTime;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:f_deleted_at;not null;comment:删除时间戳;softDelete:milli" json:"deleted_at"`                          // 删除时间戳
}
