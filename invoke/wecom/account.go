package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcFindAccountUserList(remark string, currentPage, pageSize, tagID int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindAccountAO{
		Remark:      remark,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		TagID:       int32(tagID),
	}
	vo, err := rpcClient.FindAccountUserList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcSyncAccountUser(ctx context.Context) (string, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "SyncAccountUser", &wecom_rpc.RequestAO{})
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcFlushAccountUser(openID string, ctx context.Context) error {
	data, err := utils.ParseDataToAny(openID)
	if err != nil {
		return nil
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	_, err = client.InvokeWecomRPCMethod(ctx, "FlushAccountUser", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetAccountUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAccountUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetAccountUserListByOpenIDList(openIDList []string, ctx context.Context) ([]map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.OpenIDListAO{
		OpenIDList: openIDList,
	}
	vo, err := rpcClient.GetAccountUserListByOpenIDList(ctx, ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcModifyAccountUserRemark(openID, remark string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.ModifyRemarkAO{
		OpenID: openID,
		Remark: remark,
	}
	_, err = rpcClient.ModifyAccountUserRemark(ctx, ao)
	if err != nil {
		return nil
	}
	return nil
}

func InvokeRpcGetAccountTagList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAccountTagList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetOauthCallBackURL(id string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(id)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetOauthCallBackURL", ao)
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}
