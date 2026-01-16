package client

import (
	"encoding/json"
	"fmt"

	"github.com/samber/lo"
)

const (
	AssetDataCatalog  = "data_catalog"
	AssetInterfaceSvc = "interface_svc"
)

type AssetSearch struct {
	AssetType []string `json:"asset_type" binding:"omitempty,unique,dive,oneof=data_catalog interface_svc logical_view all"` // 资产分类
	Query     string   `json:"query"  binding:"required"`                                                                    // 输入的查询内容
	Limit     int      `json:"limit" binding:"max=100"`                                                                      // 每次返回的数量
	MaxLimit  int      `json:"-"`                                                                                            //从AD返回的最大数量
	LastId    string   `json:"last_id"  binding:"omitempty"`                                                                 //最后一个结果的ID
	LastScore float64  `json:"last_score" binding:"gte=0"`                                                                   //最后一个结果的分数
	//属性相关的
	DataKind              []int            `json:"data_kind" binding:"omitempty,max=10,unique,dive,gt=0" example:"1,2"`              // 基础信息分类 1 人 2 地 4 事 8 物 16 组织 32 其他  可组合，如 人和地 即 1|2 = 3
	UpdateCycle           []int            `json:"update_cycle" binding:"omitempty,max=10,unique,dive,gt=0" example:"3,7"`           // 更新频率 参考数据字典：GXZQ，1不定时 2实时 3每日 4每周 5每月 6每季度 7每半年 8每年 9其他
	SharedType            []int            `json:"shared_type" binding:"omitempty,max=10,unique,dive,gt=0" example:"2"`              // 共享属性 1 无条件共享 2 有条件共享 3 不予共享
	StartTime             *int64           `json:"start_time" binding:"omitempty,gte=0,ltfield=EndTime" example:"1682586655"`        // 开始时间，秒时间戳
	EndTime               *int64           `json:"end_time" binding:"required_with=StartTime,omitempty,gte=0"  example:"1682586655"` // 结束时间，秒时间戳
	Stopwords             []string         `json:"stopwords" binding:"omitempty,unique" example:"应用"`                                //智能搜索对象，停用词
	StopEntities          []string         `json:"stop_entities" binding:"omitempty,unique" example:"datacatalog"`                   //智能搜索维度, 停用的实体类别
	StopEntityInfos       []StopEntityInfo `json:"stop_entity_infos" binding:"omitempty"`                                            //智能搜索维度, 停用的实体点
	DepartmentId          []string         `json:"department_id,omitempty"`
	SubjectDomainId       []string         `json:"subject_domain_id,omitempty"`
	DataOwnerId           []string         `json:"data_owner_id,omitempty"`
	InfoSystemId          []string         `json:"info_system_id,omitempty"`
	AvailableOption       int              `json:"available_option" binding:"omitempty,gte=0,lte=2"`
	SearchType            string           `json:"search_type" binding:"default=cognitive_search"`
	PublishStatusCategory []string         `json:"publish_status_category,omitempty"`
	OnlineStatus          []string         `json:"online_status,omitempty"`
	PublishStatus         []string         `json:"publish_status,omitempty"`
	CateNodeId            []CateNodeItem   `json:"cate_node_id,omitempty"`
	SubjectId             []string         `json:"subject_id,omitempty"`
	ResourceType          []string         `json:"resource_type,omitempty"`
}

type CateNodeItem struct {
	CategoryId  string   `json:"category_id"`
	SelectedIds []string `json:"selected_ids"`
}

type StopEntityInfo struct {
	ClassName string   `json:"class_name"`
	Names     []string `json:"names"`
}

type AssistantQa struct {
	Stream       bool     `json:"stream"`
	Query        string   `json:"query"`
	User         string   `json:"user"`
	Limit        int      `json:"limit"`
	StopWords    []string `json:"stopwords"`
	StopEntities []string `json:"stop_entities"`
	Filter       struct {
		AssetType       string           `json:"asset_type"`                            // 资产分类
		DataKind        string           `json:"data_kind,omitempty"`                   // 基础信息分类
		UpdateCycle     string           `json:"update_cycle"`                          // 更新频率
		SharedType      string           `json:"shared_type"`                           // 共享条件
		StartTime       string           `json:"start_time"`                            // 开始时间，毫秒时间戳
		EndTime         string           `json:"end_time"`                              // 结束时间，毫秒时间戳
		StopEntityInfos []StopEntityInfo `json:"stop_entity_infos" binding:"omitempty"` // 这个参数不用户图分析，弄错位置了，不改了
	} `json:"filter"`
}

type AssistantQaResp struct {
	Res struct {
	} `json:"res"`
}

type SSEReq struct {
	//Query string `json:"query"`
}

