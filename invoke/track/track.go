package track

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcTrackFindWcwOngoingTrackList(ao *track_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindWcwOngoingTrackList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcTrackFindWcwHistoryTrackList(ao *track_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindWcwHistoryTrackList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcTrackFinishWcwTrack(contentNo string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(contentNo)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	_, err := client.InvokeTrackRPCMethod(ctx, "FinishWcwTrack", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcTrackSyncWcwTrack(contentNo string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(contentNo)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	_, err := client.InvokeTrackRPCMethod(ctx, "SyncWcwTrack", ao)
	if err != nil {
		return err
	}
	return nil
}
