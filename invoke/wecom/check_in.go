package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcGetCheckInData fetches assembled check-in data within [startTime, endTime].
//
// It maps to the server-side BuildCheckInData(userID, departmentID, startTime, endTime).
// userID and departmentID are optional filters; pass "" to ignore them. The result is
// a nested list grouped by department/user/month depending on which filters are set.
func InvokeRpcGetCheckInData(userID, departmentID string, startTime, endTime int64, ctx context.Context) ([]map[string]interface{}, error) {
	map1 := map[string]interface{}{
		"userID":       userID,
		"departmentID": departmentID,
		"startTime":    startTime,
		"endTime":      endTime,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCheckInData", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
