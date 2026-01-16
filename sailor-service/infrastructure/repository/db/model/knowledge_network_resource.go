package model

// 数据资源

type EntityDataResource struct {
	ID                    string `gorm:"column:id" json:"id"`                                  // 唯一id，雪花算法
	Code                  string `gorm:"column:code;not null" json:"code"`                     // 用户ID
	TechnicalName         string `gorm:"column:technical_name;not null" json:"technical_name"` // 目录编码
	Name                  string `gorm:"column:name;not null" json:"name"`                     // 申请天数（7、15、30天）
	Description           string `gorm:"column:description;not null" json:"description"`       // 申请理由
	AssetType             int8   `gorm:"column:asset_type;not null" json:"asset_type"`         // 发起审核申请序号
	Color                 string `gorm:"column:color;not null" json:"color"`                   // 审核类型，默认af-data-catalog-download
	PublishedAt           int64  `gorm:"column:published_at;not null" json:"published_at"`     // 申请审核状态 1 审核中 2 审核通过 3 审核不通过
	OnlineAt              int64  `gorm:"column:online_at;not null" json:"online_at"`
	PublishStatus         string `gorm:"column:publish_status;not null" json:"publish_status"` // 审核类型，默认af-data-catalog-download
	OnlineStatus          string `gorm:"column:online_status;not null" json:"online_status"`
	PublishStatusCategory string `gorm:"column:publish_status_category;not null" json:"publish_status_category"`
	OwnerId               string `gorm:"column:owner_id;not null" json:"owner_id"`      // 审核创建时间
	OwnerName             string `gorm:"column:owner_name;not null" json:"owner_name"`  // 审核结果更新时间
	DepartmentId          string `gorm:"column:department_id" json:"department_id"`     // 审批流程实例ID
	DepartmentName        string `gorm:"column:department_name" json:"department_name"` // 审核流程key
	DepartmentPathId      string `gorm:"column:department_path_id" json:"department_path_id"`
	DepartmentPath        string `gorm:"column:department_path" json:"department_path"`
	SubjectId             string `gorm:"column:subject_id" json:"subject_id"`
	SubjectName           string `gorm:"column:subject_name" json:"subject_name"`
	SubjectPathId         string `gorm:"column:subject_path_id" json:"subject_path_id"`
	SubjectPath           string `gorm:"column:subject_path" json:"subject_path"`
}

type EntityIndicator struct {
	ID                string `gorm:"column:id" json:"id"` // 唯一id，雪花算法
	AnalysisDimension string `gorm:"column:analysis_dimension;not null" json:"analysis_dimension"`
}

type EdgeDataView2DataExploreReport struct {
	//FormviewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	//ColumnId     string `gorm:"column:column_id;not null" json:"column_id"`
	FormViewId    string `gorm:"column:form_view_id;not null" json:"form_view_id"`
	ColumnId      string `gorm:"column:column_id;not null" json:"column_id"`
	ExploreItem   string `gorm:"column:explore_item;not null" json:"explore_item"`
	ExploreResult string `gorm:"column:explore_result;not null" json:"explore_result"`
}

type EdgeInterface2ResponseField struct {
	FieldSid string `gorm:"column:field_sid;not null" json:"field_sid"`
	Id       string `gorm:"column:id;not null" json:"id"`
}

type EdgeDataView2Filed struct {
	ColumnId     string `gorm:"column:column_id;not null" json:"column_id"`
	FormViewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
}

type EdgeDataView2MetadataSchema struct {
	FormviewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	SchemaSid    string `gorm:"column:schema_sid;not null" json:"schema_sid"`
}

