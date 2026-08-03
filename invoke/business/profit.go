package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

func InvokeRPCSyncOrderProfit(ctx context.Context) (map[int][]float64, error) {
	vo, err := client.InvokeBusinessRPCMethod(ctx, "SyncOrderProfit", &business_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntFloatList(vo.Map)
	return result, nil
}
