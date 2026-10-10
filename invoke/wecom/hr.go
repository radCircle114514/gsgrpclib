package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcGetHRFields fetches the employee field configuration.
func InvokeRpcGetHRFields(ctx context.Context) (map[string]interface{}, error) {
	ao := &wecom_rpc.RequestAO{}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetHRFields", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcGetHRStaffInfo fetches one employee's roster info.
//
// params uses the WeCom-native keys: userid (required), get_all (optional),
// fieldids: [{fieldid, sub_idx?}] (required unless get_all is true).
func InvokeRpcGetHRStaffInfo(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetHRStaffInfo", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcUpdateHRStaffInfo updates one employee's roster info.
//
// params uses the WeCom-native keys: userid (required), update_items,
// remove_items, insert_items (at least one required).
func InvokeRpcUpdateHRStaffInfo(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: toAny}
	vo, err := client.InvokeWecomRPCMethod(ctx, "UpdateHRStaffInfo", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
