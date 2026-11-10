package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcRegisterWeDriveSpace creates (or registers an externally created)
// WeDrive space and persists its metadata + auth info.
//
// params: space_name (required when creating), space_type (default
// "user_qrcode"), is_default (bool), creator_userid (optional), auth_info
// (optional [{type:1 userid auth} | {type:2 departmentid auth}]), spaceid
// (optional: when provided the space is assumed pre-created and only
// persisted).
func InvokeRpcRegisterWeDriveSpace(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "RegisterWeDriveSpace", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcDismissWeDriveSpace dismisses a space and marks its local record
// inactive. params: spaceid (required). Returns {spaceid, status}.
func InvokeRpcDismissWeDriveSpace(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "DismissWeDriveSpace", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcGetDefaultWeDriveSpace returns the active default space row for a
// space type. params: space_type (optional, default user_qrcode).
func InvokeRpcGetDefaultWeDriveSpace(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetDefaultWeDriveSpace", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcClearWeDriveSpaceFiles deletes every file in the space root and
// clears the local we_drive_file records. params: spaceid (required).
// Returns map {deleted_count}.
func InvokeRpcClearWeDriveSpaceFiles(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ClearWeDriveSpaceFiles", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcUploadWeDriveQRImage uploads one base64 PNG into the space and
// upserts the we_drive_file record.
//
// params: spaceid (required), file_name (required), file_base64_content
// (required, raw base64), user_id (required), qr_type (optional, default
// "unlimited"). Returns map {fileid, file_name}.
func InvokeRpcUploadWeDriveQRImage(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "UploadWeDriveQRImage", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
