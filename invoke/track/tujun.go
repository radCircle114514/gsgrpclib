package track

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcTrackSubscribeTujun subscribes one or more Tujun (途君 GPS) devices
// via the tracking service over gRPC.
func InvokeRpcTrackSubscribeTujun(snList []string, serialNumber, userID string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"snList":       snList,
		"serialNumber": serialNumber,
		"userID":       userID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	if _, err := client.InvokeTrackRPCMethod(ctx, "SubscribeTujun", ao); err != nil {
		return err
	}
	return nil
}
