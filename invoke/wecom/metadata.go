package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcWecomQueryFieldList(table, field, value, userID string, ctx context.Context) ([]interface{}, error) {
	map1 := map[string]interface{}{
		"table":  table,
		"field":  field,
		"value":  value,
		"userID": userID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryFieldList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return dataList, nil
}

func InvokeRpcGetFindEnumList(type1 string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(type1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindEnumList", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}

func InvokeRpcQueryNameByType(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	data, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryNameByType", ao)
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}