type EntitySubdomain struct {
	Id         string `gorm:"column:id;not null" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	PrefixName string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EdgeSubdomain2DataView struct {
	ID        string `gorm:"column:id;not null" json:"id"`
	SubjectId string `gorm:"column:subject_id;not null" json:"subject_id"`
}

type EdgeSubdomain2Domain struct {
	DomainId string `gorm:"column:domain_id;not null" json:"domain_id"`
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
}

type EntityDomain struct {
	Id         string `gorm:"column:id;not null" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	PrefixName string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EdgeDomain2Subdomain struct {
	DomainId string `gorm:"column:domain_id;not null" json:"domain_id"`
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
}

type EntityMetadataSchema struct {
	SchemaSid  string `gorm:"column:schema_sid;not null" json:"schema_sid"`
	SchemaName string `gorm:"column:schema_name;not null" json:"schema_name"`
	PrefixName string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EdgeMetadataSchema2DataView struct {
	FormViewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	SchemaSid    string `gorm:"column:schema_sid;not null" json:"schema_sid"`
}

type EdgeMetadataSchema2DataSource struct {
	DataSourceUuid string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	SchemaSid      string `gorm:"column:schema_sid;not null" json:"schema_sid"`
}

type EntityDataSource struct {
	DataSourceUuid     string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	DataSourceName     string `gorm:"column:data_source_name;not null" json:"data_source_name"`
	DataSourceTypeName string `gorm:"column:data_source_type_name;not null" json:"data_source_type_name"`
	SourceTypeCode     string `gorm:"column:source_type_code;not null" json:"source_type_code"`
	SourceTypeName     string `gorm:"column:source_type_name;not null" json:"source_type_name"`
	PrefixName         string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EdgeDataSource2MetadataSchema struct {
	DataSourceUuid string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	SchemaSid      string `gorm:"column:schema_sid;not null" json:"schema_sid"`
}

type EdgeDataSourceAndMetadataSchemaByMetadataSchema struct {
	DataSourceUuid     string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	SchemaSid          string `gorm:"column:schema_sid;not null" json:"schema_sid"`
	SchemaName         string `gorm:"column:schema_name;not null" json:"schema_name"`
	DataSourceName     string `gorm:"column:data_source_name;not null" json:"data_source_name"`
	DataSourceTypeName string `gorm:"column:data_source_type_name;not null" json:"data_source_type_name"`
	SourceTypeCode     string `gorm:"column:source_type_code;not null" json:"source_type_code"`
	SourceTypeName     string `gorm:"column:source_type_name;not null" json:"source_type_name"`
	PrefixName         string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EntityDataViewFields struct {
	ColumnId      string `gorm:"column:column_id;not null" json:"column_id"`
	FormviewUuid  string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"`
	FieldName     string `gorm:"column:field_name;not null" json:"field_name"`
	DataType      string `gorm:"column:data_type;not null" json:"data_type"`
}

type EdgeDataViewFields2DataView struct {
	FieldSid string `gorm:"column:field_sid;not null" json:"field_sid"`
	CnName   string `gorm:"column:cn_name;not null" json:"cn_name"`
	EnName   string `gorm:"column:en_name;not null" json:"en_name"`
}

type EntityResponseField struct {
	FieldSid string `gorm:"column:field_sid;not null" json:"field_sid"`
	CnName   string `gorm:"column:cn_name;not null" json:"cn_name"`
	EnName   string `gorm:"column:en_name;not null" json:"en_name"`
}

type EdgeResponseField2Interface struct {
	FieldSid string `gorm:"column:field_sid;not null" json:"field_sid"`
	Id       string `gorm:"column:id;not null" json:"id"`
}

type EntityDataExploreReport struct {
	ColumnId           string `gorm:"column:column_id;not null" json:"column_id"`
	ExploreItem        string `gorm:"column:explore_item;not null" json:"explore_item"`
	ColumnName         string `gorm:"column:column_name;not null" json:"column_name"`
	ExploreResult      string `gorm:"column:explore_result;not null" json:"explore_result"`
	ExploreResultValue string `gorm:"column:explore_result_value;not null" json:"explore_result_value"`
}

type EdgeDataExploreReport2DataView struct {
	FormviewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	ColumnId     string `gorm:"column:column_id;not null" json:"column_id"`
}

type EntityDataOwner struct {
	OwnerId   string `gorm:"column:owner_id;not null" json:"owner_id"`
	OwnerName string `gorm:"column:owner_name;not null" json:"owner_name"`
}

type EdgeDataOwner2DataView struct {
	Id      string `gorm:"column:id;not null" json:"id"`
	OwnerId string `gorm:"column:owner_id;not null" json:"owner_id"`
}

type EntityDepartment struct {
	Id   string `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

type EdgeDepartment2DataView struct {
	Id           string `gorm:"column:id;not null" json:"id"`
	DepartmentId string `gorm:"column:department_id;not null" json:"department_id"`
}

type EntityDimensionModel struct {
	Id          int64  `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
}

type EdgeDimensionModel2Resource struct {
	DimensionModelId string `gorm:"column:dimension_model_id;not null" json:"dimension_model_id"`
	Id               string `gorm:"column:id;not null" json:"id"`
}

type EntityIndicatorAnalysisDimension struct {
	FormViewId         string `gorm:"column:formview_id;not null" json:"formview_id"`
	FieldId            string `gorm:"column:field_id;not null" json:"field_id"`
	FieldBusinessName  string `gorm:"column:field_business_name;not null" json:"field_business_name"`
	FieldTechnicalName string `gorm:"column:field_technical_name;not null" json:"field_technical_name"`
	FieldDataType      string `gorm:"column:field_data_type;not null" json:"field_data_type"`
}

type EdgeIndicatorAnalysisDimension2Resource struct {
	InticatorId int64  `gorm:"column:inticator_id;not null" json:"inticator_id"`
	FormviewId  string `gorm:"column:formview_id;not null" json:"formview_id"`
	FieldId     string `gorm:"column:field_id;not null" json:"field_id"`
}

// 业务架构图谱——算法

type EntityEntityDomainGroup struct {
	Id             string `gorm:"column:id;not null" json:"id"`
	Name           string `gorm:"column:name;not null" json:"name"`
	Description    string `gorm:"column:description;not null" json:"description"`
	Path           string `gorm:"column:path;not null" json:"path"`
	PathId         string `gorm:"column:path_id;not null" json:"path_id"`
	DepartmentId   string `gorm:"column:department_id;not null" json:"department_id"`
	BusinessSystem string `gorm:"column:business_system;not null" json:"business_system"`
	ModelId        string `gorm:"column:model_id;not null" json:"model_id"`
}

type EdgeRelationDomainGroup2Domain struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EntityEntityDomain struct {
	Id             string `gorm:"column:id;not null" json:"id"`
	Name           string `gorm:"column:name;not null" json:"name"`
	Description    string `gorm:"column:description;not null" json:"description"`
	Path           string `gorm:"column:path;not null" json:"path"`
	PathId         string `gorm:"column:path_id;not null" json:"path_id"`
	DepartmentId   string `gorm:"column:department_id;not null" json:"department_id"`
	BusinessSystem string `gorm:"column:business_system;not null" json:"business_system"`
	ModelId        string `gorm:"column:model_id;not null" json:"model_id"`
}

type EdgeRelationDomain2Self struct {
}

type EdgeRelationDomain2DomainFlow struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EntityEntityDomainFlow struct {
	Id             string `gorm:"column:id;not null" json:"id"`
	Name           string `gorm:"column:name;not null" json:"name"`
	Description    string `gorm:"column:description;not null" json:"description"`
	Path           string `gorm:"column:path;not null" json:"path"`
	PathId         string `gorm:"column:path_id;not null" json:"path_id"`
	DepartmentId   string `gorm:"column:department_id;not null" json:"department_id"`
	BusinessSystem string `gorm:"column:business_system;not null" json:"business_system"`
	ModelId        string `gorm:"column:model_id;not null" json:"model_id"`
}

//type EdgeRelationDomain2DomainFlow struct {
//}

type EdgeRelationDomainFlow2InfomationSystem struct {
	DomainFlowId       string `gorm:"column:domain_flow_id;not null" json:"domain_flow_id"`
	InfomationSystemId string `gorm:"column:infomation_system_id;not null" json:"infomation_system_id"`
}

type EdgeRelationBusinessModel2Department struct {
	DomainFlowId string `gorm:"column:domain_flow_id;not null" json:"domain_flow_id"`
	DepartmentId string `gorm:"column:department_id;not null" json:"department_id"`
}

type EntityEntityInfomationSystem struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
}

type EntityEntityDepartment struct {
	Id     string `gorm:"column:id;not null" json:"id"`
	Name   string `gorm:"column:name;not null" json:"name"`
	PathId string `gorm:"column:path_id;not null" json:"path_id"`
	Path   string `gorm:"column:path;not null" json:"path"`
}

type EdgeRelationDepartment2Self struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EntityEntityBusinessModel struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	DomainId    string `gorm:"column:domain_id;not null" json:"domain_id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
}

type EdgeRelationBusinessModel2DomainFlow struct {
	BusinessModelId  string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	BusinessDomainId string `gorm:"column:business_domain_id;not null" json:"business_domain_id"`
}

type EdgeRelationBusinessModel2Form struct {
	FormId          string `gorm:"column:form_id;not null" json:"form_id"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
}

type EdgeRelationBusinessModel2Flowchart struct {
	FlowchartId     string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
}

type EntityEntityFlowchart struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
	BusinessModelId string `gorm:"column:business_model_id ;not null" json:"business_model_id"`
	Path            string `gorm:"column:path;not null" json:"path"`
	PathId          string `gorm:"column:path_id;not null" json:"path_id"`
}

type EdgeRelationFlowchart2Self struct {
	FromFlowchartId string `gorm:"column:from_flowchart_id;not null" json:"from_flowchart_id"`
	ToFlowchartId   string `gorm:"column:to_flowchart_id;not null" json:"to_flowchart_id"`
}

type EdgeRelationFlowchart2FlowchartNode struct {
	FlowchartId     string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	FlowchartNodeId string `gorm:"column:flowchart_node_id;not null" json:"flowchart_node_id"`
}

type EntityEntityFlowchartNode struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	MxcellId    string `gorm:"column:mxcell_id;not null" json:"mxcell_id"`
	FlowchartId string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	Source      string `gorm:"column:source;not null" json:"source"`
	Target      string `gorm:"column:target;not null" json:"target"`
}

