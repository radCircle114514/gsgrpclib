package booking

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	booking_rpc "github.com/iznilul/gsgrpclib/proto/booking"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRPCGetGiftList(index string, ctx context.Context) ([]map[string]interface{}, error) {
	toAny, _ := utils.ParseDataToAny(index)
	ao := &booking_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBookingRPCMethod(ctx, "GetGiftList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
