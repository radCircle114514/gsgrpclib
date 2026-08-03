package booking

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	booking_rpc "github.com/iznilul/gsgrpclib/proto/booking"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeBookingRPCQueryIndicatorCount(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorCount", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapInt(vo.Map)
	return result, nil
}

func InvokeBookingRPCQueryIndicatorDetail(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorDetail", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeBookingRPCCalculateUserIndicator(queryAO map[string]map[string]interface{}, ctx context.Context) (map[string]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "CalculateUserIndicator", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeBookingRPCQueryIndicatorCountInBatch(queryAO map[int][]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "QueryIndicatorCountInBatch", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}
