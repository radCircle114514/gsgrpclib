package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcGetUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}

func InvokeRpcGetAllUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAllUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}

func InvokeRpcGetUserInfo(code string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(code)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserInfo", ao)
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcGetUserDetailList(userID string, dataList []interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	data, _ := utils.ParseDataToAny(userID)
	result, _ := utils.ParseDataListToAnyList(dataList)
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserDetailList", &wecom_rpc.RequestAO{
		Data:     data,
		DataList: result,
	})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetUserDetailByUserID(userID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserDetailByUserID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRPCListUserByCond(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toMap, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toMap,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ListUserByCond", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

func InvokeWecomRpcUpdateUser(userInfo map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(userInfo)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateUser", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetUserIdentity(code string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(code)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserIdentity", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetDepartmentTree(ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetDepartmentTree", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}
