package aliyun

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	aliyun_rpc "github.com/iznilul/gsgrpclib/proto/aliyun"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcSendSms sends a templated SMS. templateParam carries the template
// variables (e.g. {"number": "123456"}). When signName/templateCode are left
// empty the server falls back to its configured defaults.
func InvokeRpcSendSms(ctx context.Context, phoneNumbers, signName, templateCode string, templateParam map[string]interface{}) (map[string]interface{}, error) {
	req := map[string]interface{}{
		"phoneNumbers":  phoneNumbers,
		"signName":      signName,
		"templateCode":  templateCode,
		"templateParam": templateParam,
	}
	toAny, err := utils.ParseMapToAny(req)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeAliyunRPCMethod(ctx, "SendSms", &aliyun_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
