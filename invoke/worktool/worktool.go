package worktool

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	worktool_rpc "github.com/iznilul/gsgrpclib/proto/worktool"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcListRobots(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "ListRobots", &worktool_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcGetRobot(robotKey string, ctx context.Context) (map[string]interface{}, error) {
	ao, err := buildMapAO(map[string]interface{}{"robotKey": robotKey})
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "GetRobot", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

func InvokeRpcFlushAllRobotCustomers(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "FlushAllRobotCustomers", &worktool_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcFlushRobotCustomers(robotKey string, ctx context.Context) (map[string]interface{}, error) {
	ao, err := buildMapAO(map[string]interface{}{"robotKey": robotKey})
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "FlushRobotCustomers", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

func InvokeRpcFlushAllRobotGroups(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "FlushAllRobotGroups", &worktool_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcFlushRobotGroup(chatID string, ctx context.Context) (map[string]interface{}, error) {
	ao, err := buildMapAO(map[string]interface{}{"chatID": chatID})
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "FlushRobotGroup", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Data), nil
}

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

func InvokeRpcListPendingResults(robotKey string, ctx context.Context) ([]map[string]interface{}, error) {
	ao, err := buildMapAO(map[string]interface{}{"robotKey": robotKey})
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWorkToolRPCMethod(ctx, "ListPendingResults", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcClearReplyNumber(ctx context.Context) error {
	_, err := client.InvokeWorkToolRPCMethod(ctx, "ClearReplyNumber", &worktool_rpc.RequestAO{})
	return err
}

func buildMapAO(data map[string]interface{}) (*worktool_rpc.RequestAO, error) {
	anyMap, err := utils.ParseMapToAny(data)
	if err != nil {
		return nil, err
	}
	return &worktool_rpc.RequestAO{Map: anyMap}, nil
}
