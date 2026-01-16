package knowledge_network

import (
	"context"
	"fmt"
	"net/http"
)

type SearchConfig struct {
	Tag        string        `json:"tag"`
	Properties []*SearchProp `json:"properties"`
}

type SearchProp struct {
	Name      string `json:"name"`      // f_db_type f_tb_name f_db_name
	Operation string `json:"operation"` // eq
	OpValue   string `json:"op_value"`
}

type ADLineageFulltextResp struct {
	Res *FulltextSearchResp `json:"res"`
}

type ADLineageFulltextV2Resp struct {
	Res struct {
		Nodes []struct {
			Id              string `json:"id"`
			Alias           string `json:"alias"`
			Color           string `json:"color"`
			ClassName       string `json:"class_name"`
			Icon            string `json:"icon"`
			DefaultProperty struct {
				Name  string `json:"name"`
				Value string `json:"value"`
				Alias string `json:"alias"`
			} `json:"default_property"`
			Tags       []string `json:"tags"`
			Properties []struct {
				Tag   string `json:"tag"`
				Props []struct {
					Name     string `json:"name"`
					Value    string `json:"value"`
					Alias    string `json:"alias"`
					Type     string `json:"type"`
					Disabled bool   `json:"disabled"`
					Checked  bool   `json:"checked"`
				} `json:"props"`
			} `json:"properties"`
		} `json:"nodes"`
		Edges []interface{} `json:"edges"`
	} `json:"res"`
}

type FulltextSearchResp struct {
	Count  int            `json:"count"`  // tag总数
	Result []*TagInfoResp `json:"result"` // 结果列表
}

type TagInfoResp struct {
	Alias    string              `json:"alias"`   // 别名
	Color    string              `json:"color"`   // 颜色
	Icon     string              `json:"icon"`    // 图标
	Tag      string              `json:"tag"`     // 点的类型
	Vertexes []*FulltextVertexes `json:"vertexs"` // 点的列表
}

type FulltextVertexes struct {
	ID              string            `json:"id"`
	Color           string            `json:"color"`
	Icon            string            `json:"icon"`
	DefaultProperty *Property         `json:"default_property"`
	Tags            []string          `json:"tags"`
	Properties      []*PropertiesInfo `json:"properties"` //实体属性
}

type Property struct {
	A string `json:"a"` // 属性别名 alias缩写
	N string `json:"n"` // 属性字段名 name缩写
	V string `json:"v"` // 属性值 value缩写
}

type PropertiesInfo struct {
	Props []*Prop `json:"props"` // 属性集合
	Tag   string  `json:"tag"`   // 实体类名
}

type Prop struct {
	Alias string `json:"alias"` // 属性显示名
	Name  string `json:"name"`  // 属性名
	Type  string `json:"type"`  // 属性类型
	Value string `json:"value"` // 属性值
}

func NewADLineageFulltextReqBody(kgID, query string, config []*SearchConfig) any {
	requestMap := map[string]any{
		"kg_id":         kgID,
		"query":         query,
		"page":          1,
		"size":          0,
		"matching_rule": "portion",
		"matching_num":  20,
	}
	if len(config) > 0 {
		requestMap["search_config"] = config
	}

	return requestMap
}

