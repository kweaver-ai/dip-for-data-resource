package data_change_mq

//type UseCase interface {
//	IndexToES(ctx context.Context, req *IndexToESReqParam) (*IndexToESRespParam, error)
//	DeleteFromES(ctx context.Context, req *DeleteFromESReqParam) (*DeleteFromESRespParam, error)
//}

type IndexToESReqParam struct {
	DocId         string `json:"docid,omitempty"`        // doc id
	Title         string `json:"title,omitempty"`        // 数据目录名称
	Description   string `json:"description,omitempty"`  // 数据目录描述
	ID            string `json:"id,omitempty"`           // 数据目录id
	Code          string `json:"code"`                   // 目录编码
	DataKind      []int  `json:"data_kind,omitempty"`    // 基础信息分类
	DataRange     *int   `json:"data_range,omitempty"`   // 数据范围
	UpdateCycle   *int   `json:"update_cycle,omitempty"` // 更新频率
	SharedType    int    `json:"shared_type,omitempty"`  // 共享条件
	OrgCode       string `json:"orgcode,omitempty"`      // 组织架构ID
	OrgName       string `json:"orgname,omitempty"`      // 组织架构名称
	TableId       string `json:"table_id,omitempty"`     // 库表ID
	TableRows     *int64 `json:"table_rows,omitempty"`   // 数据量
	DataUpdatedAt *int64 `json:"updated_at,omitempty"`   // 数据更新时间
	PublishedAt   int64  `json:"published_at,omitempty"` // 上线发布时间

	BusinessObjects []*BusinessObjectEntity `json:"business_objects"` //业务对象ID数组，里面ID用于左侧树业务域选中节点筛选
	InfoSystems     []*InfoSystemEntity     `json:"info_systems"`     // 信息系统列表，里面ID用于左侧树信息系统选中节点筛选

	InfoSystemID   string `json:"info_system_id"`   // 信息系统ID
	InfoSystemName string `json:"info_system_name"` // 信息系统名称

	OwnerName      string `json:"owner_name"` // 数据Owner名称
	OwnerID        string `json:"owner_id"`   // 数据OwnerID
	DataSourceName string `json:"data_source_name,omitempty" binding:"omitempty"`
	DataSourceID   string `json:"data_source_id,omitempty" binding:"omitempty"`
	SchemaName     string `json:"schema_name,omitempty" binding:"omitempty"`
	SchemaID       string `json:"schema_id,omitempty" binding:"omitempty"`
}

type IndexToESRespParam struct {
	ID string
}
type DeleteFromESRespParam struct {
	ID string
}

//func NewDeleteFromESRespParam(id string) *DeleteFromESRespParam {
//	return &DeleteFromESRespParam{ID: id}
//}

type FieldEntity struct {
	RawFieldNameZH string `json:"raw_field_name_zh"`
	FieldNameZH    string `json:"field_name_zh"`
	RawFieldNameEN string `json:"raw_field_name_en"`
	FieldNameEN    string `json:"field_name_en"`
}

type BusinessObjectEntity struct {
	ID   string `json:"id" binding:"omitempty,uuid"` // 业务对象id
	Name string `json:"name"`                        // 业务对象名称
}

type InfoSystemEntity struct {
	ID   string `json:"id" binding:"omitempty,uuid"` // 业务对象id
	Name string `json:"name"`                        // 业务对象名称
}

type DeleteFromESReqParam struct {
	ID string
}

type UpdateTableRowsAndUpdatedAtReqParam struct {
	TableId       string `json:"table_id,omitempty"`   // 库表ID
	TableRows     *int64 `json:"table_rows,omitempty"` // 数据量
	DataUpdatedAt *int64 `json:"updated_at,omitempty"` // 数据更新时间
}
