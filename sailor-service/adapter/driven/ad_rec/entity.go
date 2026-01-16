package ad_rec

type PtReq struct {
	Content string `json:"content"`
}

type PtResp struct {
	Res struct {
		Response string `json:"response"`
	} `json:"res"`
}

type CPRecommendCodeReq struct {
	TableId       string `json:"table_id"`
	TableName     string `json:"table_name"`
	TableDesc     string `json:"table_desc"`
	DepartmentIDS string `json:"department_ids"`
	TableFields   []struct {
		TableFieldId   string `json:"table_field_id"`
		TableFieldName string `json:"table_field_name"`
	} `json:"table_fields"`
}

type CPRecommendCodeResp struct {
	Res struct {
		Answers struct {
			TableName   string `json:"table_name"`
			TableFields []struct {
				TableFieldName string `json:"table_field_name"`
				RecStds        []struct {
					StdChName string  `json:"std_ch_name"`
					StdCode   string  `json:"std_code"`
					Score     float64 `json:"score"`
				} `json:"rec_stds"`
			} `json:"table_fields"`
		} `json:"answers"`
	} `json:"res"`
}

type CPRecommendTableReq struct {
	BusinessModelId   string `json:"business_model_id"`
	BusinessModelName string `json:"business_model_name"`

	Domain struct {
		DomainId     string `json:"domain_id"`
		DomainName   string `json:"domain_name"`
		DomainPath   string `json:"domain_path"`
		DomainPathId string `json:"domain_path_id"`
	} `json:"domain"`

	Dept struct {
		DeptId     string `json:"dept_id"`
		DeptName   string `json:"dept_name"`
		DeptPath   string `json:"dept_path"`
		DeptPathId string `json:"dept_path_id"`
	} `json:"dept"`

	InfoSystem []InfoSystemItem `json:"info_system"`

	Table struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"table"`

	Fields []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"fields"`

	Key string `json:"key"`
}

type InfoSystemItem struct {
	InfoSystemId   string `json:"info_system_id"`
	InfoSystemName string `json:"info_system_name"`
	InfoSystemDesc string `json:"info_system_desc"`
}

type CPRecommendTableResp struct {
	Res struct {
		Answers struct {
			Tables []struct {
				ID       string  `json:"id"`
				HitScore float64 `json:"hit_score"`
				Reason   string  `json:"reason"`
			} `json:"tables"`
		} `json:"answers"`
	} `json:"res"`
}

