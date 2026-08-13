package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcGetReportList retrieves the HR report/journal record list within
// [startTime, endTime] for the given templateID.
func InvokeRpcGetReportList(startTime, endTime int64, templateID string, ctx context.Context) (map[string]interface{}, error) {
	map1 := map[string]interface{}{
		"startTime":  startTime,
		"endTime":    endTime,
		"templateID": templateID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetReportList", ao)
	if err != nil {
		return nil, err
	}
	mapResult := utils.ParseAnyToMap(vo.Map)
	return mapResult, nil
}

// InvokeRpcGetReportDetail retrieves a single HR report/journal record by its journalID.
func InvokeRpcGetReportDetail(journalID string, ctx context.Context) (map[string]interface{}, error) {
	map1 := map[string]interface{}{
		"journalID": journalID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetReportDetail", ao)
	if err != nil {
		return nil, err
	}
	mapResult := utils.ParseAnyToMap(vo.Map)
	return mapResult, nil
}
