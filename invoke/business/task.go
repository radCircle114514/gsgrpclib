package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// TaskOverviewFilterInAgent 是 Agent 侧任务总览入口的过滤条件。
// History=true 查已办任务（按流程实例聚合），否则查待办任务。
// 行级可见范围由服务端按已认证 userID 的权限判定：普通用户强制只看
// 本人（待办 assignee=本人；已办本人处理过），持有“查看所有任务”
// 权限时才认 AssigneeUserID 作为处理人过滤。
type TaskOverviewFilterInAgent struct {
	CurrentPage    int    `json:"currentPage"`
	PageSize       int    `json:"pageSize"`
	SerialNumber   string `json:"serialNumber"`
	G3Number       string `json:"g3Number"`
	Title          string `json:"title"`
	AssigneeUserID string `json:"assigneeUserId"`
	History        bool   `json:"history"`
}

// TaskDetailInAgent 是 Agent 侧单条任务详情的定位参数。TaskID 优先；
// 也可用 SerialNumber / G3Number 定位当前/最近任务。服务端先查待办表
// task，未命中再查历史表 task_history。
type TaskDetailInAgent struct {
	TaskID       int64  `json:"taskId"`
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
}

// InvokeRpcFindTaskOverviewInAgent 供 Agent 调用的任务总览：依据
// filter.History 路由到待办或已办列表，返回通用 map 行与总条数。
func InvokeRpcFindTaskOverviewInAgent(filter TaskOverviewFilterInAgent, ctx context.Context) ([]map[string]interface{}, int64, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindTaskOverviewInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	list := utils.ParseAnyToMapList(vo.MapList)
	count, _ := utils.ParseAnyToData(vo.Data).(int64)
	return list, count, nil
}

// InvokeRpcFindTaskDetailInAgent 供 Agent 调用，按 task_id 或序列号 /
// G3 单号查询单条任务完整详情（任务主体、表单、状态、orderInfo）。
// 结果是响应 Map 字段里的复合对象。
func InvokeRpcFindTaskDetailInAgent(filter TaskDetailInAgent, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindTaskDetailInAgent", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
