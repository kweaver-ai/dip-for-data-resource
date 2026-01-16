package data_subject

import (
	"context"
)

type Driven interface {
	GetDataSubjectByPath(ctx context.Context, paths *GetDataSubjectByPathReq) (res *GetDataSubjectByPathRes, err error)
	GetDataSubjectByID(ctx context.Context, ids []string) (res *DataSubjectListObject, err error)
	GetSubjectList(ctx context.Context, parentId, subjectType string) (*DataSubjectListRes, error)
}

//region GetDataSubjectByPath

type GetDataSubjectByPathReq struct {
	Paths []string `json:"paths" binding:"required,gt=0,unique"`
}

type GetDataSubjectByPathRes struct {
	DataSubjects map[string]*DataSubjectInternal `json:"departments"`
}

type DataSubjectInternal struct {
	DomainID    uint64 `json:"domain_id"`   // 雪花id
	ID          string `json:"id"`          // 对象id，uuid
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	Type        int8   `json:"type"`        // 类型：1：业务域，2：主题域，3：业务对象，4：业务活动，5：逻辑实体，6：属性
	PathID      string `json:"path_id"`     // 路径ID
	Path        string `json:"path"`        // 路径
	Owners      string `json:"owners"`      // 拥有者
}

type DataSubjectList struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	PathID      string `json:"path_id"`
	PathName    string `json:"path_name"`
	Owners      struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
	} `json:"owners"`
	CreateBy  string `json:"created_by"`
	CreateAt  int64  `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt int64  `json:"updated_at"`
}
type DataSubjectListObject struct {
	Objects []DataSubjectList `json:"object"`
}

func (d *DataSubjectListObject) ToDataSubjectMap() (dataSubjectMap map[string]DataSubjectList) {
	dataSubjectMap = make(map[string]DataSubjectList)
	for _, v := range d.Objects {
		dataSubjectMap[v.ID] = v
	}
	return
}

// GetSubjectList

type DataSubjectListRes struct {
	Entries    []DataSubject `json:"entries"`
	TotalCount int           `json:"total_count"`
}
type DataSubject struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	PathId           string   `json:"path_id"`
	PathName         string   `json:"path_name"`
	Owners           []string `json:"owners"`
	CreatedBy        string   `json:"created_by"`
	CreatedAt        int64    `json:"created_at"`
	UpdatedBy        string   `json:"updated_by"`
	UpdatedAt        int64    `json:"updated_at"`
	ChildCount       int      `json:"child_count"`
	SecondChildCount int      `json:"second_child_count"`
}

type GetAttributRes struct {
	Attributes []*GetAttributResp `json:"attributes"`
}
type GetAttributResp struct {
	ID          string `json:"id"`          // 对象id
	Name        string `json:"name"`        // 对象名称
	Description string `json:"description"` // 描述
	Type        string `json:"type"`        // 对象类型
	PathID      string `json:"path_id"`     // 路径id
	PathName    string `json:"path_name"`   // 路径名称
	LabelId     string `json:"label_id"`    // 标签ID
	LabelName   string `json:"label_name"`  // 标签名称
	LabelIcon   string `json:"label_icon"`  // 标签颜色
	LabelPath   string `json:"label_path"`  //标签路径
}

// endregion
