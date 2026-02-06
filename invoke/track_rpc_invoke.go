package invoke

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
	"github.com/pkg/errors"
)

func GenerateRequestAO(data interface{}, dataList []interface{}, map1 map[string]interface{}, mapList []map[string]interface{}) *track_rpc.RequestAO {
	ao := &track_rpc.RequestAO{}
	if data != nil {
		result, _ := utils.ParseDataToAny(data)
		ao.Data = result
	}
	if dataList != nil {
		result, _ := utils.ParseDataListToAnyList(dataList)
		ao.DataList = result
	}
	if map1 != nil {
		result, _ := utils.ParseMapToAny(map1)
		ao.Map = result
	}
	if mapList != nil {
		result, _ := utils.ParseMapListToAnyList(mapList)
		ao.MapList = result
	}
	return ao
}

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

func InvokeRpcTrackQueryFieldList(table, field, value string, ctx context.Context) ([]interface{}, error) {
	map1 := map[string]interface{}{
		"table": table,
		"field": field,
		"value": value,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "QueryFieldList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return dataList, nil
}

func InvokeRpcTrackSendSyncRequest(comment map[string]interface{}, userID string, inst map[string]string, ctx context.Context) error {
	// 判断这个任务节点是否包括API请求的字段
	startStation := comment["task_start_station_autocomplete"]
	if startStation == nil {
		return nil
	}
	endStation := comment["task_end_station_autocomplete"]
	if endStation == nil {
		return nil
	}
	departureDirection := comment["departure_direction"]
	if departureDirection == nil {
		return nil
	}

	var map1 map[string]interface{}

	globalVar := inst["globalVar"]
	if globalVar == "" {
		return errors.New("globalVar is empty")
	}
	err := json.Unmarshal([]byte(globalVar), &map1)
	if err != nil {
		return err
	}

	var contentNoList []string
	startStationStr := startStation.(string)
	endStationStr := endStation.(string)
	departureDirectionStr := departureDirection.(string)
	if departureDirectionStr == "哈萨克斯坦" {
		if inst["procDefName"] == "铁路模板" {
			boxIDListString := map1["box_id"].(string)
			boxIDList := strings.Split(boxIDListString, ";")
			contentNoList = append(contentNoList, boxIDList...)
		} else if inst["procDefName"] == "车皮模板" {
			changeCarIDListString := map1["change_car_id"].(string)
			changeCarIDList := strings.Split(changeCarIDListString, ";")
			contentNoList = append(contentNoList, changeCarIDList...)
		}
	} else if departureDirectionStr == "俄罗斯" || departureDirectionStr == "蒙古" {
		changeCarIDListString := map1["change_car_id"].(string)
		changeCarIDList := strings.Split(changeCarIDListString, ";")
		if len(changeCarIDList) == 0 {
			return errors.New("换装车号不能为空")
		}
		var repeatedIDMap = make(map[string]bool)
		for _, changeCarID := range changeCarIDList {
			if !repeatedIDMap[changeCarID] {
				repeatedIDMap[changeCarID] = true
				contentNoList = append(contentNoList, changeCarID)
			}
		}
	}

	map1 = map[string]interface{}{
		"contentNoList": contentNoList,
		"startStation":  startStationStr,
		"endStation":    endStationStr,
		"serialNumber":  inst["serialNumber"],
		"userID":        userID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}

	_, err = client.InvokeTrackRPCMethod(ctx, "SendSyncRequest", ao)
	if err != nil {
		return err
	}
	return nil
}

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

func InvokeRpcTrackQueryDataInTimeScope(table, column string, startTime, endTime *time.Time, filter map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	queryAO := map[string]interface{}{
		"table":  table,
		"column": column,
		"filter": filter,
	}
	if startTime != nil {
		queryAO["startTime"] = *startTime
	}
	if endTime != nil {
		queryAO["endTime"] = *endTime
	}
	toAny, err := utils.ParseMapToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &track_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeTrackRPCMethod(ctx, "QueryDataInTimeScope", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