func assetTypeCode(asset string) int {
	if asset == AssetInterfaceSvc {
		return 2
	}
	return 1
}
func transferAssetSlice(codes []string) (assets []int) {
	for _, code := range codes {
		assets = append(assets, assetTypeCode(code))
	}
	return assets
}

func (a *AssetSearch) Init() {
	if a.Limit <= 0 {
		a.Limit = 3
	}
	if a.AssetType == nil {
		a.AssetType = make([]string, 0)
	}
	if a.DataKind == nil {
		a.DataKind = make([]int, 0)
	}
	if a.UpdateCycle == nil {
		a.UpdateCycle = make([]int, 0)
	}
	if a.SharedType == nil {
		a.SharedType = make([]int, 0)
	}
	if a.Stopwords == nil {
		a.Stopwords = make([]string, 0)
	}
	if a.StopEntities == nil {
		a.StopEntities = make([]string, 0)
	}
	if a.StopEntityInfos == nil {
		a.StopEntityInfos = make([]StopEntityInfo, 0)
	}

	if a.DepartmentId == nil {
		a.DepartmentId = make([]string, 0)
	}
	if a.SubjectDomainId == nil {
		a.SubjectDomainId = make([]string, 0)
	}
	if a.DataOwnerId == nil {
		a.DataOwnerId = make([]string, 0)
	}
	if a.InfoSystemId == nil {
		a.InfoSystemId = make([]string, 0)
	}
	if a.PublishStatusCategory == nil {
		a.PublishStatusCategory = make([]string, 0)
	}
	if a.OnlineStatus == nil {
		a.OnlineStatus = make([]string, 0)
	}
	if a.CateNodeId == nil {
		a.CateNodeId = make([]CateNodeItem, 0)
	}
	if a.SubjectId == nil {
		a.SubjectId = make([]string, 0)
	}
	if a.ResourceType == nil {
		a.ResourceType = make([]string, 0)
	}

	//if a.PublishStatus == nil {
	//	a.PublishStatus = make([]string, 0)
	//}
}

func (a *AssetSearch) GenAssetSearchADRequest() *AssetSearchADRequest {
	reqData := &AssetSearchADRequest{
		Query:        a.Query,
		Limit:        a.MaxLimit,
		Stopwords:    a.Stopwords,
		StopEntities: a.StopEntities,
		LastScore:    a.LastScore,
	}
	reqData.Filter.StopEntityInfos = a.StopEntityInfos
	reqData.Filter.DataKind = fmt.Sprintf("%v", ArrayToInt(a.DataKind))
	reqData.Filter.AssetType = string(lo.T2(json.Marshal(transferAssetSlice(a.AssetType))).A)
	if len(a.AssetType) == 0 {
		reqData.Filter.AssetType = "[-1]"
	}
	reqData.Filter.UpdateCycle = string(lo.T2(json.Marshal(a.UpdateCycle)).A)
	if len(a.UpdateCycle) == 0 {
		reqData.Filter.UpdateCycle = "[-1]"
	}
	reqData.Filter.SharedType = string(lo.T2(json.Marshal(a.SharedType)).A)
	if len(a.SharedType) == 0 {
		reqData.Filter.SharedType = "[-1]"
	}
	reqData.Filter.StartTime = "0"
	if a.StartTime != nil {
		reqData.Filter.StartTime = fmt.Sprintf("%d", *a.StartTime)
	}
	reqData.Filter.EndTime = "0"
	if a.EndTime != nil {
		reqData.Filter.EndTime = fmt.Sprintf("%d", *a.EndTime)
	}
	return reqData
}

// AssetSearchADRequest 资产认知搜索传给AD的结构
type AssetSearchADRequest struct {
	Query        string   `json:"query"`         // 输入的查询内容
	Limit        int      `json:"limit"`         // 每次返回的数量
	Stopwords    []string `json:"stopwords"`     // 停用词
	StopEntities []string `json:"stop_entities"` // 智能搜索维度, 停用的实体
	LastScore    float64  `json:"last_score"`    // 最后一个结果的分数
	//下面的是传给图分析服务的，全部使用字符串传
	Filter struct {
		AssetType       string           `json:"asset_type"`                            // 资产分类
		DataKind        string           `json:"data_kind,omitempty"`                   // 基础信息分类
		UpdateCycle     string           `json:"update_cycle"`                          // 更新频率
		SharedType      string           `json:"shared_type"`                           // 共享条件
		StartTime       string           `json:"start_time"`                            // 开始时间，毫秒时间戳
		EndTime         string           `json:"end_time"`                              // 结束时间，毫秒时间戳
		StopEntityInfos []StopEntityInfo `json:"stop_entity_infos" binding:"omitempty"` // 这个参数不用户图分析，弄错位置了，不改了
	} `json:"filter"`
}