type EdgeRelationFlowchartNode2Self struct {
	FromFlowchartNodeId string `gorm:"column:from_flowchart_node_id;not null" json:"from_flowchart_node_id"`
	ToFlowchartNodeId   string `gorm:"column:to_flowchart_node_id;not null" json:"to_flowchart_node_id"`
}

type EntityEntityForm struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
}

type EdgeRelationForm2Self struct {
	FromFormId string `gorm:"column:from_form_id;not null" json:"from_form_id"`
	ToFormId   string `gorm:"column:to_form_id;not null" json:"to_form_id"`
}

type EdgeRelationForm2Field struct {
	FormId  string `gorm:"column:form_id;not null" json:"form_id"`
	FieldId string `gorm:"column:field_id;not null" json:"field_id"`
}

type EntityEntityField struct {
	Id               string `gorm:"column:id;not null" json:"id"`
	BusinessFormId   string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	BusinessFormName string `gorm:"column:business_form_name;not null" json:"business_form_name"`
	Name             string `gorm:"column:name;not null" json:"name"`
	NameEn           string `gorm:"column:name_en;not null" json:"name_en"`
	StandardId       string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type EdgeRelationDataElement2Field struct {
	FieldId    string `gorm:"column:field_id;not null" json:"field_id"`
	StandardId string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type EntityEntityDataElement struct {
	Id            int64  `gorm:"column:id;not null" json:"id"`
	Code          int64  `gorm:"column:code;not null" json:"code"`
	NameEn        string `gorm:"column:name_en;not null" json:"name_en"`
	NameCn        string `gorm:"column:name_cn;not null" json:"name_cn"`
	StdType       string `gorm:"column:std_type;not null" json:"std_type"`
	RuleID        *int64 `gorm:"column:rule_id" json:"rule_id"`
	State         int    `gorm:"column:state;not null" json:"state"`
	DepartmentIds string `gorm:"column:department_ids;not null" json:"department_ids"`
}

type EntityEntityLabel struct {
	Id                  int64  `gorm:"column:id;not null" json:"id"`
	Name                string `gorm:"column:name;not null" json:"name"`
	CategoryId          string `gorm:"column:category_id;not null" json:"category_id"`
	CategoryName        string `gorm:"column:category_name;not null" json:"category_name"`
	CategoryRangeType   string `gorm:"column:category_range_type;not null" json:"category_range_type"`
	CategoryDescription string `gorm:"column:category_description;not null" json:"category_description"`
	FPath               string `gorm:"column:f_path;not null" json:"f_path"`
	FSort               int    `gorm:"column:f_sort;not null" json:"f_sort"`
}

type EntityEntityBusinessIndicator struct {
	Id                 string `gorm:"column:id;not null" json:"id"`
	BusinessModelId    string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	Name               string `gorm:"column:name;not null" json:"name"`
	Description        string `gorm:"column:description;not null" json:"description"`
	CalculationFormula string `gorm:"column:calculation_formula;not null" json:"calculation_formula"`
	Unit               string `gorm:"column:unit;not null" json:"unit"`
	StatisticsCycle    string `gorm:"column:statistics_cycle;not null" json:"statistics_cycle"`
	StatisticalCaliber string `gorm:"column:statistical_caliber;not null" json:"statistical_caliber"`
}

type EntityEntityRule struct {
	Id            int64  `gorm:"column:id;not null" json:"id"`
	Name          string `gorm:"column:name;not null" json:"name"`
	CategoryId    string `gorm:"column:category_id;not null" json:"category_id"`
	OrgType       string `gorm:"column:org_type;not null" json:"org_type"`
	Description   string `gorm:"column:description;not null" json:"description"`
	RuleType      string `gorm:"column:rule_type;not null" json:"rule_type"`
	Expression    string `gorm:"column:expression;not null" json:"expression"`
	DepartmentIds string `gorm:"column:department_ids;not null" json:"department_ids"`
}

type EdgeRelationViewField2DataElement struct {
	ViewFieldId string `gorm:"column:view_field_id;not null" json:"view_field_id"`
	StandardId  string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type EdgeRelationSubjectProperty2EntityDataElement struct {
	SubjectPropId string `gorm:"column:subject_prop_id;not null" json:"subject_prop_id"`
	StandardId    string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type EntityEntityFormViewField struct {
	Id            string `gorm:"column:id;not null" json:"id"`
	FormViewId    string `gorm:"column:form_view_id;not null" json:"form_view_id"`
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"`
	Name          string `gorm:"column:name;not null" json:"name"`
	StandardCode  string `gorm:"column:standard_code;not null" json:"standard_code"`
	Standard      string `gorm:"column:standard;not null" json:"standard"`
	CodeTableId   string `gorm:"column:code_table_id;not null" json:"code_table_id"`
}

type EdgeRelationFormView2Field struct {
	ViewFormId      string `gorm:"column:view_form_id;not null" json:"view_form_id"`
	ViewFormFieldId string `gorm:"column:view_form_field_id;not null" json:"view_form_field_id"`
}

type EntityEntityFormView struct {
	Id            string `gorm:"column:id;not null" json:"id"`
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"`
	Name          string `gorm:"column:name;not null" json:"name"`
	Type          string `gorm:"column:type;not null" json:"type"`
	DatasourceId  string `gorm:"column:datasource_id;not null" json:"datasource_id"`
	SubjectId     string `gorm:"column:subject_id;not null" json:"subject_id"`
	Description   string `gorm:"column:description;not null" json:"description"`
}

type EntityEntitySubjectProperty struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	PathId      string `gorm:"column:path_id;not null" json:"path_id"`
	Path        string `gorm:"column:path;not null" json:"path"`
	StandardId  string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type EntityEntitySubjectModel struct {
	Id            string `gorm:"column:id;not null" json:"id"`
	BusinessName  string `gorm:"column:business_name;not null" json:"business_name"`
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"`
	Description   string `gorm:"column:description;not null" json:"description"`
	DataViewId    string `gorm:"column:data_view_id;not null" json:"data_view_id"`
	DeletedAt     string `gorm:"column:deleted_at;not null" json:"deleted_at"`
}

type EntityEntitySubjectModelLabel struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	RelatedModelIds string `gorm:"column:related_model_ids;not null" json:"related_model_ids"`
}

type EntityEntitySubjectEntity struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	PathId      string `gorm:"column:path_id;not null" json:"path_id"`
	Path        string `gorm:"column:path;not null" json:"path"`
}

type EdgeRelationSubjectEntity2SubjectProp struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EntityEntitySubjectObject struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	PathId      string `gorm:"column:path_id;not null" json:"path_id"`
	Path        string `gorm:"column:path;not null" json:"path"`
	RefId       string `gorm:"column:ref_id;not null" json:"ref_id"`
}

type EdgeRelationSubjectObject2SubjectEntity struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EdgeRelationSubjectObject2Self struct {
	Id    string `gorm:"column:id;not null" json:"id"`
	RefId string `gorm:"column:ref_id;not null" json:"ref_id"`
}

type EntityEntitySubjectDomain struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	PathId      string `gorm:"column:path_id;not null" json:"path_id"`
	Path        string `gorm:"column:path;not null" json:"path"`
}

