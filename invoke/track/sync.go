package track

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/iznilul/gsgrpclib/client"
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/pkg/errors"
)

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
