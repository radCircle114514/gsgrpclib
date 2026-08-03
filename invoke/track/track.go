package track

import (
	"context"
	"fmt"
	"strings"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRPCTrackFindTrackList(contentNo string, ctx context.Context) (data map[string]interface{}, err error) {
	requestAO := GenerateRequestAO(contentNo, nil, nil, nil)

	var vo *track_rpc.ResponseVO
	vo, err = client.InvokeTrackRPCMethod(ctx, "FindTrackList", requestAO)
	if err != nil {
		return nil, err
	}
	var trackList []map[string]interface{}
	var markerList []map[string]interface{}
	dataList := utils.ParseAnyToMapList(vo.MapList)
	if len(dataList) == 0 {
		return nil, nil
	} else {
		track := dataList[0]
		startStationPosition := track["StartStationPosition"].(string)
		endStationPosition := track["EndStationPosition"].(string)
		startSplit := strings.Split(startStationPosition, ",")
		endSplit := strings.Split(endStationPosition, ",")
		start := map[string]interface{}{
			"start":       true,
			"stationName": fmt.Sprintf("起运站:%s", track["StartStationName"]),
			"longitude":   startSplit[0],
			"latitude":    startSplit[1],
		}
		end := map[string]interface{}{
			"end":         true,
			"stationName": fmt.Sprintf("目的站:%s", track["EndStationName"]),
			"longitude":   endSplit[0],
			"latitude":    endSplit[1],
		}
		markerList = append(markerList, start)
		for index, data := range dataList {
			position := data["CurrentStationPosition"].(string)
			split := strings.Split(position, ",")
			marker := map[string]interface{}{
				"index":       index + 1,
				"stationName": fmt.Sprintf("%d.当前站:%s\n距离目的地:%skm\n", index+1, data["CurrentStationName"], data["DistanceEnd"]),
				"longitude":   split[0],
				"latitude":    split[1],
			}
			if index == len(dataList)-1 {
				marker["newest"] = true
			}
			markerList = append(markerList, marker)
			track := map[string]interface{}{
				"index":              index + 1,
				"currentStation":     data["CurrentStation"],
				"currentStationName": data["CurrentStationName"],
				"distanceEnd":        data["DistanceEnd"],
				"operation":          data["Operation"],
				"createTime":         data["CreateTime"],
			}
			trackList = append(trackList, track)
		}
		markerList = append(markerList, end)
		trackList = utils.ReverseMapArray(trackList)
		data := map[string]interface{}{
			"trackList":  trackList,
			"markerList": markerList,
		}
		return data, nil
	}
}

func InvokeRpcTrackFindOngoingTrackList(ao *track_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindOngoingTrackList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcTrackFindHistoryTrackList(ao *track_rpc.RequestAO, ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeTrackRPCMethod(ctx, "FindHistoryTrackList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcTrackFinishTrack(contentNo string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(contentNo)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	_, err := client.InvokeTrackRPCMethod(ctx, "FinishTrack", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcTrackSyncTrack(contentNo string, ctx context.Context) error {
	data, _ := utils.ParseDataToAny(contentNo)
	ao := &track_rpc.RequestAO{
		Data: data,
	}
	_, err := client.InvokeTrackRPCMethod(ctx, "SyncTrack", ao)
	if err != nil {
		return err
	}
	return nil
}
