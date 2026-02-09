package invoke

import (
	"context"
	"time"

	"github.com/iznilul/gsgrpclib/client"
	wecom_rpc "github.com/iznilul/gsgrpclib/proto/wecom"
	"github.com/iznilul/gsgrpclib/utils"
	"github.com/mumushuiding/util"
)

func InvokeRpcGetUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}

func InvokeRpcGetAllUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAllUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}

func InvokeRpcGetUserInfo(code string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(code)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserInfo", ao)
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcGetRoleList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetRoleList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetRoleByUserID(userID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetRoleByUserID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetUserByRoleName(roleName string, ctx context.Context) ([]interface{}, error) {
	data, err := utils.ParseDataToAny(roleName)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserByRoleName", ao)
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToDataList(vo.DataList)
	return userList, nil
}

func InvokeRpcGetUserDetailList(userID string, dataList []interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	data, _ := utils.ParseDataToAny(userID)
	result, _ := utils.ParseDataListToAnyList(dataList)
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserDetailList", &wecom_rpc.RequestAO{
		Data:     data,
		DataList: result,
	})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetUserDetailByUserID(userID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetUserDetailByUserID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcSendTextMsg(requestUserID, userID, msg string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"requestUserID": requestUserID,
		"userID":        userID,
		"message":       msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendTextMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendWarnMsg(userID, msg string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"userID":  userID,
		"message": msg,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendWarnMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcFindCustomerList(remark, userID, searcher string, currentPage, pageSize int, all bool, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindCustomerAO{
		Remark:      remark,
		UserID:      userID,
		Searcher:    searcher,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		All:         all,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.FindCustomerList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	dataList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return dataList, count, nil
}

func InvokeRpcFindCustomerGroupList(name, userID, searcher string, currentPage, pageSize int, all bool, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindCustomerGroupAO{
		Name:        name,
		UserID:      userID,
		Searcher:    searcher,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		All:         all,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.FindCustomerGroupList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	dataList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return dataList, count, nil
}

func InvokeRpcAddCustomer(name, userID, markName, phone string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.AddCustomerAO{
		Name:     name,
		UserID:   userID,
		MarkName: markName,
		Phone:    phone,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	_, err = rpcClient.AddCustomer(ctx, ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcFindCustomerCouldBeSelectedList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindCustomerCouldBeSelectedList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcUpdateCustomerRemark(userID, markName, oldMarkName string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.UpdateCustomerAO{
		UserID:      userID,
		MarkName:    markName,
		OldMarkName: oldMarkName,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	_, err = rpcClient.UpdateCustomerRemark(ctx, ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetCustomerInRobotChat(externalUserID, remark string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerInRobotChatAO{
		ExternalUserID: externalUserID,
		Remark:         remark,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.GetCustomerInRobotChat(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcFlushCustomer(userID string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	res, err := client.InvokeWecomRPCMethod(ctx, "FlushCustomer", ao)
	if err != nil {
		return "", err
	}
	return res.Msg, nil
}

func InvokeRpcSyncCustomerGroup(userID string, ctx context.Context) error {
	data, err := utils.ParseDataToAny(userID)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncCustomerGroup", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcFlushCustomerGroup(chatID string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(chatID)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	res, err := client.InvokeWecomRPCMethod(ctx, "FlushCustomerGroup", ao)
	if err != nil {
		return "", err
	}
	return res.Msg, nil
}

func InvokeRpcQueryCustomerGroupByCond(table, field, value string, flag bool, ctx context.Context) ([]map[string]interface{}, []interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.OptionAO{
		Table: table,
		Field: field,
		Value: value,
		Flag:  flag,
	}
	ctx, cancel := client.SetTimeout(ctx)
	defer cancel()

	vo, err := rpcClient.QueryCustomerGroupByCond(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return mapList, dataList, nil
}

func InvokeRpcGetCustomerGroup(chatID, groupName string, ctx context.Context) (map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerGroupAO{
		ChatID:    chatID,
		GroupName: groupName,
	}
	vo, err := rpcClient.GetCustomerGroup(ctx, ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetCustomerGroupList(chatID string, ctx context.Context) ([]map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.GetCustomerGroupAO{
		ChatID: chatID,
	}
	vo, err := rpcClient.GetCustomerGroupList(ctx, ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetCustomerGroupChat(name string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(name)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerGroupChat", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}

func InvokeRpcGetCustomerGroupChatByChatID(chatID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(chatID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerGroupChatByChatID", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcFindAccountUserList(remark string, currentPage, pageSize, tagID int, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.FindAccountAO{
		Remark:      remark,
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
		TagID:       int32(tagID),
	}
	vo, err := rpcClient.FindAccountUserList(ctx, ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcSyncAccountUser(ctx context.Context) (string, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "SyncAccountUser", &wecom_rpc.RequestAO{})
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcFlushAccountUser(openID string, ctx context.Context) error {
	data, err := utils.ParseDataToAny(openID)
	if err != nil {
		return nil
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	_, err = client.InvokeWecomRPCMethod(ctx, "FlushAccountUser", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcGetAccountUserList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAccountUserList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetAccountUserListByOpenIDList(openIDList []string, ctx context.Context) ([]map[string]interface{}, error) {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.OpenIDListAO{
		OpenIDList: openIDList,
	}
	vo, err := rpcClient.GetAccountUserListByOpenIDList(ctx, ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcModifyAccountUserRemark(openID, remark string, ctx context.Context) error {
	conn, err := client.InitWecomRpcClientConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	rpcClient := wecom_rpc.NewWecomRPCClient(conn)
	ao := &wecom_rpc.ModifyRemarkAO{
		OpenID: openID,
		Remark: remark,
	}
	_, err = rpcClient.ModifyAccountUserRemark(ctx, ao)
	if err != nil {
		return nil
	}
	return nil
}

func InvokeRpcGetAccountTagList(ctx context.Context) ([]map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetAccountTagList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRpcGetOauthCallBackURL(id string, ctx context.Context) (string, error) {
	data, err := utils.ParseDataToAny(id)
	if err != nil {
		return "", err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetOauthCallBackURL", ao)
	if err != nil {
		return "", err
	}
	return vo.Msg, nil
}

func InvokeRpcGetFindEnumList(type1 string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(type1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Data: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindEnumList", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}

func InvokeRpcSendAccountTrackMsg(openID, contentNO, track string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"contentNO": contentNO,
		"track":     track,
		"openID":    openID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendAccountTrackMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSendAccountProcMsg(openID, item, dest, taskName, status, serialNumber string, ctx context.Context) error {
	map1 := map[string]interface{}{
		"item":         item,
		"dest":         dest,
		"taskName":     taskName,
		"status":       status,
		"serialNumber": serialNumber,
		"openID":       openID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SendAccountProcMsg", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcWecomQueryFieldList(table, field, value, userID string, ctx context.Context) ([]interface{}, error) {
	map1 := map[string]interface{}{
		"table":  table,
		"field":  field,
		"value":  value,
		"userID": userID,
	}
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryFieldList", ao)
	if err != nil {
		return nil, err
	}
	dataList := utils.ParseAnyToDataList(vo.DataList)
	return dataList, nil
}

func InvokeRpcWecomFindSupplierList(map2 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseMapToAny(map2)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindSupplierList", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcWecomSyncSupplier(ctx context.Context) error {
	_, err := client.InvokeWecomRPCMethod(ctx, "SyncSupplier", &wecom_rpc.RequestAO{})
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcUpdateSupplier(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateSupplier", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcWecomFindBusinessCustomerList(map2 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toAny, err := utils.ParseMapToAny(map2)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindBusinessCustomerList", ao)
	if err != nil {
		return nil, nil, err
	}

	mapList := utils.ParseAnyToMapList(vo.MapList)
	count := utils.ParseAnyToData(vo.Data)
	return mapList, count, nil
}

func InvokeRpcWecomSyncBusinessCustomer(ctx context.Context) error {
	_, err := client.InvokeWecomRPCMethod(ctx, "SyncBusinessCustomer", &wecom_rpc.RequestAO{})
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcWecomUpdateBusinessCustomer(map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateBusinessCustomer", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcQueryNameByType(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	data, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{Map: data}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryNameByType", ao)
	if err != nil {
		return nil, err
	}
	userList := utils.ParseAnyToMapList(vo.MapList)
	return userList, nil
}

func InvokeRpcGetBusinessCustomer(spNo string, ctx context.Context) (map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(spNo)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetBusinessCustomer", ao)
	if err != nil {
		return nil, err
	}
	toMap := utils.ParseAnyToMap(vo.Map)
	return toMap, nil
}

func InvokeRPCGetNeedNotifyCustomerList(spNoList []string, ctx context.Context) ([]map[string]interface{}, error) {
	jsonStr, _ := util.ToJSONStr(spNoList)
	anyList, _ := utils.ParseJsonStrToAnyList(jsonStr)
	ao := &wecom_rpc.RequestAO{
		DataList: anyList,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetNeedNotifyCustomerList", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}

func InvokeRPCFindClaimedCustomer(ctx context.Context) (map[string]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "FindClaimedCustomer", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRPCSaveCustomerPointRecord(data int, map1 map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(map1)
	if err != nil {
		return err
	}
	dataToAny, err := utils.ParseDataToAny(data)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map:  toAny,
		Data: dataToAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SaveCustomerPointRecord", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncContactWayList(ctx context.Context) ([]interface{}, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "SyncContactWayList", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	list := utils.ParseAnyToDataList(vo.DataList)
	return list, nil
}

func InvokeRpcFlushContactWay(configID string, ctx context.Context) (map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(configID)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: data,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "FlushContactWay", ao)
	if err != nil {
		return nil, err
	}
	map1 := utils.ParseAnyToMap(vo.Map)
	return map1, nil
}

func InvokeRpcGetCustomerPointRecord(openID string, ctx context.Context) (interface{}, []map[string]interface{}, error) {
	data, err := utils.ParseDataToAny(openID)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Data: data,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "GetCustomerPointRecord", ao)
	if err != nil {
		return nil, nil, err
	}
	point := utils.ParseAnyToData(vo.Data)
	list := utils.ParseAnyToMapList(vo.MapList)
	return point, list, nil
}

func InvokeRPCListCustomerPointRecord(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toMap, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toMap,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ListCustomerPointRecord", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

func InvokeRPCListUserByCond(map1 map[string]interface{}, ctx context.Context) ([]map[string]interface{}, interface{}, error) {
	toMap, err := utils.ParseMapToAny(map1)
	if err != nil {
		return nil, nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toMap,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "ListUserByCond", ao)
	if err != nil {
		return nil, nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	data := utils.ParseAnyToData(vo.Data)
	return mapList, data, nil
}

func InvokeRPCJudgeTodayIsWorkday(ctx context.Context) (bool, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "JudgeTodayIsWorkday", &wecom_rpc.RequestAO{})
	if err != nil {
		return false, err
	}
	isWorkday := utils.ParseAnyToData(vo.Data).(bool)
	return isWorkday, nil
}

func InvokeWecomRPCQueryIndicatorCount(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorCount", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapInt(vo.Map)
	return result, nil
}

func InvokeWecomRPCQueryIndicatorDetail(queryAO map[int]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseMapIntToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorDetail", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeWecomRPCCalculateUserIndicator(queryAO map[string]map[string]interface{}, ctx context.Context) (map[string]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "CalculateUserIndicator", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeRPCQueryHalfDayLeaveMap(ctx context.Context) (map[string]bool, error) {
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryHalfDayLeaveMap", &wecom_rpc.RequestAO{})
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapBool(vo.Map)
	return result, nil
}

func InvokeRPCGenerateReportRecord(queryAO map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(queryAO)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "GenerateReportRecord", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncAcademy(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncAcademy", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeRpcSyncTraining(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncTraining", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeWecomRPCQueryIndicatorCountInBatch(queryAO map[int][]map[string]interface{}, ctx context.Context) (map[int][]map[string]interface{}, error) {
	toAny, err := utils.ParseDataToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryIndicatorCountInBatch", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapIntList(vo.Map)
	return result, nil
}

func InvokeRpcSyncCourse(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncCourse", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeQuerySupplierBySpNoList(spNoList []string, ctx context.Context) (map[string]map[string]interface{}, error) {
	jsonStr, _ := util.ToJSONStr(spNoList)
	anyList, _ := utils.ParseJsonStrToAnyList(jsonStr)
	ao := &wecom_rpc.RequestAO{
		DataList: anyList,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QuerySupplierBySpNoList", ao)
	if err != nil {
		return nil, err
	}
	result := utils.ParseAnyToMapStringMap(vo.Map)
	return result, nil
}

func InvokeRpcSyncInternshipPlan(mapList []map[string]interface{}, ctx context.Context) error {
	jsonStr, _ := util.ToJSONStr(mapList)
	toAny, err := utils.ParseJsonStrToAnyList(jsonStr)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		MapList: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "SyncInternshipPlan", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeWecomRpcUpdateUser(userInfo map[string]interface{}, ctx context.Context) error {
	toAny, err := utils.ParseMapToAny(userInfo)
	if err != nil {
		return err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	_, err = client.InvokeWecomRPCMethod(ctx, "UpdateUser", ao)
	if err != nil {
		return err
	}
	return nil
}

func InvokeWecomRpcQueryDataInTimeScope(table, column string, startTime, endTime *time.Time, filter map[string]interface{}, ctx context.Context) ([]map[string]interface{}, error) {
	queryAO := map[string]interface{}{
		"table":  table,
		"column": column,
		"filter": filter,
	}
	if startTime != nil {
		queryAO["startTime"] = *startTime
	}
	if endTime != nil {
		queryAO["endTime"] = *endTime
	}
	toAny, err := utils.ParseMapToAny(queryAO)
	if err != nil {
		return nil, err
	}
	ao := &wecom_rpc.RequestAO{
		Map: toAny,
	}
	vo, err := client.InvokeWecomRPCMethod(ctx, "QueryDataInTimeScope", ao)
	if err != nil {
		return nil, err
	}
	mapList := utils.ParseAnyToMapList(vo.MapList)
	return mapList, nil
}
