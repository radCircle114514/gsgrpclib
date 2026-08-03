package track

import (
	track_rpc "github.com/iznilul/gsgrpclib/proto/track"
	"github.com/iznilul/gsgrpclib/utils"
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
