package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcFindCustomerList(remark, userID, searcher string, currentPage, pageSize int, all bool, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindCustomerAO{
		Remark:      remark,
		UserID:      userID,
		Searcher:    searcher,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		All:         all,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.FindCustomerList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	dataList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return dataList, count, nil
}

func InvokeRpcFindCustomerCouldBeSelectedList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindCustomerCouldBeSelectedList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcAddCustomer(name, userID, markName, phone string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.AddCustomerAO{
		Name:     name,
		UserID:   userID,
		MarkName: markName,
		Phone:    phone,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	_, err = rpcClient.AddCustomer(ctx, ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcUpdateCustomerRemark(userID, markName, oldMarkName string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.UpdateCustomerAO{
		UserID:      userID,
		MarkName:    markName,
		OldMarkName: oldMarkName,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	_, err = rpcClient.UpdateCustomerRemark(ctx, ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetCustomerInRobotChat(externalUserID, remark string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerInRobotChatAO{
		ExternalUserID: externalUserID,
		Remark:         remark,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.GetCustomerInRobotChat(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcFlushCustomer(userID string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	res, err := client.InvokeWecomRPCMethod(ctx, "FlushCustomer", ao)
	if err != nil {
		return "", err
	}
	return res.Msg, nil
}

func InvokeRPCGetNeedNotifyCustomerList(spNoList []string, ctx context.Context) ([]map[string]interface{}, error) {
	jsonStr, _ := util.ToJSONStr(spNoList)
	anyList, _ := utils.ParseJsonStrToAnyList(jsonStr)
	ao := &wecom_rpc.RequestAO{
		DataList: anyList,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetNeedNotifyCustomerList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRPCFindClaimedCustomer(ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindClaimedCustomer", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}
