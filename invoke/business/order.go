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

// OrderFilterInAgent 是 Agent 侧查询订单列表的过滤条件，作为
// FindOrderListInAgent 的 Data（JSON）载荷。用户身份不在此传递，
// 由服务端从 gRPC metadata 中的认证信息解析，防止越权冒充。
type OrderFilterInAgent struct {
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
	Topic        string `json:"topic"`
	G3Number     string `json:"g3Number"`
	SerialNumber string `json:"serialNumber"`
	Status       int    `json:"status"`
}

// InvokeRpcFindOrderListInAgent 供 Agent 调用，按过滤条件分页查询当前
// 认证用户可见的订单列表，返回订单行与总条数。行级数据权限由服务端
// 依据 metadata 中的 userID 对应权限校验，调用方无法自行扩大可见范围。
func InvokeRpcFindOrderListInAgent(filter OrderFilterInAgent, ctx context.Context) ([]map[string]interface{}, int64, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindOrderListInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count, _ := utils.ParseAnyToData(vo.Data).(int64)
	return mapList, count, nil
}

// OrderDetailInAgent 是 Agent 侧查询订单详情的请求参数，作为
// DetailOrderInfoInAgent 的 Data（JSON）载荷。用户身份由服务端从
// gRPC metadata 中的认证信息解析，不在此传递，防止越权冒充。
// 三个定位字段优先级为 SerialNumber > G3Number；SerialNumber 支持
// 只传序列号末尾 6 位（如 "150405"），服务端按后缀匹配。
type OrderDetailInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
}

// InvokeRpcDetailOrderInfoInAgent 供 Agent 调用，按序列号（支持末 6 位
// 后缀匹配）或 G3 单号定位当前认证用户可见的订单，返回订单详情
// （含附件、更新记录、绑定用户等），详情结构与 HTTP /v1/order/:id
// 接口一致，位于响应的 Map 字段。
func InvokeRpcDetailOrderInfoInAgent(filter OrderDetailInAgent, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, err
	}
	ao := &business_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "DetailOrderInfoInAgent", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
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
