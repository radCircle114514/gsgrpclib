package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRPCSaveCustomerPointRecord(data int, map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	dataToAny, err := utils.ParseDataToAny(data)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map:  toAny,
		Data: dataToAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SaveCustomerPointRecord", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetCustomerPointRecord(openID string, ctx context.Context) (interface{}, []map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(openID)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: data,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerPointRecord", ao)
	if err != nil {
		return nil, nil, err
	}
	point := utils.ParseAnyToData(vo.Data)
	list := utils.ParseAnyToMapList(vo.MapList)
	return point, list, nil
}

func InvokeRPCListCustomerPointRecord(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toMap, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toMap,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ListCustomerPointRecord", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

func InvokeRPCJudgeTodayIsWorkday(ctx context.Context) (bool, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "JudgeTodayIsWorkday", &wecom_rpc.RequestAO{})
	if err != nil {
		return false, err
	}
	isWorkday := utils.ParseAnyToData(vo.Data).(bool)
	return isWorkday, nil
}