type AssetSearchEntity struct {
	VID             string `json:"vid"`
	Type            string `json:"type"`
	DataCatalogName string `json:"datacatalogname"`
	DataCatalogId   string `json:"datacatalogid"`
	DescriptionName string `json:"description_name,omitempty"`
	Description     string `json:"description"`
	AssetType       string `json:"asset_type"`
	Code            string `json:"code"`
	MetadataSchema  string `json:"metadata_schema"`
	Datasource      string `json:"datasource"`
	DataOwner       string `json:"data_owner"`
	PublishedAt     string `json:"published_at"`
	DepartmentId    string `json:"department_id"`
	Department      string `json:"department"`
	DepartmentPath  string `json:"department_path"`
	ResourceId      string `json:"resourceid"`
	ResourceName    string `json:"resourcename"`
	OwnerID         string `json:"owner_id"`
	OwnerName       string `json:"owner_name"`
	SubjectId       string `json:"subject_id"`
	SubjectName     string `json:"subject_name"`
	SubjectPath     string `json:"subject_path"`
	InfoSystemId    string `json:"info_system_id"`
	InfoSystemName  string `json:"info_system_name"`

	TechnicalName string `json:"technical_name"`
	PublishStatus string `json:"publish_status"`
	OnlineAt      string `json:"online_at"`
	OnlineStatus  string `json:"online_status"`
	IsPermissions string `json:"is_permissions"`

	FormViewUuid string `json:"formview_uuid"`
	FormViewCode string `json:"formview_code"`
	BusinessName string `json:"business_name"`
	SubjectNodes string `json:"subject_nodes"`
	OnlineTime   string `json:"online_time"`
	ResourceType string `json:"resource_type"`
}

type AssetSearchAnswerEntity struct {
	SearchStartInfos []SearchStartInfo          `json:"starts"`
	Subgraph         GraphSynSearchSubgraphPath `json:"subgraph"`
	Entity           AssetSearchEntity          `json:"entity"`
	Score            float64                    `json:"score"`
	TotalKeys        []string                   `json:"total_keys"`
}

type SearchStartInfo struct {
	//Synonyms []struct {
	//	Source  string   `json:"source"`
	//	Synonym []string `json:"synonym"`
	//} `json:"synonyms"`
	ClassName string `json:"class_name"`
	Alias     string `json:"alias"`
	Name      string `json:"name"`
	Hit       struct {
		Prop  string   `json:"prop"`
		Value string   `json:"value"`
		Keys  []string `json:"keys"`
		Alias string   `json:"alias"`
	} `json:"hit"`
}

type AssetSearchData struct {
	Entities        []AssetSearchAnswerEntity `json:"entities"`
	QueryCuts       []QueryCut                `json:"query_cuts"`
	WordCountInfos  []WordCountInfo           `json:"word_count_infos"`
	ClassCountInfos []ClassCountInfo          `json:"class_count_infos"`
	Total           int                       `json:"total"`
}

type WordCountInfo struct {
	Word      string `json:"word"`
	IsSynonym bool   `json:"isSynonym"`
	Count     int    `json:"count"`
}

type ClassCountInfo struct {
	ClassName        string `json:"class_name"`
	Alias            string `json:"alias"`
	Count            int    `json:"count"`
	EntityCounter    *Counter
	EntityCountInfos []EntityCountInfo `json:"entity_count_infos"`
}

type EntityCountInfo struct {
	Alias string `json:"alias"`
	Count int    `json:"count"`
}

type AssetSearchResp struct {
	Data AssetSearchData `json:"data"`
}

type GraphSynSearchDAG struct {
	Outputs GraphSynSearchDAGOutputs `json:"outputs"`
}
type GraphSynSearchDAGOutputs struct {
	Count     int                          `json:"count"`
	Entities  []EntityObj                  `json:"entities"`
	Answer    string                       `json:"answer"`
	Subgraphs []GraphSynSearchSubgraphPath `json:"subgraphs"`
	QueryCuts []QueryCut                   `json:"query_cuts"`
}
type GraphSynSearchSubgraphPath struct {
	Starts []string `json:"starts"`
	End    string   `json:"end"`
}
type QueryCut struct {
	Source     string   `json:"source"`
	Synonym    []string `json:"synonym"`
	IsStopword bool     `json:"is_stopword"`
}

type EntityObj struct {
	SearchStartInfos []SearchStartInfo `json:"starts"`
	Entity           Entity            `json:"entity"`
	Score            float64           `json:"score"`
	IsPermissions    string            `json:"is_permissions"`
}

type Entity struct {
	Id         string   `json:"id"`
	Tags       []string `json:"tags"`
	Properties []struct {
		Tag   string `json:"tag"`
		Props []struct {
			Name  string `json:"name"`
			Alias string `json:"alias"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"props"`
	} `json:"properties"`
	Type  string  `json:"type"`
	Color string  `json:"color"`
	Alias string  `json:"alias"`
	Score float64 `json:"score"`
}
