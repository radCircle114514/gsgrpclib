package track

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcTrackFinishTujunLocation stops syncing one Tujun device (sn) and
// moves its subscribe record and stored location points into history.
func InvokeRpcTrackFinishTujunLocation(sn string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(sn)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	if _, err := client.InvokeTrackRPCMethod(ctx, "FinishTujunLocation", ao); err != nil {
		return err
	}
	return nil
}

// InvokeRpcTrackSyncTujunLocation manually triggers a single-device location
// pull for the given sn.
func InvokeRpcTrackSyncTujunLocation(sn string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(sn)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	if _, err := client.InvokeTrackRPCMethod(ctx, "SyncTujunLocation", ao); err != nil {
		return err
	}
	return nil
}

// InvokeRpcTrackFindTujunLocationList returns stored location points for one
// sn. The query map carries mode (ongoing|history), optional startTime/endTime
// and limit.
func InvokeRpcTrackFindTujunLocationList(query map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(query)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindTujunLocationList", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}