type CPRecommendFlowReq struct {
	BusinessModelId string `json:"business_model_id"`
	Node            struct {
		Id          string `json:"id"`
		MxcellId    string `json:"mxcell_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"node"`
	ParentNode struct {
		Id          string   `json:"id"`
		MxcellId    string   `json:"mxcell_id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Tables      []string `json:"tables"`
	} `json:"parent_node"`
	Flowchart struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		BusinessModelId string `json:"business_model_id"`
		Nodes           []struct {
			Id          string `json:"id"`
			MxcellId    string `json:"mxcell_id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"nodes"`
	} `json:"flowchart"`
}

type CPRecommendFlowResp struct {
	Res struct {
		Answers struct {
			Flowcharts []struct {
				Id       string  `json:"id"`
				HitScore float64 `json:"hit_score"`
				Reason   string  `json:"reason"`
			} `json:"flowcharts"`
		} `json:"answers"`
	} `json:"res"`
}

type CPRecommendCheckCodeReq struct {
	Data []struct {
		BusinessModelId   string `json:"business_model_id"`
		BusinessModelName string `json:"business_model_name"`
		Domain            struct {
			DomainId     string `json:"domain_id"`
			DomainName   string `json:"domain_name"`
			DomainPath   string `json:"domain_path"`
			DomainPathId string `json:"domain_path_id"`
		} `json:"domain"`
		Dept struct {
			DeptId     string `json:"dept_id"`
			DeptName   string `json:"dept_name"`
			DeptPath   string `json:"dept_path"`
			DeptPathId string `json:"dept_path_id"`
		} `json:"dept"`
		InfoSystem []struct {
			InfoSystemId   string `json:"info_system_id"`
			InfoSystemName string `json:"info_system_name"`
			InfoSystemDesc string `json:"info_system_desc"`
		} `json:"info_system"`
		Tables []struct {
			TableId     string `json:"table_id"`
			TableName   string `json:"table_name"`
			TableDesc   string `json:"table_desc"`
			TableFields []struct {
				FieldId        string `json:"field_id"`
				TableFieldName string `json:"table_field_name"`
				TableFieldDesc string `json:"table_field_desc"`
				StandardId     string `json:"standard_id"`
			} `json:"table_fields"`
		} `json:"tables"`
	} `json:"data"`
}

type CPRecommendViewReq struct {
	RecommendViewTypes []int `json:"recommend_view_types"`
	Table              struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"table"`
	Fields []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"fields"`
}

type CPTestLLMReq struct {
}
type CConditions struct {
	Index  string `json:"index"`
	Title  string `json:"title"`
	Schema string `json:"schema"`
	Source string `json:"source"`
}

type CPText2SqlReq struct {
	Query  string        `json:"query"`
	Search []CConditions `json:"search"`
}

type CPRecommendCheckCodeResp struct {
	Res struct {
		Answers []struct {
			TableId           string `json:"table_id"`
			FieldsCheckResult []struct {
				FieldId    string `json:"field_id"`
				StandardId string `json:"standard_id"`
				Consistent []struct {
					TableId string `json:"table_id"`
					FieldId string `json:"field_id"`
				} `json:"consistent"`
				Inconsistent []struct {
					StandardId string `json:"standard_id"`
					Fields     []struct {
						TableId string `json:"table_id"`
						FieldId string `json:"field_id"`
					} `json:"fields"`
				} `json:"inconsistent"`
			} `json:"fields_check_result"`
		} `json:"answers"`
	} `json:"res"`
}

type CPRecommendAssetSearchReq struct {
	Query string `json:"query"  binding:"required"`
	Limit int    `json:"limit" binding:"max=100"` // 每次返回的数量
	//智能搜索对象，停用词
	Stopwords    []string `json:"stopwords" binding:"omitempty,unique" example:"应用"`
	StopEntities []string `json:"stop_entities" binding:"omitempty,unique" example:"datacatalog"` //智能搜索维度, 停用的实体类别
	Filter       struct {
		AssetType       string           `json:"asset_type"`
		DataKind        string           `json:"data_kind"`
		UpdateCycle     string           `json:"update_cycle"`
		SharedType      string           `json:"shared_type"`
		StartTime       string           `json:"start_time"`
		EndTime         string           `json:"end_time"`
		StopEntityInfos []StopEntityInfo `json:"stop_entity_infos" binding:"omitempty"`
	} `json:"filter"`

	//AdAppid        string `json:"ad_appid"`
	//KgId           int    `json:"kg_id"`
	//Entity2Service struct {
	//} `json:"entity2service"`
	//RequiredResource struct {
	//} `json:"required_resource"`
}

type StopEntityInfo struct {
	ClassName string   `json:"class_name"`
	Names     []string `json:"names"`
}

type CPRecommendAssetSearchResp struct {
	Res struct {
		Answers []struct {
			TableId           string `json:"table_id"`
			FieldsCheckResult []struct {
				FieldId    string `json:"field_id"`
				StandardId string `json:"standard_id"`
				Consistent []struct {
					TableId string `json:"table_id"`
					FieldId string `json:"field_id"`
				} `json:"consistent"`
				Inconsistent []struct {
					StandardId string `json:"standard_id"`
					Fields     []struct {
						TableId string `json:"table_id"`
						FieldId string `json:"field_id"`
					} `json:"fields"`
				} `json:"inconsistent"`
			} `json:"fields_check_result"`
		} `json:"answers"`
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
