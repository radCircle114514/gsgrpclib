package track

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcTrackFindStationList(value string, pageSize, currentPage int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
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
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindStationList", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

func InvokeRpcTrackFindContentNoListBySerialNumberList(serialNumberList []string, ctx context.Context) ([]string, error) {
	jsonStr, _ := util.ToJSONStr(serialNumberList)
	anyList, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		DataList: anyList,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindContentNoListBySerialNumberList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	var contentNoList []string
	for _, data := range dataList {
		contentNoList = append(contentNoList, data.(string))
	}
	return contentNoList, nil
}
