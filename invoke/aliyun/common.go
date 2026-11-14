package aliyun

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	aliyun_rpc "github.com/iznilul/gsgrpclib/proto/aliyun"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcDoCommonRequest calls an arbitrary Aliyun OpenAPI RPC-style action
// through the service's generic CommonRequest client.
//
//	domain    e.g. "ecs.cn-hangzhou.aliyuncs.com"
//	version   e.g. "2014-05-26"
//	apiAction e.g. "DescribeInstances"
//	method    "GET" or "POST" (defaults to GET when empty)
//	regionId  e.g. "cn-hangzhou" (defaults to the service's configured region)
//	params    action-specific query parameters
func InvokeRpcDoCommonRequest(ctx context.Context, domain, version, apiAction, method, regionId string, params map[string]interface{}) (map[string]interface{}, error) {
	req := map[string]interface{}{
		"domain":    domain,
		"version":   version,
		"apiAction": apiAction,
		"method":    method,
		"regionId":  regionId,
		"params":    params,
	}
	toAny, err := utils.ParseMapToAny(req)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeAliyunRPCMethod(ctx, "DoCommonRequest", &aliyun_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
