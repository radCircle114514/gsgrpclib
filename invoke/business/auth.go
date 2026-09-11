package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcWhoami returns the identity of the caller as resolved by the
// business server (typically from the auth token attached to the gRPC
// metadata). The server publishes the result as a single map in vo.Map,
// which is returned as map[string]interface{}.
func InvokeRpcWhoami(ctx context.Context) (map[string]interface{}, error) {
	ao := &business_rpc.RequestAO{}
	vo, err := client.InvokeBusinessRPCMethod(ctx, "Whoami", ao)
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
