package model

import (
	//"github.com/kweaver-ai/idrm-go-frame/core/utils/utilities"
	//"devops.xxx.cn/AISHUDevOps/AnyFabric/_git/data-catalog/common/util"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
	"unsafe"
	//"gorm.io/gorm"
)

type TDataCatalog struct {
	ID    uint64 `gorm:"column:id" json:"id,string"`         // 唯一id，雪花算法
	Code  string `gorm:"column:code;not null" json:"code"`   // 目录编码
	Title string `gorm:"column:title;not null" json:"title"` // 目录名称
	//GroupID          uint64 `gorm:"column:group_id;not null;default:0" json:"group_id,string,omitempty"` // 数据资源目录分类ID
	//GroupName        string `gorm:"column:group_name;not null" json:"group_name,omitempty"`              // 数据资源目录分类名称
	//ThemeID          uint64 `gorm:"column:theme_id" json:"theme_id,string,omitempty"`                    // 主题分类ID
	//ThemeName        string `gorm:"column:theme_name" json:"theme_name,omitempty"`                       // 主题分类名称
	//ForwardVersionID uint64 `gorm:"column:forward_version_id" json:"forward_version_id,omitempty"`       // 当前目录前一版本目录ID
	Description string `gorm:"column:description" json:"description,omitempty"` // 资源目录描述
	//Version          string `gorm:"column:version;not null;default:0.0.0.1" json:"version"`              // 目录版本号，默认初始版本为0.0.0.1
	//DataRange        int32  `gorm:"column:data_range" json:"data_range,omitempty"`                       // 数据范围：字典DM_DATA_SJFW，01全市 02市直 03区县
	//UpdateCycle      int32  `gorm:"column:update_cycle" json:"update_cycle,omitempty"`                   // 更新频率 参考数据字典：GXZQ，1不定时 2实时 3每日 4每周 5每月 6每季度 7每半年 8每年 9其他
	//DataKind         int32  `gorm:"column:data_kind;not null" json:"data_kind"`                          // 基础信息分类 1 人 2 地 4 事 8 物 16 组织 32 其他  可组合，如 人和地 即 1|2 = 3
	//SharedType       int8   `gorm:"column:shared_type;not null" json:"shared_type"`                      // 共享属性 1 无条件共享 2 有条件共享 3 不予共享
	//SharedCondition  string `gorm:"column:shared_condition" json:"shared_condition,omitempty"`           // 共享条件
	//OpenType         int8   `gorm:"column:open_type;not null" json:"open_type,omitempty"`                // 开放属性 1 向公众开放 2 不向公众开放
	//OpenCondition    string `gorm:"column:open_condition" json:"open_condition,omitempty"`               // 开放条件
	//SharedMode       int8   `gorm:"column:shared_mode;not null" json:"shared_mode,omitempty"`            // 共享方式 1 共享平台方式 2 邮件方式 3 介质方式
	//PhysicalDeletion *int8  `gorm:"column:physical_deletion" json:"physical_deletion,omitempty"`         // 挂接实体资源是否存在物理删除(1 是 ; 0 否)
	//SyncMechanism    int8   `gorm:"column:sync_mechanism" json:"sync_mechanism,omitempty"`               // 数据归集机制(1 增量 ; 2 全量) ----归集到数据中台
	//SyncFrequency    string `gorm:"column:sync_frequency" json:"sync_frequency,omitempty"`               // 数据归集频率 ----归集到数据中台
	//TableCount       int32   `gorm:"column:table_count;not null" json:"table_count"`                               // 挂接库表数量
	//FileCount      int32   `gorm:"column:file_count;not null" json:"file_count"`                                 // 挂接文件数量
	//State          int8    `gorm:"column:state;not null;default:1" json:"state"`                                 // 目录状态，1 草稿  3 已发布  5 已上线  8 已下线
	OnlineStatus string `gorm:"column:online_status" json:"online_status,omitempty"` // 目录状态，1 草稿  3 已发布  5 已上线  8 已下线
	//FlowNodeID     string  `gorm:"column:flow_node_id" json:"flow_node_id"`                                      // 目录当前所处审核流程结点ID
	//FlowNodeName   string  `gorm:"column:flow_node_name" json:"flow_node_name"`                                  // 目录当前所处审核流程结点名称
	//FlowType       *int8   `gorm:"column:flow_type" json:"flow_type"`                                            // 审批流程类型，1 上线  2 变更  3 下线  4 发布
	//FlowID         string  `gorm:"column:flow_id" json:"flow_id"`                                                // 审批流程实例ID
	//FlowName       string  `gorm:"column:flow_name" json:"flow_name"`                                            // 审批流程名称
	//FlowVersion    string  `gorm:"column:flow_version" json:"flow_version"`                                      // 审批流程版本
	DepartmentId   string `gorm:"column:department_id" json:"department_id,omitempty"`     // 所属部门ID
	DepartmentName string `gorm:"column:department_name" json:"department_name,omitempty"` // 所属部门名称
	//CreatedAt      *Time  `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"` // 创建时间
	//CreatorUID     string `gorm:"column:creator_uid;not null" json:"creator_uid,omitempty"`                 // 创建用户ID
	//CreatorName    string `gorm:"column:creator_name;not null" json:"creator_name,omitempty"`               // 创建用户名称
	//UpdatedAt      *Time  `gorm:"column:updated_at;not null" json:"updated_at"`                             // 更新时间
	//UpdaterUID     string `gorm:"column:updater_uid" json:"updater_uid,omitempty"`                          // 更新用户ID
	//UpdaterName    string `gorm:"column:updater_name" json:"updater_name,omitempty"`                        // 更新用户名称
	//DeletedAt      *Time  `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`                           // 删除时间
	//DeleteUID      string  `gorm:"column:delete_uid" json:"delete_uid,omitempty"`                                // 删除用户ID
	//DeleteName     string  `gorm:"column:delete_name" json:"delete_name,omitempty"`                              // 删除用户名称
	//Source         int8    `gorm:"column:source;not null;default:1" json:"source"`                               // 数据来源 1 认知平台自动创建 2 人工创建
	//TableType      int8    `gorm:"column:table_type" json:"table_type,omitempty"`                                // 库表类型 1 贴源表 2 标准表
	//CurrentVersion *int8   `gorm:"column:current_version;not null;default:1" json:"current_version,omitempty"`   // 是否先行版本 0 否 1 是
	//PublishFlag    *int8   `gorm:"column:publish_flag;not null;default:0" json:"publish_flag,omitempty"`         // 是否发布到超市 (1 是 ; 0 否)
	//DataKindFlag   *int8   `gorm:"column:data_kind_flag;not null;default:0" json:"data_kind_flag,omitempty"`     // 基础信息分类是否智能推荐 (1 是 ; 0 否)
	//LabelFlag      *int8   `gorm:"column:label_flag;not null;default:0" json:"label_flag,omitempty"`             // 标签是否智能推荐 (1 是 ; 0 否)
	//SrcEventFlag   *int8   `gorm:"column:src_event_flag;not null;default:0" json:"src_event_flag,omitempty"`     // 来源业务场景是否智能推荐 (1 是 ; 0 否)
	//RelEventFlag   *int8   `gorm:"column:rel_event_flag;not null;default:0" json:"rel_event_flag,omitempty"`     // 关联业务场景是否智能推荐 (1 是 ; 0 否)
	//SystemFlag     *int8   `gorm:"column:system_flag;not null;default:0" json:"system_flag,omitempty"`           // 关联信息系统是否智能推荐 (1 是 ; 0 否)
	//RelCatalogFlag *int8   `gorm:"column:rel_catalog_flag;not null;default:0" json:"rel_catalog_flag,omitempty"` // 关联目录是否智能推荐 (1 是 ; 0 否)
	PublishedAt *Time `gorm:"column:published_at" json:"published_at,omitempty"` // 上线发布时间
	//IsIndexed      *int8   `gorm:"column:is_indexed" json:"is_indexed"`                                          // 是否已建ES索引，0 否 1 是，默认为0
	//AuditApplySN   uint64  `gorm:"column:audit_apply_sn;not null" json:"audit_apply_sn"`                         // 发起审核申请序号
	//AuditAdvice    string  `gorm:"audit_advice" json:"audit_advice"`
	OwnerId   string `gorm:"column:owner_id" json:"owner_id,omitempty"`     // 目录数据owner的用户ID
	OwnerName string `gorm:"column:owner_name" json:"owner_name,omitempty"` // 审核意见，仅驳回时有用
	//AuditState     *int8   `gorm:"audit_state" json:"audit_state"`                        // 审核状态，1 审核中  2 通过  3 驳回
	//IsCanceled     *int8   `gorm:"is_canceled" json:"is_canceled"`                        // 目录下线时针对该目录的关联申请是否已撤销，0 待撤销 1 已撤销
	//ProcDefKey     string  `gorm:"column:proc_def_key" json:"proc_def_key"`               // 审核流程key
	//FlowApplyId    string  `gorm:"column:flow_apply_id" json:"flow_apply_id"`             // 审核流程ID
	//ExploreJobId   *string `gorm:"column:explore_job_id" json:"explore_job_id"`           // 探查作业ID
	//ExploreJobVer  *int    `gorm:"column:explore_job_version" json:"explore_job_version"` // 探查作业版本
}

