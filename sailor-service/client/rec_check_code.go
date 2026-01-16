package client

type CheckCodeReq struct {
	Data []struct {
		TableId     string `json:"table_id"`
		TableName   string `json:"table_name"`
		TableFields []struct {
			FieldId         string `json:"field_id"`
			TableFieldName  string `json:"table_field_name"`
			StandardId      string `json:"standard_id"`
			TableFieldType  string `json:"table_field_type"`
			TableFieldRange string `json:"table_field_range"`
		} `json:"table_fields"`
	} `json:"data"`
}

type CheckCodeResp struct {
	Data []struct {
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
	} `json:"data"`
}

type CheckCodeReqV2 struct {
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

type TestLLMResp struct {
	Res bool `json:"res"`
}
