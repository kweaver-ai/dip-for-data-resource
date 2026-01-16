package es_subject_model

import (
	"context"
)

const (
	ngram = ".ngram"
	graph = ".graph"

	Score            = "_score" // 算分
	ID               = "id"     // ID
	Code             = "code"   //
	Name             = "name"   // 名称
	NameNgram        = Name + ".ngram"
	NameGraph        = Name + ".graph"
	Description      = "description" // 数据目录描述
	DescriptionNgram = Description + ".ngram"
	DescriptionGraph = Description + ".graph"

	OrgCode      = "orgcode" // 组织架构ID
	OrgName      = "orgname" // 组织架构名称
	OrgNameNgram = OrgName + ngram
	OrgNameGraph = OrgName + graph

	SubjectDomainID = "subject_domain_id" // 组织架构ID

	SubjectDomainName      = "subject_domain_name" // 组织架构名称
	SubjectDomainIDNgram   = SubjectDomainID + ngram
	SubjectDomainNameGraph = SubjectDomainName + graph

	DataOwnerName      = "data_owner_name"
	DataOwnerNameGraph = "data_owner_name" + ngram
	DataOwnerNameNgram = "data_owner_name" + graph

	Fields      = "fields"        // 字段
	FieldNameZH = "field_name_zh" // 字段中文名称
	FieldNameEN = "field_name_en" // 字段英文名称

	FieldNameZHNgram = FieldNameZH + ngram
	FieldNameZHGraph = FieldNameZH + graph
	FieldNameENNgram = FieldNameEN + ngram
	FieldNameENGraph = FieldNameEN + graph

	UpdatedAt   = "updated_at"   // 更新时间
	OnlineAt    = "online_at"    // 发布时间
	PublishedAt = "published_at" // 发布时间
)

type ESSubjectModel interface {
	//Search(ctx context.Context, param *SearchParam) (*SearchResult, error)
	Index(ctx context.Context, doc *SubjectModelDoc) error
	IndexLabel(ctx context.Context, doc *SubjectModelLabelDoc) error
	IndexEntityRule(ctx context.Context, doc *EntityRuleDoc) (err error)
	IndexEntityFormDoc(ctx context.Context, doc *EntityFormDoc) (err error)
	//IndexEntityFlowchart(ctx context.Context, doc *EntityFlowchart) (err error)
	IndexEntityDataElement(ctx context.Context, doc *EntityDataElement) (err error)
	IndexEntityField(ctx context.Context, doc *EntityField) (err error)
	//IndexEntityBusinessIndicator(ctx context.Context, doc *EntityBusinessIndicator) (err error)
	IndexEntityFormView(ctx context.Context, doc *EntityFormView) (err error)
	IndexEntitySubjectProperty(ctx context.Context, doc *EntitySubjectProperty) (err error)
	Delete(ctx context.Context, id string) error
	DeleteLabel(ctx context.Context, id string) (err error)
	DeleteIndex(ctx context.Context, id string, indexAlias string) (err error)
}

type SearchParam struct {
	BaseSearchParam
	Orders   []Order
	Size     int
	NextFlag []string
}

type Order struct {
	Direction string
	Sort      string
}

type BaseSearchParam struct {
	Keyword         string   // 关键字
	OrgCode         []string // 组织架构ID
	SubjectDomainID []string
}

type SubjectModelDoc struct {
	DocID string `json:"doc_id"`
	BaseObj
}

type BaseObj struct {
	ID             string `json:"id"`             // 逻辑视图id
	BusinessName   string `json:"business_name"`  // 逻辑视图名称
	TechnicalName  string `json:"technical_name"` // 逻辑视图名称
	Description    string `json:"description"`    // 逻辑视图描述
	DataViewID     string `json:"data_view_id"`
	DisplayFieldID string `json:"display_field_id"`
	DeletedAt      int64  `json:"deleted_at"`
}

type SubjectModelLabelDoc struct {
	DocID           string   `json:"doc_id"`
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	RelatedModelIds []string `json:"related_model_ids"`
}

type EntityRuleDoc struct {
	DocID         string `json:"doc_id"`
	ID            string `json:"id"`
	CatalogId     string `json:"catalog_id"`
	Name          string `json:"name"`
	OrgType       string `json:"org_type"`
	Description   string `json:"description"`
	RuleType      string `json:"rule_type"`
	Expression    string `json:"expression"`
	DepartmentIds string `json:"department_ids"`
}

type EntityFormDoc struct {
	DocID           string `json:"doc_id"`
	ID              string `json:"id"`
	BusinessModelID string `json:"business_model_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
}

type EntityFlowchart struct {
	DocID           string `json:"doc_id"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	PathID          string `json:"path_id"`
	Description     string `json:"description"`
	BusinessModelID string `json:"business_model_id"`
}

type EntityDataElement struct {
	DocID         string `json:"doc_id"`
	ID            string `json:"id"`
	Code          int64  `json:"code"`
	NameEN        string `json:"name_en"`
	RuleID        string `json:"rule_id"`
	DepartmentIds string `json:"department_ids"`
	NameCn        string `json:"name_cn"`
	StdType       string `json:"std_type"`
}

type EntityField struct {
	DocID            string `json:"doc_id"`
	ID               string `json:"id"`
	BusinessFormID   string `json:"business_form_id"`
	BusinessFormName string `json:"business_form_name"`
	Name             string `json:"name"`
	NameEn           string `json:"name_en"`
	StandardID       string `json:"standard_id"`
}

type EntityBusinessIndicator struct {
	DocID              string `json:"doc_id"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	CalculationFormula string `json:"calculation_formula"`
	Unit               string `json:"unit"`
	StatisticsCycle    string `json:"statistics_cycle"`
	StatisticalCaliber string `json:"statistical_caliber"`
	BusinessModelID    string `json:"business_model_id"`
}

type EntityFormView struct {
	DocID         string `json:"doc_id"`
	ID            string `json:"id"`
	TechnicalName string `json:"technical_name"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	DatasourceID  string `json:"datasource_id"`
	SubjectID     string `json:"subject_id"`
	Description   string `json:"description"`
}

type EntitySubjectProperty struct {
	DocID       string `json:"doc_id"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PathID      string `json:"path_id"`
	Path        string `json:"path"`
	StandardID  string `json:"standard_id"`
}

type EntityLabel struct {
	DocID               string `json:"doc_id"`
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	CategoryID          string `json:"category_id"`
	CategoryName        string `json:"category_name"`
	CategoryRangeType   string `json:"category_range_type"`
	CategoryDescription string `json:"category_description"`
	Path                string `json:"path"`
	Sort                int    `json:"sort"`
}

type SearchResult struct {
	Items      []SearchResultItem
	TotalCount int64
	NextFlag   []string
}

type SearchResultItem struct {
	BaseObj
	RawName          string `json:"raw_name,omitempty"`
	RawCode          string `json:"raw_code,omitempty"`
	RawDescription   string `json:"raw_description,omitempty"`
	RawOrgName       string `json:"raw_orgname"`
	RawDataOwnerName string `json:"raw_data_owner_name"`
}
