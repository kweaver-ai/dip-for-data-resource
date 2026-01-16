package client

type GraphFullTextReq struct {
	KgType string `json:"kg_type"` // 图谱类型
	KgId   string `json:"kg_id"`
	Query  string `json:"query"`
	//Page         int    `json:"page"`
	//Size         int    `json:"size"`
	//MatchingRule string `json:"matching_rule"`
	//MatchingNum  int    `json:"matching_num"`
	SearchConfig []SearchConfigItem `json:"search_config"`
}

type SearchConfigItem struct {
	Tag        string        `json:"tag"`
	Properties []*SearchProp `json:"properties"`
}

type SearchProp struct {
	Name      string `json:"name"`      // f_db_type f_tb_name f_db_name
	Operation string `json:"operation"` // eq
	OpValue   string `json:"op_value"`
}

type GraphFullTextResp struct {
	Res struct {
		Count  int `json:"count"`
		Result []struct {
			Alias   string `json:"alias"`
			Color   string `json:"color"`
			Icon    string `json:"icon"`
			Tag     string `json:"tag"`
			Vertexs []struct {
				Id              string `json:"id"`
				Color           string `json:"color"`
				Icon            string `json:"icon"`
				DefaultProperty struct {
					A string `json:"a"`
					N string `json:"n"`
					V string `json:"v"`
				} `json:"default_property"`
				Tags       []string `json:"tags"`
				Properties []struct {
					Props []struct {
						Alias string `json:"alias"`
						Name  string `json:"name"`
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"props"`
					Tag string `json:"tag"`
				} `json:"properties"`
			} `json:"vertexs"`
		} `json:"result"`
	} `json:"res"`
}
