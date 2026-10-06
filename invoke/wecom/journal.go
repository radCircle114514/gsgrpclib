package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcGetReportList fetches all journal record UUIDs within
// [startTime, endTime]. filters narrows the result; pass nil to fetch
// everything. Each map entry maps to one item in the WeCom "filters" array,
// e.g. {"key":"template_id","value":"..."}, {"key":"creator","value":"..."},
// {"key":"department","value":"1"}.
func InvokeRpcGetReportList(filters []map[string]interface{}, startTime, endTime int64, ctx context.Context) ([]string, error) {
	map1 := map[string]interface{}{
		"startTime": startTime,
		"endTime":   endTime,
		"filters":   filters,
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
	result := utils.ParseAnyToMap(vo.Map)
	journalUuidList := make([]string, 0)
	if list, ok := result["journalUuidList"].([]interface{}); ok {
		for _, item := range list {
			if uuid, ok := item.(string); ok {
				journalUuidList = append(journalUuidList, uuid)
			}
		}
	}
	return journalUuidList, nil
}

// InvokeRpcGetReportDetail fetches one journal record by its UUID, including
// submitter/receivers, apply_data contents and comments, with user names enriched.
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
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcGetReportStatList fetches the aggregated journal statistics for a
// template within [startTime, endTime]. Only fully aggregated cycles are
// returned by the WeCom server.
func InvokeRpcGetReportStatList(templateID string, startTime, endTime int64, ctx context.Context) ([]map[string]interface{}, error) {
	map1 := map[string]interface{}{
		"templateID": templateID,
		"startTime":  startTime,
		"endTime":    endTime,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetReportStatList", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

// InvokeRpcGetJournalData fetches complete journal entries within
// [startTime, endTime]. filters narrows the result; pass nil to fetch
// everything. List and detail calls are merged server-side; each returned
// map is a ready-to-use journal record with submitter, apply_data contents,
// comments and enriched user names.
func InvokeRpcGetJournalData(filters []map[string]interface{}, startTime, endTime int64, ctx context.Context) ([]map[string]interface{}, error) {
	map1 := map[string]interface{}{
		"startTime": startTime,
		"endTime":   endTime,
		"filters":   filters,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetJournalData", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}
