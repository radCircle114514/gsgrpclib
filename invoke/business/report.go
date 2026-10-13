package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// ReportOverviewFilterInAgent 是 Agent 侧报表总览入口的过滤条件。该入口
// 只返回当前“待汇报”的活跃记录（不支持历史 / 详情），因此只有分页参数；
// 服务端不按调用方做行级过滤，返回全公司当日待汇报队列。
type ReportOverviewFilterInAgent struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

// InvokeRpcFindReportOverviewInAgent 供 Agent 调用，分页返回当前所有
// “待汇报”的活跃 report_info 记录（delete_flag=0 且 is_finished=0）。
// 返回通用 map 行与总条数；不涉及历史表。
func InvokeRpcFindReportOverviewInAgent(filter ReportOverviewFilterInAgent, ctx context.Context) ([]map[string]interface{}, int64, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindReportOverviewInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	list := utils.ParseAnyToMapList(vo.MapList)
	count, _ := utils.ParseAnyToData(vo.Data).(int64)
	return list, count, nil
}
