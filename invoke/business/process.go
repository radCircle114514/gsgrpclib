package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// ProcessFilterInAgent 是 Agent 侧查询流程实例列表的过滤条件，作为
// FindOngoingProcessInfoInAgent / FindFinishedProcessInfoInAgent 的
// Data（JSON）载荷。用户身份由服务端从 gRPC metadata 中的认证信息
// 解析，不在此传递，防止越权冒充。
type ProcessFilterInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
	Title        string `json:"title"`
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
}

// invokeProcessInAgent 是进行中/已完成流程实例两个 Agent 入口的公共调用：
// 过滤条件放在 ao.Data，流程实例列表在 vo.ProcInstVOList，总条数在 vo.Data。
// 行级可见范围由服务端依据已认证 userID 的权限判定，调用方无法自行扩大。
func invokeProcessInAgent(methodName string, filter ProcessFilterInAgent, ctx context.Context) ([]*business_rpc.ProcInstVO, interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, methodName, ao)
	if err != nil {
		return nil, 0, err
	}
	return vo.ProcInstVOList, utils.ParseAnyToData(vo.Data), nil
}

// InvokeRpcFindOngoingProcessInfoInAgent 供 Agent 调用，分页查询当前
// 认证用户可见的进行中流程实例列表，返回实例与总条数。
func InvokeRpcFindOngoingProcessInfoInAgent(filter ProcessFilterInAgent, ctx context.Context) ([]*business_rpc.ProcInstVO, interface{}, error) {
	return invokeProcessInAgent("FindOngoingProcessInfoInAgent", filter, ctx)
}

// InvokeRpcFindFinishedProcessInfoInAgent 供 Agent 调用，分页查询当前
// 认证用户可见的已完成流程实例列表，返回实例与总条数。
func InvokeRpcFindFinishedProcessInfoInAgent(filter ProcessFilterInAgent, ctx context.Context) ([]*business_rpc.ProcInstVO, interface{}, error) {
	return invokeProcessInAgent("FindFinishedProcessInfoInAgent", filter, ctx)
}

// ProcessOverviewFilterInAgent 是 Agent 侧统一流程总览入口的过滤条件。
// History=true 查已完成流程（历史表），否则查进行中流程。行级可见范围
// 由服务端依据已认证 userID 的权限判定：持有“查看所有流程”权限返回全部，
// 否则只返回本人发起 / 操作 / 销售相关的流程。
type ProcessOverviewFilterInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
	Title        string `json:"title"`
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
	History      bool   `json:"history"`
}

// ProcessDetailInAgent 是 Agent 侧单条流程详情的定位参数。SerialNumber
// 优先于 G3Number；服务端先查进行中表，再查历史表，因此已结束流程也可取到。
type ProcessDetailInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
}

// InvokeRpcFindProcessOverviewInAgent 供 Agent 调用的统一流程总览：
// 依据 filter.History 路由到进行中或已完成列表，返回实例列表与总条数。
func InvokeRpcFindProcessOverviewInAgent(filter ProcessOverviewFilterInAgent, ctx context.Context) ([]*business_rpc.ProcInstVO, interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindProcessOverviewInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	return vo.ProcInstVOList, utils.ParseAnyToData(vo.Data), nil
}

// InvokeRpcFindProcessDetailInAgent 供 Agent 调用，按序列号或 G3 单号
// 查询单条流程完整详情（主体、审批时间线、当前任务、表单变量），并内嵌
// 同业务单号的订单详情 orderInfo。结果是响应 Map 字段里的复合对象：
// {"procInst": {...}, "orderInfo": {...}}（无订单可见权限时缺省 orderInfo）。
func InvokeRpcFindProcessDetailInAgent(filter ProcessDetailInAgent, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindProcessDetailInAgent", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
