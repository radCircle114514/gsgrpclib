package wecom

import (
	"context"
	"time"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeWecomRpcQueryDataInTimeScope(table, column string, startTime, endTime *time.Time, filter map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
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
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryDataInTimeScope", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
