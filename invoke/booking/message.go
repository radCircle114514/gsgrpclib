package booking

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	booking_rpc "github.com/iznilul/gsgrpclib/proto/booking"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcBookingSendMiniMsg(toMap map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(toMap)
	if err != nil {
		return err
	}
	ao := &booking_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBookingRPCMethod(ctx, "SendMiniMsg", ao)
	if err != nil {
		return err
	}
	return nil
}
