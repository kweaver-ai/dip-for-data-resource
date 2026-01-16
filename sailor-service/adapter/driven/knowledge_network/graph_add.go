package knowledge_network

import (
	"context"
	"fmt"
	"net/http"
)

type InsertEntityReqBody struct {
	Action    string           `json:"action"`
	DataType  string           `json:"data_type"`
	GraphData []map[string]any `json:"graph_data"`
	GraphId   int              `json:"graph_id"`
	Name      string           `json:"name"`
}

type InsertSideReqBody struct {
	Action    string                         `json:"action"`
	DataType  string                         `json:"data_type"`
	GraphData []map[string]map[string]string `json:"graph_data"`
	GraphId   int                            `json:"graph_id"`
	Name      string                         `json:"name"`
}

type DeleteEntityReqBody struct {
	Action    string              `json:"action"`
	DataType  string              `json:"data_type"`
	GraphData []map[string]string `json:"graph_data"`
	GraphId   int                 `json:"graph_id"`
	Name      string              `json:"name"`
}

type InsertGraphResp struct {
	Res string `json:"res"`
}

func (a *ad) InsertEntity(ctx context.Context, dataType string, graphData []map[string]any, kgId int, entityName string) (*InsertGraphResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/builder/v1/open/graph/data", a.baseUrl)
	body := InsertEntityReqBody{}
	body.Action = "upsert"
	body.DataType = dataType
	body.GraphData = graphData
	body.GraphId = kgId
	body.Name = entityName
	return httpPostDo[InsertGraphResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func (a *ad) InsertSide(ctx context.Context, dataType string, graphData []map[string]map[string]string, kgId int, sideName string) (*InsertGraphResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/builder/v1/open/graph/data", a.baseUrl)
	body := InsertSideReqBody{}
	body.Action = "upsert"
	body.DataType = dataType
	body.GraphData = graphData
	body.GraphId = kgId
	body.Name = sideName
	return httpPostDo[InsertGraphResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func (a *ad) DeleteEntity(ctx context.Context, graphData []map[string]string, entityName string, kgId int) (*InsertGraphResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/builder/v1/open/graph/data", a.baseUrl)
	body := DeleteEntityReqBody{}
	body.Action = "delete"
	body.DataType = "entity"
	body.GraphData = graphData
	body.GraphId = kgId
	body.Name = entityName
	return httpPostDo[InsertGraphResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}

func (a *ad) DeleteEdge(ctx context.Context, graphData []map[string]map[string]string, sideName string, kgId int) (*InsertGraphResp, error) {
	neighborUrl := fmt.Sprintf("%s/api/builder/v1/open/graph/data", a.baseUrl)
	body := InsertSideReqBody{}
	body.Action = "delete"
	body.DataType = "edge"
	body.GraphData = graphData
	body.GraphId = kgId
	body.Name = sideName
	return httpPostDo[InsertGraphResp](ctx, neighborUrl, body, http.Header{"Content-Type": []string{"application/json"}}, a)
}
