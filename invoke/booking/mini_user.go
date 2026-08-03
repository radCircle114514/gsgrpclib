package booking

import (
	"context"

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
