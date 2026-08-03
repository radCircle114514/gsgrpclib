package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcSyncContactWayList(ctx context.Context) ([]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "SyncContactWayList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	list := utils.ParseAnyToDataList(vo.DataList)
	return list, nil
}

func InvokeRpcFlushContactWay(configID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(configID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: data,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FlushContactWay", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcSendTextMsg(requestUserID, userID, msg string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"requestUserID": requestUserID,
		"userID":        userID,
		"message":       msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendTextMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendWarnMsg(userID, msg string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"userID":  userID,
		"message": msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendWarnMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendAccountTrackMsg(openID, contentNO, track string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"contentNO": contentNO,
		"track":     track,
		"openID":    openID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendAccountTrackMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendAccountProcMsg(openID, item, dest, taskName, status, serialNumber string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"item":         item,
		"dest":         dest,
		"taskName":     taskName,
		"status":       status,
		"serialNumber": serialNumber,
		"openID":       openID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendAccountProcMsg", ao)
	if err != nil {
		return err
	}
	return nil
}
