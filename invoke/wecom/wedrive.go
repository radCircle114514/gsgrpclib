package wecom

import (
	"context"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
)

// InvokeRpcCreateWeDriveSpace creates a WeDrive space.
//
// params: space_name (required), space_sub_type (optional, 0 only),
// auth_info (optional): [{type:1 userid auth} | {type:2 departmentid auth}],
// auth: 1 download-only, 4 preview, 7 admin (max 3, users only).
// Returns map containing spaceid.
func InvokeRpcCreateWeDriveSpace(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "CreateWeDriveSpace", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcAddWeDriveSpaceACL adds members/departments to a space.
//
// params: spaceid (required), auth_info (required) with the same shape as
// CreateWeDriveSpace.
func InvokeRpcAddWeDriveSpaceACL(params map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return err
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "AddWeDriveSpaceACL", &wecom_rpc.RequestAO{Map: toAny})
	return err
}

// InvokeRpcListWeDriveFiles lists files under one folder of a space.
//
// params: spaceid (required), fatherid (required, equals spaceid at root),
// sort_type (1-6), start (0 first page), limit (<=1000), fetch_all (optional
// bool: follow has_more/next_start and return every page aggregated).
// Returns map: {file_list: [...], has_more, next_start}.
func InvokeRpcListWeDriveFiles(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ListWeDriveFiles", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcRenameWeDriveFile renames a file/folder.
//
// params: fileid (required), new_name (required, max 255 chars, Chinese=2).
// Returns map containing the updated file object.
func InvokeRpcRenameWeDriveFile(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "RenameWeDriveFile", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcDeleteWeDriveFiles deletes one or more files (batch).
//
// params: fileids (required, []string).
func InvokeRpcDeleteWeDriveFiles(params map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return err
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "DeleteWeDriveFiles", &wecom_rpc.RequestAO{Map: toAny})
	return err
}

// InvokeRpcUploadWeDriveFile uploads a small file (<=10MB) in one base64 call.
//
// params: spaceid, fatherid (or selected_ticket), file_name (required),
// file_base64_content (required, raw base64 without data-URI prefix).
// Returns map containing fileid.
func InvokeRpcUploadWeDriveFile(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "UploadWeDriveFile", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcInitWeDriveFileUpload starts a chunked upload (files up to 20GB,
// 2MiB chunks).
//
// params: spaceid, fatherid (or selected_ticket), file_name, size (bytes),
// block_sha ([]string cumulative SHA-1 per chunk), skip_push_card (optional).
// Returns map: {hit_exist, upload_key, fileid}. When hit_exist is true the
// upload is already finished and upload_key is empty.
func InvokeRpcInitWeDriveFileUpload(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "InitWeDriveFileUpload", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}

// InvokeRpcUploadWeDriveFilePart uploads one 2MiB chunk.
//
// params: upload_key (required), index (required, starts at 1),
// file_base64_content (required).
func InvokeRpcUploadWeDriveFilePart(params map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return err
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UploadWeDriveFilePart", &wecom_rpc.RequestAO{Map: toAny})
	return err
}

// InvokeRpcFinishWeDriveFileUpload completes a chunked upload.
//
// params: upload_key (required). Returns map containing fileid.
func InvokeRpcFinishWeDriveFileUpload(params map[string]interface{}, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseMapToAny(params)
	if err != nil {
		return nil, err
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FinishWeDriveFileUpload", &wecom_rpc.RequestAO{Map: toAny})
	if err != nil {
		return nil, err
	}
	return utils.ParseAnyToMap(vo.Map), nil
}
