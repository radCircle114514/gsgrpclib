package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcFindCustomerGroupList(name, userID, searcher string, currentPage, pageSize int, all bool, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindCustomerGroupAO{
		Name:        name,
		UserID:      userID,
		Searcher:    searcher,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		All:         all,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.FindCustomerGroupList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	dataList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return dataList, count, nil
}

func InvokeRpcSyncCustomerGroup(userID string, ctx context.Context) error {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncCustomerGroup", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcFlushCustomerGroup(chatID string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(chatID)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	res, err := client.InvokeWecomRPCMethod(ctx, "FlushCustomerGroup", ao)
	if err != nil {
		return "", err
	}
	return res.Msg, nil
}

func InvokeRpcQueryCustomerGroupByCond(table, field, value string, flag bool, ctx context.Context) ([]map[string]interface{}, []interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.OptionAO{
		Table: table,
		Field: field,
		Value: value,
		Flag:  flag,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.QueryCustomerGroupByCond(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return mapList, dataList, nil
}

func InvokeRpcGetCustomerGroup(chatID, groupName string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerGroupAO{
		ChatID:    chatID,
		GroupName: groupName,
	}
	vo, err := rpcClient.GetCustomerGroup(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetCustomerGroupList(chatID string, ctx context.Context) ([]map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerGroupAO{
		ChatID: chatID,
	}
	vo, err := rpcClient.GetCustomerGroupList(ctx, ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetCustomerGroupChat(name string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(name)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerGroupChat", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}

func InvokeRpcGetCustomerGroupChatByChatID(chatID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(chatID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerGroupChatByChatID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}
