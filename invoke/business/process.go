package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// ProcessFilterInAgent mirrors the filter bound by v1.FindOngoingProcess /
// v1.FindFinishedProcess and their "All" variants. When UserID is empty the
// server behaves like FindAll* (no per-user scoping); when set it behaves
// like Find* for that user.
type ProcessFilterInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
	Title        string `json:"title"`
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
	UserID       string `json:"userID"`
}

func InvokeRpcFindOngoingProcessInfoInAgent(filter ProcessFilterInAgent, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindOngoingProcessInfoInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcFindFinishedProcessInfoInAgent(filter ProcessFilterInAgent, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseDataToAny(filter)
	if err != nil {
		return nil, 0, err
	}
	ao := &business_rpc.RequestAO{Data: toAny}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindFinishedProcessInfoInAgent", ao)
	if err != nil {
		return nil, 0, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}