type EdgeRelationSubjectDomain2SubjectObject struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

type EntityEntitySubjectGroup struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	PathId      string `gorm:"column:path_id;not null" json:"path_id"`
	Path        string `gorm:"column:path;not null" json:"path"`
}

type EdgeRelationSubjectGroup2SubjectDomains struct {
	Id       string `gorm:"column:id;not null" json:"id"`
	ParentId string `gorm:"column:parent_id;not null" json:"parent_id"`
}

// 认知搜索-数据目录图谱

type EntityDataCatalog struct {
	Sid              string `gorm:"column:sid;not null" json:"sid"`
	Code             string `gorm:"column:code;not null" json:"code"`
	Name             string `gorm:"column:name;not null" json:"name"`
	Description      string `gorm:"column:description;not null" json:"description"`
	AssetType        int    `gorm:"column:asset_type;not null" json:"asset_type"`
	Color            string `gorm:"column:color;not null" json:"color"`
	DataKind         int    `gorm:"column:data_kind;not null" json:"data_kind"`
	SharedType       int    `gorm:"column:shared_type;not null" json:"shared_type"`
	PublishedAt      int64  `gorm:"column:published_at;not null" json:"published_at"`
	UpdateCycle      int    `gorm:"column:update_cycle;not null" json:"update_cycle"`
	OwnerId          string `gorm:"column:owner_id;not null" json:"owner_id"`
	OwnerName        string `gorm:"column:owner_name;not null" json:"owner_name"`
	DepartmentId     string `gorm:"column:department_id;not null" json:"department_id"`
	Department       string `gorm:"column:department;not null" json:"department"`
	DepartmentPathId string `gorm:"column:department_path_id;not null" json:"department_path_id"`
	DepartmentPath   string `gorm:"column:department_path;not null" json:"department_path"`
	Datasource       string `gorm:"column:datasource;not null" json:"datasource"`
	VesCatalogName   string `gorm:"column:ves_catalog_name;not null" json:"ves_catalog_name"`
	MetadataSchema   string `gorm:"column:metadata_schema;not null" json:"metadata_schema"`
	InfoSystemId     string `gorm:"column:info_system_id;not null" json:"info_system_id"`
	InfoSystemName   string `gorm:"column:info_system_name;not null" json:"info_system_name"`
}

