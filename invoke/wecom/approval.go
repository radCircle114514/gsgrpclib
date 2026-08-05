package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcGetApprovalDetail(spNo string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetApprovalDetailAO{
		SpNo: spNo,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.GetApprovalDetail(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetApprovalList(startTime, endTime int64, ctx context.Context) ([]map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetApprovalListAO{
		StartTime: startTime,
		EndTime:   endTime,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.GetApprovalList(ctx, ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcCreateApproval(dataJson string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.CreateApprovalAO{
		DataJson: dataJson,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.CreateApproval(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetApprovalTemplate(templateID string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetApprovalTemplateAO{
		TemplateID: templateID,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.GetApprovalTemplate(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}
