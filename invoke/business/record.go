package business

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	business_rpc "github.com/iznilul/gsgrpclib/proto/business"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeGenerateRecordInMonth(userIDList []string, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(userIDList)
	anyList, _ := utils.ParseJsonStrToAnyList(jsonStr)
	ao := &business_rpc.RequestAO{
		DataList: anyList,
	}
	_, err := client.InvokeBusinessRPCMethod(ctx, "GenerateRecordInMonth", ao)
	if err != nil {
		return err
	}
	return nil
}