type EdgeCatalogTag2DataCatalog struct {
	TagSid         string `gorm:"column:tag_sid;not null" json:"tag_sid"`
	DataCatalogSid string `gorm:"column:data_catalog_sid;not null" json:"data_catalog_sid"`
}

type EdgeInfoSystem2DataCatalog struct {
	SysSid         string `gorm:"column:sys_sid;not null" json:"sys_sid"`
	DataCatalogSid string `gorm:"column:data_catalog_sid;not null" json:"data_catalog_sid"`
}

type EdgeDepartment2DataCatalog struct {
	DepartmentId   string `gorm:"column:department_id;not null" json:"department_id"`
	DataCatalogSid string `gorm:"column:data_catalog_sid;not null" json:"data_catalog_sid"`
}

type EdgeDataOwner2DataCatalog struct {
	OwnerId        string `gorm:"column:owner_id;not null" json:"owner_id"`
	DataCatalogSid string `gorm:"column:data_catalog_sid;not null" json:"data_catalog_sid"`
}

type EdgeFormView2DataCatalog struct {
	DataCatalogCode string `gorm:"column:datacatalog_code;not null" json:"datacatalog_code"`
	FormViewUuid    string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
}

type EntityCatalogTag struct {
	TagSid  string `gorm:"column:tag_sid;not null" json:"tag_sid"`
	TagCode string `gorm:"column:tag_code;not null" json:"tag_code"`
	TagName string `gorm:"column:tag_name;not null" json:"tag_name"`
}

type EntityInfoSystem struct {
	SysSid                string `gorm:"column:sys_sid;not null" json:"sys_sid"`
	InfoSystemUuid        string `gorm:"column:info_system_uuid;not null" json:"info_system_uuid"`
	InfoSystemName        string `gorm:"column:info_system_name;not null" json:"info_system_name"`
	InfoSystemDescription string `gorm:"column:info_system_description;not null" json:"info_system_description"`
}

