package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

func invokeChatbotMapList(method string, payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	anyMap, err := utils.ParseMapToAny(payload)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeBusinessRPCMethod(ctx, method, &business_rpc.RequestAO{Map: anyMap})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}

func InvokeRpcFindChatbotOngoingInsts(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotOngoingInsts", payload, ctx)
}

func InvokeRpcFindChatbotHistoryInsts(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotHistoryInsts", payload, ctx)
}

func InvokeRpcFindChatbotIdentityLinks(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotIdentityLinksByProcInstID", payload, ctx)
}

func InvokeRpcFindChatbotBoxInfos(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotBoxInfos", payload, ctx)
}

func InvokeRpcFindChatbotTraceInfos(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotTraceInfos", payload, ctx)
}

func InvokeRpcFindChatbotInspectionInfos(payload map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	return invokeChatbotMapList("FindChatbotInspectionInfos", payload, ctx)
}

func InvokeRpcFindChatbotFormItems(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeBusinessRPCMethod(ctx, "FindChatbotFormItems", &business_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMapList(vo.MapList), nil
}
