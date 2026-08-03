package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeWecomRPCQueryIndicatorCount(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorCount", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapInt(vo.Map)
	return result, nil
}

func InvokeWecomRPCQueryIndicatorDetail(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorDetail", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeWecomRPCCalculateUserIndicator(queryAO map[string]map[string]interface{}, ctx context.Context) (map[string]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "CalculateUserIndicator", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeWecomRPCQueryIndicatorCountInBatch(queryAO map[int][]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorCountInBatch", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeRPCQueryHalfDayLeaveMap(ctx context.Context) (map[string]bool, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryHalfDayLeaveMap", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapBool(vo.Map)
	return result, nil
}

func InvokeRPCGenerateReportRecord(queryAO map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(queryAO)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "GenerateReportRecord", ao)
	if err != nil {
		return err
	}
	return nil
}