type EntityDepartmentV2 struct {
	Id   string `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

type EntityDataOwnerV2 struct {
	OwnerId   string `gorm:"column:owner_id;not null" json:"owner_id"`
	OwnerName string `gorm:"column:owner_name;not null" json:"owner_name"`
}

type EntityFormViewV2 struct {
	FormviewUuid     string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	FormviewCode     string `gorm:"column:formview_code;not null" json:"formview_code"`
	TechnicalName    string `gorm:"column:technical_name;not null" json:"technical_name"`
	BusinessName     string `gorm:"column:business_name;not null" json:"business_name"`
	Type             int    `gorm:"column:type;not null" json:"type"`
	DatasourceId     string `gorm:"column:datasource_id;not null" json:"datasource_id"`
	PublishAt        string `gorm:"column:publish_at;not null" json:"publish_at"`
	OwnerId          string `gorm:"column:owner_id;not null" json:"owner_id"`
	SubjectId        string `gorm:"column:subject_id;not null" json:"subject_id"`
	DepartmentId     string `gorm:"column:department_id;not null" json:"department_id"`
	Description      string `gorm:"column:description;not null" json:"description"`
	OwnerName        string `gorm:"column:owner_name;not null" json:"owner_name"`
	Department       string `gorm:"column:department;not null" json:"department"`
	SubjectName      string `gorm:"column:subject_name;not null" json:"subject_name"`
	DepartmentPathId string `gorm:"column:department_path_id;not null" json:"department_path_id"`
	DepartmentPath   string `gorm:"column:department_path;not null" json:"department_path"`
	SubjectPathId    string `gorm:"column:subject_path_id;not null" json:"subject_path_id"`
	SubjectPath      string `gorm:"column:subject_path;not null" json:"subject_path"`
}

type EdgeDataExploreReport2MetadataTable struct {
	ColumnId      string `gorm:"column:column_id;not null" json:"column_id"`
	FormViewId    string `gorm:"column:form_view_id;not null" json:"form_view_id"`
	ExploreItem   string `gorm:"column:explore_item;not null" json:"explore_item"`
	ExploreResult string `gorm:"column:explore_result;not null" json:"explore_result"`
}

type EdgeMetadataTableField2MetadataTable struct {
	ColumnId     string `gorm:"column:column_id;not null" json:"column_id"`
	FormViewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
}

type EdgeMetadataSchema2MetadataTable struct {
	SchemaSid    string `gorm:"column:schema_sid;not null" json:"schema_sid"`
	FormViewUuid string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
}

type EntityMetadataSchemaV2 struct {
	SchemaSid  string `gorm:"column:schema_sid;not null" json:"schema_sid"`
	SchemaName string `gorm:"column:schema_name;not null" json:"schema_name"`
	PrefixName string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EntityDataSourceV2 struct {
	DataSourceUuid     string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	DataSourceName     string `gorm:"column:data_source_name;not null" json:"data_source_name"`
	DataSourceTypeName string `gorm:"column:data_source_type_name;not null" json:"data_source_type_name"`
	SourceTypeCode     string `gorm:"column:source_type_code;not null" json:"source_type_code"`
	SourceTypeName     string `gorm:"column:source_type_name;not null" json:"source_type_name"`
	PrefixName         string `gorm:"column:prefix_name;not null" json:"prefix_name"`
}

type EdgeDatasource2MetaDataSchema struct {
	DataSourceUuid string `gorm:"column:data_source_uuid;not null" json:"data_source_uuid"`
	SchemaSid      string `gorm:"column:schema_sid;not null" json:"schema_sid"`
}

type EntityDataExploreReportV2 struct {
	ColumnId      string `gorm:"column:column_id;not null" json:"column_id"`
	ExploreItem   string `gorm:"column:explore_item;not null" json:"explore_item"`
	ColumnName    string `gorm:"column:column_name;not null" json:"column_name"`
	ExploreResult string `gorm:"column:explore_result;not null" json:"explore_result"`
}

type EntityFormViewFieldV2 struct {
	ColumnId      string `gorm:"column:column_id;not null" json:"column_id"`
	FormViewUuid  string `gorm:"column:formview_uuid;not null" json:"formview_uuid"`
	TechnicalName string `gorm:"column:technical_name;not null" json:"technical_name"`
	BusinessName  string `gorm:"column:business_name;not null" json:"business_name"`
	DataType      string `gorm:"column:data_type;not null" json:"data_type"`
}

type TableSubjectDomain struct {
	Id   string `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
	Type int    `gorm:"column:type;not null" json:"type"`
}

// 业务架构图谱

type BRGEntityBusinessDomain struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	Owners      string `gorm:"column:owners;not null" json:"owners"`
}

type BRGEdgeEntityBusinessDomain2EntityThemeDomain struct {
	DomainId string `gorm:"column:domain_id;not null" json:"domain_id"`
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
}

type BRGEntityThemeDomain struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	Owners      string `gorm:"column:owners;not null" json:"owners"`
}

type BRGEdgeEntityThemeDomain2EntityBusinessDomain struct {
	DomainId string `gorm:"column:domain_id;not null" json:"domain_id"`
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
}

type BRGEdgeEntityThemeDomain2EntityBusinessObject struct {
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
	ObjectId string `gorm:"column:object_id;not null" json:"object_id"`
}

type BRGEntityBusinessObject struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	Owners      string `gorm:"column:owners;not null" json:"owners"`
}

type BRGEdgeEntityBusinessObject2EntityThemeDomain struct {
	ThemeId  string `gorm:"column:theme_id;not null" json:"theme_id"`
	ObjectId string `gorm:"column:object_id;not null" json:"object_id"`
}

type BRGEdgeEntityBusinessObject2EntityBusinessForm struct {
	BusinessObjectId string `gorm:"column:business_object_id;not null" json:"business_object_id"`
	BusinessFormId   string `gorm:"column:business_form_id;not null" json:"business_form_id"`
}

type BRGEdgeEntityBusinessObject2EntityDataCatalog struct {
	ObjectId  string `gorm:"column:object_id;not null" json:"object_id"`
	CatalogId string `gorm:"column:catalog_id;not null" json:"catalog_id"`
}

