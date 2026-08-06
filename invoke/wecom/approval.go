package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcGetApprovalDetail(spNo string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(spNo)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetApprovalDetail", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

// InvokeRpcGetApprovalList fetches approval numbers in [startTime, endTime].
// startTime/endTime are unix timestamps (seconds). filters narrows the result;
// pass nil to fetch everything. Each map entry maps to one item in the WeCom
// "filters" array, e.g. {"key":"template_id","value":"..."} or
// {"key":"sp_no","obj_value":{"new_id":"..."}}.
func InvokeRpcGetApprovalList(startTime, endTime int64, filters []map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
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
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetApprovalList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
