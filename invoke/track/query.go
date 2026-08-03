package track

import (
	"context"
	"time"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcTrackQueryFieldList(table, field, value string, ctx context.Context) ([]interface{}, error) {
	map1 := map[string]interface{}{
		"table": table,
		"field": field,
		"value": value,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "QueryFieldList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return dataList, nil
}

func InvokeRpcTrackQueryDataInTimeScope(table, column string, startTime, endTime *time.Time, filter map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	queryAO := map[string]interface{}{
		"table":  table,
		"column": column,
		"filter": filter,
	}
	if startTime != nil {
		queryAO["startTime"] = *startTime
	}
	if endTime != nil {
		queryAO["endTime"] = *endTime
	}
	toAny, err := utils.ParseMapToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "QueryDataInTimeScope", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
