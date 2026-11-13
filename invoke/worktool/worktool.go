package worktool

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	worktool_rpc "github.com/iznilul/gsgrpclib/proto/worktool"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcSendTextMessage(ao *worktool_rpc.RequestAO, ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "SendTextMessage", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

func InvokeRpcAddCustomer(ao *worktool_rpc.RequestAO, ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "AddCustomer", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

func InvokeRpcUpdateCustomerRemark(ao *worktool_rpc.RequestAO, ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "UpdateCustomerRemark", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

func InvokeRpcQueryCommandResults(ao *worktool_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "QueryCommandResults", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcQueryCommandResultByMessageID(ao *worktool_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "QueryCommandResultByMessageID", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcSyncCommandResults(ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "SyncCommandResults", &worktool_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	if vo.Data == nil {
		return map[string]interface{}{}, nil
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

func InvokeRpcSaveCommandResults(ao *worktool_rpc.RequestAO, ctx context.Context) error {
	_, err := client.InvokeWorkToolRPCMethod(ctx, "SaveCommandResults", ao)
	return err
}

func InvokeRpcListPendingResults(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "ListPendingResults", &worktool_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcClearReplyNumber(ctx context.Context) error {
	_, err := client.InvokeWorkToolRPCMethod(ctx, "ClearReplyNumber", &worktool_rpc.RequestAO{})
	return err
}
