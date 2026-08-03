package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRPCQueryIndicatorCount(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "QueryIndicatorCount", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapInt(vo.Map)
	return result, nil
}

func InvokeRPCQueryIndicatorDetail(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "QueryIndicatorDetail", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeBusinessRPCCalculateUserIndicator(queryAO map[string]map[string]interface{}, ctx context.Context) (map[string]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "CalculateUserIndicator", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeBusinessRPCQueryIndicatorCountInBatch(queryAO map[int][]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "QueryIndicatorCountInBatch", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}
