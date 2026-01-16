package knowledge_network

import "context"

type SearchEngineResponse struct {
	Res struct {
		RecommdedDAG      any `json:"RecommdedDAG,omitempty"`
		RecommendCodeDAG  any `json:"RecommendCodeDAG,omitempty"`
		GraphSynSearchDAG any `json:"GraphSynSearchDAG"`
	} `json:"res"`
}

type TableReq struct {
	MainBusinessID string     `json:"main_business_ID"`
	Table          *AdTable   `json:"table"`
	Fields         []*AdField `json:"fields"`
	Key            string     `json:"key"`
}

type AdTable struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//DataRange    string   `json:"data_range"`
	Guideline    string   `json:"guideline"`
	SourceSystem []string `json:"source_system"`
	ResourceTag  []string `json:"resource_tag"`
}

type AdField struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	NameEn         string `json:"name_en"`
	StandardId     string `json:"standard_id"`
	CodeTable      string `json:"code_table"`
	DataAccuracy   int    `json:"data_accuracy"`
	DataLength     int    `json:"data_length"`
	DataType       string `json:"data_type"`
	EncodingRule   string `json:"encoding_rule"`
	Explanation    string `json:"explanation"`
	FormulateBasis string `json:"formulate_basis"`
}

type TableResp struct {
	Tables []struct {
		ID       string  `json:"id"`
		HitScore float64 `json:"hit_score"`
		Reason   string  `json:"reason"`
	} `json:"tables"`
}

type CheckFieldsReq struct {
	TableId     string    `json:"table_id"`
	TableName   string    `json:"table_name"`
	TableFields []*Fields `json:"table_fields"`
}

type Fields struct {
	FieldId         string `json:"field_id"`
	TableFieldName  string `json:"table_field_name"`
	StandardId      string `json:"standard_id"`
	TableFieldType  string `json:"table_field_type"`
	TableFieldRange string `json:"table_field_range"`
}

type CheckFieldsResp struct {
	TableId           string               `json:"table_id"`
	FieldsCheckResult []*FieldsCheckResult `json:"fields_check_result"`
}

type FieldsCheckResult struct {
	FieldId      string          `json:"field_id"`
	StandardId   string          `json:"standard_id"`
	Consistent   []*TableField   `json:"consistent"`
	Inconsistent []*Inconsistent `json:"inconsistent"`
}

type Inconsistent struct {
	StandardId string        `json:"standard_id"`
	Fields     []*TableField `json:"fields"`
}

type TableField struct {
	TableId string `json:"table_id"`
	FieldId string `json:"field_id"`
}

type FlowReq struct {
	BusinessModelId string     `json:"business_model_id"`
	Node            Node       `json:"node"`
	ParentNode      ParentNode `json:"parent_node"`
	Flowchart       Flowchart  `json:"flowchart"`
}

type FlowResp struct {
	Flowcharts []*ADRecResp `json:"flowcharts"`
}

type ADRecResp struct {
	Id       string  `json:"id"`
	HitScore float64 `json:"hit_score"`
	Reason   string  `json:"reason"`
}

type Node struct {
	Id          string `json:"id"`
	MxcellId    string `json:"mxcell_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ParentNode struct {
	Node
	Id     string   `json:"id"`
	Tables []string `json:"tables"`
}

type Flowchart struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nodes       []*Node `json:"nodes"`
}

type FieldStandardizationReq struct {
	TableName   string `json:"table_name"`
	TableFields []struct {
		TableFieldName string `json:"table_field_name"`
		TableFieldType string `json:"table_field_type"`
	} `json:"table_fields"`
}

type FieldStandardizationResp struct {
	TableFields []struct {
		RecStds []struct {
			Info      string `json:"info"`
			Rank      int    `json:"rank"`
			RecReason string `json:"rec_reason"`
			StdChName string `json:"std_ch_name"`
			StdCode   int64  `json:"std_code"`
		} `json:"rec_stds"`
		TableFieldName string `json:"table_field_name"`
	} `json:"table_fields"`
	TableName string `json:"table_name"`
}

type GraphAnalysisResp map[string]any

//SearchEngine   自定义认知服务，表单推荐等接口
func (a *ad) SearchEngine(ctx context.Context, serviceId string, content any) (*SearchEngineResponse, error) {
	uri := a.baseUrl + "/api/search-engine/v1/open/custom-services/" + serviceId
	headers := make(map[string][]string)
	headers["source-type"] = []string{"2"}
	return httpPostDo[SearchEngineResponse](ctx, uri, content, headers, a)
}
