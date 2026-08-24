package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcBusinessFindProcInstByChatID(chatID string, ctx context.Context) ([]map[string]interface{}, error) {
	toAny, _ := utils.ParseDataToAny(chatID)
	ao := &business_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindProcInstByChatID", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcFindOrderInfo(openID string, currentPage, pageSize int, ctx context.Context) (*business_rpc.ResponseVO, interface{}, error) {
	map1 := map[string]interface{}{
		"openID":      openID,
		"currentPage": currentPage,
		"pageSize":    pageSize,
	}
	toAny, _ := utils.ParseMapToAny(map1)
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindOrderInfo", ao)
	if err != nil {
		return nil, 0, err
	}

	vos := vo.ProcInstVOList
	if len(vos) == 0 {
		return nil, 0, nil
	}
	count := utils.ParseAnyToData(vo.Data)
	return vo, count, nil
}

func InvokeRpcGenerateOrder(inquiry map[string]interface{}, ctx context.Context) error {
	toAny, _ := utils.ParseMapToAny(inquiry)
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	_, err := client.InvokeBusinessRPCMethod(ctx, "GenerateOrder", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcFindContentNoList(orderID string, ctx context.Context) ([]map[string]string, error) {
	toAny, _ := utils.ParseDataToAny(orderID)
	ao := &business_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindContentNoList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	var contentNoList []map[string]string
	for _, data := range dataList {
		map1 := map[string]string{
			"value": data.(string),
			"text":  data.(string),
		}
		contentNoList = append(contentNoList, map1)
	}
	return contentNoList, nil
}

func InvokeRpcUpdateTrack(serialNumber, userID, trackMsg string, validTrackMsgList []map[string]string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"serialNumber":      serialNumber,
		"userID":            userID,
		"track":             trackMsg,
		"validTrackMsgList": validTrackMsgList,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBusinessRPCMethod(ctx, "UpdateTrack", ao)
	return err
}

// OrderFilterInAgent mirrors the filter bound by v1.FindOrder in
// business-contact-back, but without the login/data-permission fields that
// do not apply to agent callers.
type OrderFilterInAgent struct {
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
	Topic        string `json:"topic"`
	G3Number     string `json:"g3Number"`
	SerialNumber string `json:"serialNumber"`
	Status       int    `json:"status"`
}

func InvokeRpcFindOrderInfoInAgent(filter OrderFilterInAgent, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindOrderInfoInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}