type BRGEntityDataCatalog struct {
	Id          string `gorm:"column:id;not null" json:"id"`
	Code        string `gorm:"column:code;not null" json:"code"`
	Title       string `gorm:"column:title;not null" json:"title"`
	GroupId     int    `gorm:"column:group_id;not null" json:"group_id"`
	GroupName   string `gorm:"column:group_name;not null" json:"group_name"`
	ThemeId     string `gorm:"column:theme_id;not null" json:"theme_id"`
	ThemeName   string `gorm:"column:theme_name;not null" json:"theme_name"`
	Description string `gorm:"column:description;not null" json:"description"`
	DataRange   string `gorm:"column:data_range;not null" json:"data_range"`
	UpdateCycle string `gorm:"column:update_cycle;not null" json:"update_cycle"`
	DataKind    string `gorm:"column:data_kind;not null" json:"data_kind"`
	OrgCode     string `gorm:"column:orgcode;not null" json:"orgcode"`
	OrgName     string `gorm:"column:orgname;not null" json:"orgname"`
}

type BRGEdgeEntityDataCatalog2EntityInfoSystem struct {
	CatalogId    string `gorm:"column:catalog_id;not null" json:"catalog_id"`
	InfoSystemId string `gorm:"column:info_system_id;not null" json:"info_system_id"`
}

type BRGEdgeEntityCatalog2EntityBusinessSceneSource struct {
	CatalogId       string `gorm:"column:catalog_id;not null" json:"catalog_id"`
	BusinessSceneId string `gorm:"column:business_scene_id;not null" json:"business_scene_id"`
}

type BRGEdgeEntityCatalog2EntityBusinessSceneRelated struct {
	CatalogId       string `gorm:"column:catalog_id;not null" json:"catalog_id"`
	BusinessSceneId string `gorm:"column:business_scene_id;not null" json:"business_scene_id"`
}

type BRGEdgeEntityDataCatalog2EntityDataCatalogColumn struct {
	ObjectId string `gorm:"column:object_id;not null" json:"object_id"`
	FieldId  string `gorm:"column:field_id;not null" json:"field_id"`
}

