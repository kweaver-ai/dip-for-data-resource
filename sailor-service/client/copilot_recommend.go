package client

type PtTableReq struct {
	Content string `json:"content"`
}

type PtTableResp struct {
	Res struct {
		Response string `json:"response"`
	} `json:"res"`
}

type CpRecommendCodeReq struct {
	Query struct {
		TableName   string `json:"table_name"`
		TableFields []struct {
			TableFieldName string `json:"table_field_name"`
		} `json:"table_fields"`
	} `json:"query"`
}

type CpRecommendCodeResp struct {
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

type CpRecommendTableReq struct {
	AfQuery struct {
		MainBusinessId string `json:"main_business_id"`
		Table          struct {
			Name         string   `json:"name"`
			Description  string   `json:"description"`
			DataRange    string   `json:"data_range"`
			Guideline    string   `json:"guideline"`
			ResourceTag  string   `json:"resource_tag"`
			SourceSystem []string `json:"source_system"`
		} `json:"table"`
		Fields []struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			NameEn         string `json:"name_en"`
			StdCode        string `json:"std_code"`
			CodeTable      string `json:"code_table"`
			DataAccuracy   int    `json:"data_accuracy"`
			DataLength     int    `json:"data_length"`
			DataType       string `json:"data_type"`
			EncodingRule   string `json:"encoding_rule"`
			FormulateBasis string `json:"formulate_basis"`
		} `json:"fields"`
		Key string `json:"key"`
	} `json:"af_query"`
}

type CpRecommendTableResp struct {
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

type CpRecommendFlowReq struct {
	AfQuery struct {
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
	} `json:"af_query"`
}

type CpRecommendFlowResp struct {
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

type CpRecommendCheckCodeReq struct {
	CheckAfQuery []struct {
		TableId     string `json:"table_id"`
		TableName   string `json:"table_name"`
		TableFields []struct {
			FieldId         string `json:"field_id"`
			TableFieldName  string `json:"table_field_name"`
			StandardId      string `json:"standard_id"`
			TableFieldType  string `json:"table_field_type"`
			TableFieldRange string `json:"table_field_range"`
		} `json:"table_fields"`
	} `json:"check_af_query"`
}

type CpRecommendCheckCodeResp struct {
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

type CpRecommendCodeDAG struct {
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

type CpRecommendTableDAG struct {
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

type CpRecommendFlowDAG struct {
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

type CpRecommendCheckCodeDAG struct {
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

type CpRecommendViewDAG struct {
	Res struct {
		Answers struct {
			Views []struct {
				Id       string  `json:"id"`
				HitScore float64 `json:"hit_score"`
				Reason   string  `json:"reason"`
			} `json:"views"`
		} `json:"answers"`
	} `json:"res"`
}

type CpTestLLM struct {
	Res bool `json:"res"`
}

//type CpRecommendAssetSearchDAG struct {
//	Res struct {
//		Count    int `json:"count"`
//		Entities []struct {
//			Starts []struct {
//				Relation  string `json:"relation"`
//				ClassName string `json:"class_name"`
//				Name      string `json:"name"`
//				Hit       struct {
//					Prop  string   `json:"prop"`
//					Value string   `json:"value"`
//					Keys  []string `json:"keys"`
//				} `json:"hit"`
//				Alias string `json:"alias"`
//			} `json:"starts"`
//			Entity struct {
//				Id              string `json:"id"`
//				Alias           string `json:"alias"`
//				Color           string `json:"color"`
//				ClassName       string `json:"class_name"`
//				Icon            string `json:"icon"`
//				DefaultProperty struct {
//					Name  string `json:"name"`
//					Value string `json:"value"`
//					Alias string `json:"alias"`
//				} `json:"default_property"`
//
//				Tags       []string `json:"tags"`
//				Properties []struct {
//					Tag   string `json:"tag"`
//					Props []struct {
//						Name     string `json:"name"`
//						Value    string `json:"value"`
//						Alias    string `json:"alias"`
//						Type     string `json:"type"`
//						Disabled int    `json:"disabled"`
//						Checked  int    `json:"checked"`
//					} `json:"props"`
//				} `json:"properties"`
//				Score float64 `json:"score"`
//			} `json:"entity"`
//			Score float64 `json:"float"`
//		} `json:"entities"`
//
//		Answer    string `json:"answer"`
//		Subgraphs []struct {
//			Starts []string `json:"starts"`
//			End    string   `json:"end"`
//		} `json:"subgraphs"`
//		QueryCuts []struct {
//			Source     string   `json:"source"`
//			Synonym    []string `json:"synonym"`
//			IsStopword bool     `json:"is_stopword"`
//		} `json:"query_cuts"`
//	} `json:"res"`
//}
