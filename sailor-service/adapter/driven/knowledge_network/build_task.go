package knowledge_network

import (
	"context"
	"fmt"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

type GraphBuildTaskReq struct {
	TaskType string `json:"tasktype"`
}

type GraphBuildTaskResp struct {
	GraphTaskID int `json:"graph_task_id"`
}

type GraphGroupTaskListReq struct {
	Page        int    `json:"page" form:"page,default=1"`
	Size        int    `json:"size" form:"size,default=20"`
	Order       int    `json:"order" form:"order,default=desc"`
	Status      string `json:"status" form:"status,default=all"`
	GraphName   string `json:"graph_name" form:"graph_name"`
	TaskType    string `json:"task_type" form:"task_type,default=all"`
	TriggerType string `json:"trigger_type" form:"trigger_type,default=all"`
	Rule        string `json:"rule" form:"rule,default=start_time"`
}

type TaskDetailListRes struct {
	Count       int           `json:"count"`
	DF          []*TaskDetail `json:"df"`
	GraphStatus string        `json:"graph_status"`
}

type TaskDetail struct {
	AllTime          string      `json:"all_time"`
	CeleryTaskId     string      `json:"celery_task_id"`
	CountStatus      int         `json:"count_status"`
	CreateUser       string      `json:"create_user"`
	CreateUserName   string      `json:"create_user_name"`
	Edge             string      `json:"edge"`
	EdgeNum          int         `json:"edge_num"`
	EdgeProNum       int         `json:"edge_pro_num"`
	EffectiveStorage bool        `json:"effective_storage"`
	EndTime          string      `json:"end_time"`
	Entity           string      `json:"entity"`
	EntityNum        int         `json:"entity_num"`
	EntityProNum     int         `json:"entity_pro_num"`
	ErrorReport      string      `json:"error_report"`
	ExtractType      interface{} `json:"extract_type"`
	Files            interface{} `json:"files"`
	GraphEdge        int         `json:"graph_edge"`
	GraphEntity      int         `json:"graph_entity"`
	GraphId          int         `json:"graph_id"`
	GraphName        string      `json:"graph_name"`
	KgId             int         `json:"kg_id"`
	Parent           interface{} `json:"parent"`
	StartTime        string      `json:"start_time"`
	SubgraphId       int         `json:"subgraph_id"`
	TaskId           int         `json:"task_id"`
	TaskName         string      `json:"task_name"`
	TaskStatus       string      `json:"task_status"`
	TaskType         string      `json:"task_type"`
	TriggerType      int         `json:"trigger_type"`
}

func (a *ad) CreateGraphTask(ctx context.Context, graphID int, taskType string) (id int, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + fmt.Sprintf(`/api/builder/v1/task/%v`, graphID)

	args := &GraphBuildTaskReq{
		TaskType: taskType,
	}
	resp, err := httpPostDo[commonResp[GraphBuildTaskResp]](ctx, rawURL, args, nil, a)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return resp.Res.GraphTaskID, nil
}

func (a *ad) GraphTaskList(ctx context.Context, graphID int, req *GraphGroupTaskListReq) (data *TaskDetailListRes, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	rawURL := a.baseUrl + fmt.Sprintf(`/api/builder/v1/task/%v`, graphID)

	resp, err := httpPostDo[commonResp[TaskDetailListRes]](ctx, rawURL, req, nil, a)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &resp.Res, nil
}

func (a *ad) DeleteTask(ctx context.Context, graphID int, taskIDSlice []int) (err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	body := make(map[string][]int)
	body["task_ids"] = taskIDSlice
	rawURL := a.baseUrl + fmt.Sprintf(`/api/builder/v1/task/delete/%v`, graphID)

	_, err = httpPostDo[commonResp[string]](ctx, rawURL, body, nil, a)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
