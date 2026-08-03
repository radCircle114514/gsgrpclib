package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRpcBusinessUpdateNotifyMode(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBusinessRPCMethod(ctx, "UpdateNotifyMode", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcBusinessUpdateCustomerRemark(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBusinessRPCMethod(ctx, "UpdateCustomerRemark", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcBusinessUpdateMiniAndAccount(spNo string, miniOpenIDList, accountOpenIDList []string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"spNo":              spNo,
		"miniOpenIDList":    miniOpenIDList,
		"accountOpenIDList": accountOpenIDList,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &business_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeBusinessRPCMethod(ctx, "UpdateMiniAndAccount", ao)
	if err != nil {
		return err
	}
	return nil
}
