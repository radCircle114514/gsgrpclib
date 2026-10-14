package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcFindOrderBillOverviewInAgent 供 Agent 调用：分页返回当前认证
// 用户可见的每张订单账单合计（应收 / 应付 / 毛利）。过滤条件与订单列表
// 一致，复用 OrderFilterInAgent。返回通用 map 行与总条数。
func InvokeRpcFindOrderBillOverviewInAgent(filter OrderFilterInAgent, ctx context.Context) ([]map[string]interface{}, int64, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindOrderBillOverviewInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	list := utils.ParseAnyToMapList(vo.MapList)
	count, _ := utils.ParseAnyToData(vo.Data).(int64)
	return list, count, nil
}

// InvokeRpcDetailOrderBillInAgent 供 Agent 调用：按序列号或 G3 单号查询
// 单张订单的账单明细（list）与毛利 / 毛利率（op_profit / op_profit_rate）。
// 复用 OrderDetailInAgent。结果是响应 Map 字段里的复合对象。
func InvokeRpcDetailOrderBillInAgent(filter OrderDetailInAgent, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "DetailOrderBillInAgent", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
