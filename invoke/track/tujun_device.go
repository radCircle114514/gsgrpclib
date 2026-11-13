package track

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcTrackFindTujunUnsubscribedDevices returns active Tujun devices that
// have no live subscribe record. value filters the sn by substring and the
// result is paginated by currentPage/pageSize.
func InvokeRpcTrackFindTujunUnsubscribedDevices(value string, pageSize, currentPage int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	map1 := map[string]interface{}{
		"value":       value,
		"pageSize":    pageSize,
		"currentPage": currentPage,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindTujunUnsubscribedDevices", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

// InvokeRpcTrackFindTujunSubscribedDevices returns devices that currently have a
// live subscribe record (the "途骏-定位" tab SN search). value filters the sn by
// substring and the result is paginated by currentPage/pageSize.
func InvokeRpcTrackFindTujunSubscribedDevices(value string, pageSize, currentPage int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	map1 := map[string]interface{}{
		"value":       value,
		"pageSize":    pageSize,
		"currentPage": currentPage,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindTujunSubscribedDevices", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

// InvokeRpcTrackFindTujunHistorySubscribedDevices returns archived (finished)
// subscribed devices (the "途骏-历史定位" tab SN search). value filters the sn by
// substring and the result is paginated by currentPage/pageSize.
func InvokeRpcTrackFindTujunHistorySubscribedDevices(value string, pageSize, currentPage int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	map1 := map[string]interface{}{
		"value":       value,
		"pageSize":    pageSize,
		"currentPage": currentPage,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindTujunHistorySubscribedDevices", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}
