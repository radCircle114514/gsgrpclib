package invoke

import (
	"context"
	"time"

	"github.com/iznilul/gsgrpclib/client"
	booking_rpc "github.com/iznilul/gsgrpclib/proto/booking"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcBookingFindMiniUser(userID string, ctx context.Context) (map[string]interface{}, error) {
	toAny, _ := utils.ParseDataToAny(userID)
	ao := &booking_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "FindMiniUser", ao)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	mapList := utils.ParseAnyToMapList(vo.MapList)
	for _, m := range mapList {
		openID := m["openID"].(string)
		result[openID] = m["remark"]
	}
	return result, nil
}

func InvokeRpcBookingFindMiniUserList(ctx context.Context) ([]map[string]interface{}, error) {
	ao := &booking_rpc.RequestAO{}
	vo, err := client.InvokeBookingRPCMethod(ctx, "FindMiniUserList", ao)
	if err != nil {
		return nil, err
	}
	miniUserList := utils.ParseAnyToMapList(vo.MapList)
	return miniUserList, nil
}

func InvokeRpcBookingFindMiniUserListByOpenIDList(openIDList []string, ctx context.Context) ([]map[string]interface{}, error) {
	jsonStr, _ := util.ToJSONStr(openIDList)
	datalist, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		DataList: datalist,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "FindMiniUserListByOpenIDList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcBookingSendMiniMsg(toMap map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(toMap)
	if err != nil {
		return err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBookingRPCMethod(ctx, "SendMiniMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeBookingRPCQueryIndicatorCount(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorCount", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapInt(vo.Map)
	return result, nil
}

func InvokeBookingRPCQueryIndicatorDetail(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorDetail", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeBookingRPCCalculateUserIndicator(queryAO map[string]map[string]interface{}, ctx context.Context) (map[string]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "CalculateUserIndicator", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeRPCGetGiftList(index string, ctx context.Context) ([]map[string]interface{}, error) {
	toAny, _ := utils.ParseDataToAny(index)
	ao := &booking_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "GetGiftList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeBookingRPCQueryIndicatorCountInBatch(queryAO map[int][]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorCountInBatch", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeBookingRPCQueryDataInTimeScope(table, column string, startTime, endTime *time.Time, filter map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
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
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryDataInTimeScope", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