func (a *ad) FulltextSearch(ctx context.Context, kgID string, query string, config []*SearchConfig) (*ADLineageFulltextResp, error) {
	fulltextUrl := fmt.Sprintf("%s/api/alg-server/v1/open/graph-search/kgs/%s/full-text", a.baseUrl, kgID)
	body := NewADLineageFulltextReqBody(kgID, query, config)
	return httpPostDo[ADLineageFulltextResp](ctx, fulltextUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func (a *ad) FulltextSearchV2(ctx context.Context, kgID string, query string, config []*SearchConfig) (*ADLineageFulltextV2Resp, error) {
	fulltextUrl := fmt.Sprintf("%s/api/engine/v1/open/basic-search/kgs/%s/full-text", a.baseUrl, kgID)
	body := NewADLineageFulltextReqBody(kgID, query, config)
	return httpPostDo[ADLineageFulltextV2Resp](ctx, fulltextUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func (a *ad) GetSearchConfig(tagName string, propertyName string, propertyValue string) ([]*SearchConfig, error) {
	searchCfg := make([]*SearchConfig, 0)

	searchProperties := make([]*SearchProp, 0)
	searchProperties = append(searchProperties, &SearchProp{propertyName, "eq", propertyValue})
	searchCfg = append(searchCfg, &SearchConfig{tagName, searchProperties})

	return searchCfg, nil
}

func NewADLineageNeighborsReqBody(kgID, vid string, steps int) any {
	requestMap := map[string]any{
		"id":        kgID,
		"steps":     steps,
		"direction": "reverse",
		"vids": []string{
			vid,
		},
		"final_step": true,
		"page":       1,
		"size":       -1,
		"filters":    []string{},
	}

	return requestMap
}

func NewADLineageNeighborsReqBodyV2(kgID, vid string, steps int) any {
	requestMap := map[string]any{
		"id":        kgID,
		"steps":     steps,
		"direction": "positive",
		"vids": []string{
			vid,
		},
		"final_step": true,
		"page":       1,
		"size":       -1,
		"filters":    []string{},
	}

	return requestMap
}

type SearchConfigItem struct {
	Tag        string             `json:"tag"`
	Properties []PropertiesStruct `json:"properties"`
}

type PropertiesStruct struct {
	Name      string `json:"name"`
	Operation string `json:"operation"`
	OpValue   string `json:"op_value"`
}

type ADLineageNeighborsResp struct {
	Res *LineageNeighborsResp `json:"res"`
}

type LineageNeighborsResp struct {
	VCount  int                `json:"v_count"`  // 点的数量
	VResult []*NeighborsVGroup `json:"v_result"` // 点的结果
}

type NeighborsVGroup struct {
	Alias    string              `json:"alias"`    // 别名
	Color    string              `json:"color"`    // 颜色
	Icon     string              `json:"icon"`     // 图标
	Tag      string              `json:"tag"`      // 点的类型
	Vertexes []*NeighborsVResult `json:"vertexes"` // 点的列表
}

type NeighborsVResult struct {
	DefaultProperty *Property         `json:"default_property"` // 属性
	ID              string            `json:"id"`               // 点的id
	InEdges         []string          `json:"in_edges"`         // 进边列表
	OutEdges        []string          `json:"out_edges"`        // 出边列表
	Properties      []*PropertiesInfo `json:"properties"`       // 属性列表
	Tags            []string          `json:"tags"`             // 实体tag
}

func (a *ad) NeighborSearch(ctx context.Context, vid string, steps int, kgId string) (*ADLineageNeighborsResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/alg-server/v1/open/explore/kgs/%s/neighbors", a.baseUrl, kgId)
	body := NewADLineageNeighborsReqBody(kgId, vid, steps)
	return httpPostDo[ADLineageNeighborsResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func NewADEntitySearchReqBody(kgID, entityId string) any {
	requestMap := map[string]any{
		"page":  1,
		"size":  0,
		"query": "",
		"search_config": []SearchConfigItem{
			{"logical_entity", []PropertiesStruct{
				{"id", "eq", entityId},
			}},
		},
		"kg_id":         kgID,
		"matching_rule": "completeness",
		"matching_num":  20}

	return requestMap
}

type ADFullResp struct {
	Res *FullSearchResp `json:"res"`
}

type FullSearchResp struct {
	Nodes []NodeItem `json:"nodes"`
	Edges []NodeItem `json:"edges"`
}

type NodeItem struct {
	Id string `json:"id"`

	Alias           string `json:"alias"`
	Color           string `json:"color"`
	ClassName       string `json:"class_name"`
	Icon            string `json:"icon"`
	DefaultProperty struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Alias string `json:"alias"`
	} `json:"default_property"`
	Tags       []string `json:"tags"`
	Properties []struct {
		Tag   string `json:"tag"`
		Props []struct {
			Name     string `json:"name"`
			Value    string `json:"value"`
			Alias    string `json:"alias"`
			Type     string `json:"type"`
			Disabled bool   `json:"disabled"`
			Checked  bool   `json:"checked"`
		} `json:"props"`
	} `json:"properties"`
}

// 使用anydata engine 接口
func (a *ad) EntitySearchByEngine(ctx context.Context, entityId string, kgId string) (*ADFullResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/engine/v1/open/basic-search/kgs/%s/full-text", a.baseUrl, kgId)
	body := NewADEntitySearchReqBody(kgId, entityId)
	return httpPostDo[ADFullResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

type ADLineageNeighborsRespV2 struct {
	Res *LineageNeighborsRespV2 `json:"res"`
}

type LineageNeighborsRespV2 struct {
	NodesCount int                  `json:"nodes_count"` // 点的数量
	Nodes      []*NeighborsVGroupV2 `json:"nodes"`       // 点的结果
}

type NeighborsVGroupV2 struct {
	Id              string `json:"id"`
	Alias           string `json:"alias"`
	Color           string `json:"color"`
	ClassName       string `json:"class_name"`
	Icon            string `json:"icon"`
	DefaultProperty struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Alias string `json:"alias"`
	} `json:"default_property"`
	Tags       []string `json:"tags"`
	Properties []struct {
		Tag   string `json:"tag"`
		Props []struct {
			Name     string `json:"name"`
			Value    string `json:"value"`
			Alias    string `json:"alias"`
			Type     string `json:"type"`
			Disabled bool   `json:"disabled"`
			Checked  bool   `json:"checked"`
		} `json:"props"`
	} `json:"properties"`
}

// 使用anydata engine 接口
func (a *ad) NeighborSearchByEngine(ctx context.Context, vid string, steps int, kgId string) (*ADLineageNeighborsRespV2, error) {
	neighborUrl := fmt.Sprintf("%s/api/engine/v1/open/graph-explore/kgs/%s/neighbors", a.baseUrl, kgId)
	body := NewADLineageNeighborsReqBody(kgId, vid, steps)
	return httpPostDo[ADLineageNeighborsRespV2](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

// 使用anydata engine 接口
func (a *ad) NeighborSearchByEngineV2(ctx context.Context, vid string, steps int, kgId string) (*ADLineageNeighborsRespV2, error) {
	neighborUrl := fmt.Sprintf("%s/api/engine/v1/open/graph-explore/kgs/%s/neighbors", a.baseUrl, kgId)
	body := NewADLineageNeighborsReqBodyV2(kgId, vid, steps)
	return httpPostDo[ADLineageNeighborsRespV2](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}
