package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcSyncAcademy(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncAcademy", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncTraining(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncTraining", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncCourse(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncCourse", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncInternshipPlan(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncInternshipPlan", ao)
	if err != nil {
		return err
	}
	return nil
}
