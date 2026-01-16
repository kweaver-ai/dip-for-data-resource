package client

type GraphNeighborsReq struct {
	KgType string `json:"kg_type"` // 图谱类型
	Id     string `json:"id"`
	Steps  int    `json:"steps"`
	Vid    string `json:"vid"`
}

type GraphNeighborsResp struct {
	Res struct {
		VCount  int `json:"v_count"`
		VResult []struct {
			Alias    string `json:"alias"`
			Color    string `json:"color"`
			Icon     string `json:"icon"`
			Tag      string `json:"tag"`
			Vertexes []struct {
				DefaultProperty struct {
					A string `json:"a"`
					N string `json:"n"`
					V string `json:"v"`
				} `json:"default_property"`
				Id         string   `json:"id"`
				InEdges    []string `json:"in_edges"`
				OutEdges   []string `json:"out_edges"`
				Properties []struct {
					Tag   string `json:"tag"`
					Props []struct {
						Alias string `json:"alias"`
						Name  string `json:"name"`
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"props"`
				} `json:"properties"`
				Tags []string `json:"tags"`
			} `json:"vertexes"`
		} `json:"v_result"`
	} `json:"res"`
}
