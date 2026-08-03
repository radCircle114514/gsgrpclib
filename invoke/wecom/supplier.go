package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcWecomFindSupplierList(map2 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseMapToAny(map2)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindSupplierList", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcWecomSyncSupplier(ctx context.Context) error {
	_, err := client.InvokeWecomRPCMethod(ctx, "SyncSupplier", &wecom_rpc.RequestAO{})
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcUpdateSupplier(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateSupplier", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeQuerySupplierBySpNoList(spNoList []string, ctx context.Context) (map[string]map[string]interface{}, error) {
	jsonStr, _ := util.ToJSONStr(spNoList)
	anyList, _ := utils.ParseJsonStrToAnyList(jsonStr)
	ao := &wecom_rpc.RequestAO{
		DataList: anyList,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QuerySupplierBySpNoList", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeRpcWecomFindBusinessCustomerList(map2 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseMapToAny(map2)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindBusinessCustomerList", ao)
	if err != nil {
		return nil, nil, err
	}

	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcWecomSyncBusinessCustomer(ctx context.Context) error {
	_, err := client.InvokeWecomRPCMethod(ctx, "SyncBusinessCustomer", &wecom_rpc.RequestAO{})
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcWecomUpdateBusinessCustomer(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateBusinessCustomer", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetBusinessCustomer(spNo string, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(spNo)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetBusinessCustomer", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}
