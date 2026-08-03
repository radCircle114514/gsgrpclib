package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcGetRoleList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetRoleList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetRoleByUserID(userID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetRoleByUserID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetUserByRoleName(roleName string, ctx context.Context) ([]interface{}, error) {
	data, err := utils.ParseDataToAny(roleName)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserByRoleName", ao)
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToDataList(vo.DataList)
	return userList, nil
}
