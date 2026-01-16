package client

type RecommdedDAG struct {
	Outputs struct {
		GraphRecommendAnswerList any `json:"graph_recommend_answer_list"`
	} `json:"outputs"`
	BlackBoard struct {
		RecommedCodeExecutor struct {
			ExecuteTime float64 `json:"execute_time"`
			KeyProcess  string  `json:"key_process"`
		} `json:"RecommedCodeExecutor"`
	} `json:"black_board"`
}

type RecCodeReq struct {
	TableId       string `json:"table_id"`
	TableName     string `json:"table_name"`
	TableDesc     string `json:"table_desc"`
	DepartmentIDS string `json:"department_ids"`
	TableFields   []struct {
		TableFieldId   string `json:"table_field_id"`
		TableFieldName string `json:"table_field_name"`
		//TableFieldType string `json:"table_field_type"`
	} `json:"table_fields"`
}

type RecCodeResp struct {
	TableFields []struct {
		RecStds []struct {
			Info      string `json:"info"`
			Rank      int    `json:"rank"`
			RecReason string `json:"rec_reason"`
			StdChName string `json:"std_ch_name"`
			StdCode   any    `json:"std_code"`
		} `json:"rec_stds"`
		TableFieldName string `json:"table_field_name"`
	} `json:"table_fields"`
	TableName string `json:"table_name"`
}