type BRGEntityInfoSystem struct {
	Id        string `gorm:"column:id;not null" json:"id"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Path      string `gorm:"column:path;not null" json:"path"`
	Attribute int    `gorm:"column:attribute;not null" json:"attribute"`
}

type BRGEntityBusinessScene struct {
	Id        string `gorm:"column:id;not null" json:"id"`
	Name      string `gorm:"column:name;not null" json:"name"`
	Path      string `gorm:"column:path;not null" json:"path"`
	Attribute int    `gorm:"column:attribute;not null" json:"attribute"`
}

type BRGEntityDataCatalogColumn struct {
	Id            string `gorm:"column:id;not null" json:"id"`
	CatalogId     string `gorm:"column:catalog_id;not null" json:"catalog_id"`
	ColumnName    string `gorm:"column:column_name;not null" json:"column_name"`
	NameCn        int    `gorm:"column:name_cn;not null" json:"name_cn"`
	Description   string `gorm:"column:description;not null" json:"description"`
	DataFormat    string `gorm:"column:data_format;not null" json:"data_format"`
	DataLength    string `gorm:"column:data_length;not null" json:"data_length"`
	DatametaId    string `gorm:"column:datameta_id;not null" json:"datameta_id"`
	DatametaName  string `gorm:"column:datameta_name;not null" json:"datameta_name"`
	Ranges        string `gorm:"column:ranges;not null" json:"ranges"`
	CodesetId     string `gorm:"column:codeset_id;not null" json:"codeset_id"`
	CodesetName   string `gorm:"column:codeset_name;not null" json:"codeset_name"`
	TimestampFlag string `gorm:"column:timestamp_flag;not null" json:"timestamp_flag"`
}

type BRGEntitySourceTable struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
	SchemaName      string `gorm:"column:schema_name;not null" json:"schema_name"`
	VeCatalogId     int    `gorm:"column:ve_catalog_id;not null" json:"ve_catalog_id"`
	MetadataTableId int    `gorm:"column:metadata_table_id;not null" json:"metadata_table_id"`
}

type BRGEdgeEntitySourceTable2EntityDataCatalog struct {
	SourceId string `gorm:"column:source_id;not null" json:"source_id"`
	Code     string `gorm:"column:code;not null" json:"code"`
}

type BRGEntityStandardTable struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
	MetadataTableId int    `gorm:"column:metadata_table_id;not null" json:"metadata_table_id"`
	SchemaName      int    `gorm:"column:schema_name;not null" json:"schema_name"`
	VeCatalogId     int    `gorm:"column:ve_catalog_id;not null" json:"ve_catalog_id"`
}

type BRGEdgeEntityStandardTable2EntityDataCatalog struct {
	StandardId string `gorm:"column:standard_id;not null" json:"standard_id"`
	Code       string `gorm:"column:code;not null" json:"code"`
}

type BRGEntityDepartment struct {
	Id   string `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

type BRGEdgeEntityDepartment2EntityDataCatalog struct {
	CatalogId    string `gorm:"column:catalog_id;not null" json:"catalog_id"`
	DepartmentId string `gorm:"column:department_id;not null" json:"department_id"`
}

type BRGEntityBusinessFormStandard struct {
	Id              string `gorm:"column:id;not null" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	Description     string `gorm:"column:description;not null" json:"description"`
	Guideline       string `gorm:"column:guideline;not null" json:"guideline"`
}

type BRGEdgeEntityBusinessForm2EntityBusinessObject struct {
	BusinessFormId   string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	BusinessObjectId string `gorm:"column:business_object_id;not null" json:"business_object_id"`
}

type BRGEdgeEntityBusinessFormStandard2Self struct {
	BusinessFormIdP string `gorm:"column:business_form_id_p;not null" json:"business_form_id_p"`
	BusinessFormIdC string `gorm:"column:business_form_id_c;not null" json:"business_form_id_c"`
}

type BRGEdgeEntityBusinessForm2EntityStandardTable struct {
	BusinessFormId string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	StandardId     string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type BRGEdgeEntityBusinessForm2EntityBusinessModel struct {
	BusinessFormId  string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
}

type BRGEdgeEntityBusinessFormStandard2EntityField struct {
	BusinessFormId string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	FieldId        string `gorm:"column:field_id;not null" json:"field_id"`
}

type BRGEntityField struct {
	FieldId           string `gorm:"column:field_id;not null" json:"field_id"`
	BusinessFormId    string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	BusinessFormName  string `gorm:"column:business_form_name;not null" json:"business_form_name"`
	Name              string `gorm:"column:name;not null" json:"name"`
	NameEn            string `gorm:"column:name_en;not null" json:"name_en"`
	DataType          string `gorm:"column:data_type;not null" json:"data_type"`
	DataLength        string `gorm:"column:data_length;not null" json:"data_length"`
	ValueRange        string `gorm:"column:value_range;not null" json:"value_range"`
	FieldRelationship string `gorm:"column:field_relationship;not null" json:"field_relationship"`
	RefId             string `gorm:"column:ref_id;not null" json:"ref_id"`
	StandardId        string `gorm:"column:standard_id;not null" json:"standard_id"`
}

type BRGEdgeEntityField2EntityBusinessIndicator struct {
	IndicatorId string `gorm:"column:indicator_id;not null" json:"indicator_id"`
	FieldId     string `gorm:"column:field_id;not null" json:"field_id"`
}

type BRGEntityBusinessModel struct {
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	MainBusinessId  string `gorm:"column:main_business_id;not null" json:"main_business_id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
}

type BRGEdgeEntityBusinessModel2EntityBusinessForm struct {
	BusinessFormId  string `gorm:"column:business_form_id;not null" json:"business_form_id"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
}

type BRGEdgeEntityBusinessModel2EntityDepartment struct {
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	DepartmentId    string `gorm:"column:department_id;not null" json:"department_id"`
}

type BRGEdgeEntityBusinessModel2EntityFlowchart struct {
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	FlowchartId     string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
}

type BRGEdgeEntityBusinessModel2EntityBusinessIndicator struct {
	BusinessModelId   string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	BusinessIndicator string `gorm:"column:business_indicator;not null" json:"business_indicator"`
}

type BRGEntityBusinessIndicator struct {
	BusinessIndicatorId string `gorm:"column:business_indicator_id;not null" json:"business_indicator_id"`
	IndicatorId         string `gorm:"column:indicator_id;not null" json:"indicator_id"`
	BusinessModelId     string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	Name                string `gorm:"column:name;not null" json:"name"`
	Desc                string `gorm:"column:desc;not null" json:"desc"`
}

type BRGEdgeEntityBusinessIndicator2EntityBusinessModel struct {
	BusinessModelId   string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	BusinessIndicator string `gorm:"column:business_indicator;not null" json:"business_indicator"`
}

type BRGEntityFlowchart struct {
	FlowchartId     string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	Name            string `gorm:"column:name;not null" json:"name"`
	Description     string `gorm:"column:description;not null" json:"description"`
	BusinessModelId string `gorm:"column:business_model_id;not null" json:"business_model_id"`
	MainBusinessId  string `gorm:"column:main_business_id;not null" json:"main_business_id"`
}

type BRGEdgeEntityFlowchart2EntityFlowchart struct {
	FlowchartIdP string `gorm:"column:flowchart_id_p;not null" json:"flowchart_id_p"`
	FlowchartIdC string `gorm:"column:flowchart_id_c;not null" json:"flowchart_id_c"`
}

type BRGEdgeEntityFlowchart2EntityFlowchartNode struct {
	FlowchartId string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	NodeId      string `gorm:"column:node_id;not null" json:"node_id"`
}

type BRGEntityFlowchartNode struct {
	NodeId      string `gorm:"column:node_id;not null" json:"node_id"`
	FlowchartId string `gorm:"column:flowchart_id;not null" json:"flowchart_id"`
	DiagramId   string `gorm:"column:diagram_id;not null" json:"diagram_id"`
	DiagramName string `gorm:"column:diagram_name;not null" json:"diagram_name"`
	Name        string `gorm:"column:name;not null" json:"name"`
	Description string `gorm:"column:description;not null" json:"description"`
	Target      string `gorm:"column:target;not null" json:"target"`
	Source      string `gorm:"column:source;not null" json:"source"`
}
type BRGEdgeEntityFlowchartNode2EntityFlowchartNode struct {
	NodeIdP string `gorm:"column:node_id_p;not null" json:"node_id_p"`
	NodeIdC string `gorm:"column:node_id_c;not null" json:"node_id_c"`
}