type Time struct {
	time.Time
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte{}, nil
	}

	// return StringToBytes(fmt.Sprintf("\"%s\"", t.Format(constant.LOCAL_TIME_FORMAT))), nil
	return StringToBytes(fmt.Sprintf("%d", t.UnixMilli())), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}

	// str = strings.Trim(str, "\"")
	// val, err := time.Parse(constant.LOCAL_TIME_FORMAT, str)
	// *t = Time{val}
	ts, err := strconv.ParseInt(str, 10, 64)
	*t = Time{time.UnixMilli(ts)}
	return err
}

func (t *Time) Scan(value interface{}) error {
	val, ok := value.(time.Time)
	if ok {
		*t = Time{val}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}

func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

type TDataCatalogColumn struct {
	ID              uint64 `gorm:"column:id" json:"id,string"`                                // 唯一id，雪花算法
	CatalogID       uint64 `gorm:"column:catalog_id;not null" json:"catalog_id,string"`       // 数据资源目录ID
	TechnicalName   string `gorm:"column:technical_name;not null" json:"technical_name"`      // 字段名称
	BusinessName    string `gorm:"column:business_name;not null" json:"business_name"`        // 信息项名称
	DataFormat      *int32 `gorm:"column:data_format;not null" json:"data_format"`            // 字段类型
	DataLength      *int32 `gorm:"column:data_length;not null" json:"data_length"`            // 字段长度
	DatametaID      string `gorm:"column:datameta_id" json:"datameta_id"`                     // 关联数据元ID
	DatametaName    string `gorm:"column:datameta_name" json:"datameta_name"`                 // 关联数据元名称
	Ranges          string `gorm:"column:ranges" json:"ranges"`                               // 字段值域
	CodesetID       string `gorm:"column:codeset_id" json:"-"`                                // 关联代码集ID
	CodesetName     string `gorm:"column:codeset_name" json:"-"`                              // 关联代码集名称
	SharedType      *int8  `gorm:"column:shared_type;not null" json:"shared_type"`            // 共享属性 1 无条件共享 2 有条件共享 3 不予共享
	OpenType        *int8  `gorm:"column:open_type;not null" json:"open_type"`                // 开放属性 1 向公众开放 2 不向公众开放
	TimestampFlag   *int8  `gorm:"column:timestamp_flag;not null" json:"timestamp_flag"`      // 是否时间戳(1 是 ; 0 否)
	PrimaryFlag     *int8  `gorm:"column:primary_flag;not null" json:"primary_flag"`          // 是否主键(1 是 ; 0 否)
	NullFlag        *int8  `gorm:"column:null_flag;not null" json:"null_flag"`                // 是否为空(1 是 ; 0 否)
	ClassifiedFlag  *int8  `gorm:"column:classified_flag;not null" json:"classified_flag"`    // 是否涉密属性(1 是 ; 0 否)
	SensitiveFlag   *int8  `gorm:"column:sensitive_flag;not null" json:"sensitive_flag"`      // 是否敏感属性(1 是 ; 0 否)
	Description     string `gorm:"column:description" json:"description"`                     // 字段描述
	AIDescription   string `gorm:"column:ai_description" json:"ai_description"`               // 在AI数据理解下生成的字段描述
	SharedCondition string `gorm:"column:shared_condition" json:"shared_condition,omitempty"` // 共享条件
	OpenCondition   string `gorm:"column:open_condition" json:"open_condition,omitempty"`     // 开放条件
}

type TDataCatalogInfo struct {
	ID        uint64 `gorm:"column:id" json:"id"`                          // 唯一id，雪花算法
	CatalogID uint64 `gorm:"column:catalog_id;not null" json:"catalog_id"` // 数据资源目录ID
	InfoType  int8   `gorm:"column:info_type;not null" json:"info_type"`   // 关联信息类型 1 标签 2 来源业务场景 3 关联业务场景 4 关联系统 5 关联目录、信息项 6 业务域
	InfoKey   string `gorm:"column:info_key;not null" json:"info_key"`     // 关联信息key（仅当info_type为5时为关联目录ID，其它情况下为ID或枚举值）
	InfoValue string `gorm:"column:info_value;not null" json:"info_value"` // 关联信息名称（info_type为5时表示关联目录及其信息项的json字符串）
}

type TDataCatalogResourceMount struct {
	ID        uint64  `gorm:"column:id" json:"id"`                          // 唯一id，雪花算法
	CatalogID uint64  `gorm:"column:catalog_id;not null" json:"catalog_id"` // 数据资源目录ID
	ResType   int8    `gorm:"column:res_type;not null" json:"res_type"`     // 挂接资源类型 1 库表 2 文件
	ResID     uint64  `gorm:"column:res_id;not null" json:"res_id"`         // 挂接资源ID
	ResIDStr  *string `gorm:"column:s_res_id;not null" json:"s_res_id"`     // 挂接资源ID（字符型），res_type为1即库表时填写数据表视图ID
	ResName   string  `gorm:"column:res_name;not null" json:"res_name"`     // 挂接资源名称
	Code      string  `gorm:"column:code;not null" json:"code"`             // 目录编码
}

type TDataResource struct {
	ID            uint64 `gorm:"column:id" json:"id"`                                  // 唯一id，雪花算法
	ResourceId    string `gorm:"column:resource_id;not null" json:"resource_id"`       // 资源ID
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"` // 挂接资源名称
	CatalogId     uint64 `gorm:"column:catalog_id;not null" json:"catalog_id"`         // 目录编码
}

type TUserDataCatalogRel struct {
	ID          uint64 `gorm:"column:id" json:"id"`                                                      // 唯一id，雪花算法
	UID         string `gorm:"column:uid;not null" json:"uid"`                                           // 用户ID
	Code        string `gorm:"column:code;not null" json:"code"`                                         // 目录编码
	ApplyID     uint64 `gorm:"column:apply_id;not null" json:"apply_id"`                                 // 申请记录ID
	CreatedAt   *Time  `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"` // 记录创建时间
	UpdatedAt   *Time  `gorm:"column:updated_at;not null;default:current_timestamp()" json:"updated_at"` // 记录更新时间
	ExpiredAt   *Time  `gorm:"column:expired_at" json:"expired_at"`                                      // 权限过期时间
	ExpiredFlag int8   `gorm:"column:expired_flag;not nul" json:"expired_flag"`                          // 权限过期标记 1 未过期 2 已过期
}

type TDataCatalogDownloadApply struct {
	ID           uint64 `gorm:"column:id" json:"id"`                                                      // 唯一id，雪花算法
	UID          string `gorm:"column:uid;not null" json:"uid"`                                           // 用户ID
	Code         string `gorm:"column:code;not null" json:"code"`                                         // 目录编码
	ApplyDays    int    `gorm:"column:apply_days;not null" json:"apply_days"`                             // 申请天数（7、15、30天）
	ApplyReason  string `gorm:"column:apply_reason;not null" json:"apply_reason"`                         // 申请理由
	AuditApplySN uint64 `gorm:"column:audit_apply_sn;not null" json:"audit_apply_sn"`                     // 发起审核申请序号
	AuditType    string `gorm:"column:audit_type;not null" json:"audit_type"`                             // 审核类型，默认af-data-catalog-download
	State        int8   `gorm:"column:state;not null;default:1" json:"state"`                             // 申请审核状态 1 审核中 2 审核通过 3 审核不通过
	CreatedAt    *Time  `gorm:"column:created_at;not null;default:current_timestamp()" json:"created_at"` // 审核创建时间
	UpdatedAt    *Time  `gorm:"column:updated_at;not null" json:"updated_at"`                             // 审核结果更新时间
	FlowID       string `gorm:"column:flow_id" json:"flow_id"`                                            // 审批流程实例ID
	ProcDefKey   string `gorm:"column:proc_def_key" json:"proc_def_key"`                                  // 审核流程key
	FlowApplyId  string `gorm:"column:flow_apply_id" json:"flow_apply_id"`                                // 审核流程ID
}
