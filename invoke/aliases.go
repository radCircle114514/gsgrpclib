package invoke

import (
	"github.com/iznilul/gsgrpclib/invoke/booking"
	"github.com/iznilul/gsgrpclib/invoke/business"
	"github.com/iznilul/gsgrpclib/invoke/track"
	"github.com/iznilul/gsgrpclib/invoke/wecom"
)

// Re-exports for backwards compatibility. Prefer importing the subpackages directly.

type (
	OrderFilterInAgent   = business.OrderFilterInAgent
	ProcessFilterInAgent = business.ProcessFilterInAgent
)

var (
	InvokeBookingRPCCalculateUserIndicator            = booking.InvokeBookingRPCCalculateUserIndicator
	InvokeBookingRPCQueryDataInTimeScope              = booking.InvokeBookingRPCQueryDataInTimeScope
	InvokeBookingRPCQueryIndicatorCount               = booking.InvokeBookingRPCQueryIndicatorCount
	InvokeBookingRPCQueryIndicatorCountInBatch        = booking.InvokeBookingRPCQueryIndicatorCountInBatch
	InvokeBookingRPCQueryIndicatorDetail              = booking.InvokeBookingRPCQueryIndicatorDetail
	InvokeRPCGetGiftList                              = booking.InvokeRPCGetGiftList
	InvokeRpcBookingFindMiniUser                      = booking.InvokeRpcBookingFindMiniUser
	InvokeRpcBookingFindMiniUserList                  = booking.InvokeRpcBookingFindMiniUserList
	InvokeRpcBookingFindMiniUserListByOpenIDList      = booking.InvokeRpcBookingFindMiniUserListByOpenIDList
	InvokeRpcBookingSendMiniMsg                       = booking.InvokeRpcBookingSendMiniMsg
	InvokeBusinessRPCCalculateUserIndicator           = business.InvokeBusinessRPCCalculateUserIndicator
	InvokeBusinessRPCQueryIndicatorCountInBatch       = business.InvokeBusinessRPCQueryIndicatorCountInBatch
	InvokeGenerateRecordInMonth                       = business.InvokeGenerateRecordInMonth
	InvokeRPCQueryIndicatorCount                      = business.InvokeRPCQueryIndicatorCount
	InvokeRPCQueryIndicatorDetail                     = business.InvokeRPCQueryIndicatorDetail
	InvokeRPCSyncOrderProfit                          = business.InvokeRPCSyncOrderProfit
	InvokeRpcBusinessFindProcInstByChatID             = business.InvokeRpcBusinessFindProcInstByChatID
	InvokeRpcBusinessUpdateCustomerRemark             = business.InvokeRpcBusinessUpdateCustomerRemark
	InvokeRpcBusinessUpdateMiniAndAccount             = business.InvokeRpcBusinessUpdateMiniAndAccount
	InvokeRpcBusinessUpdateNotifyMode                 = business.InvokeRpcBusinessUpdateNotifyMode
	InvokeRpcFindContentNoList                        = business.InvokeRpcFindContentNoList
	InvokeRpcFindFinishedProcessInfoInAgent           = business.InvokeRpcFindFinishedProcessInfoInAgent
	InvokeRpcFindOngoingProcessInfoInAgent            = business.InvokeRpcFindOngoingProcessInfoInAgent
	InvokeRpcFindOrderInfo                            = business.InvokeRpcFindOrderInfo
	InvokeRpcFindOrderInfoInAgent                     = business.InvokeRpcFindOrderInfoInAgent
	InvokeRpcGenerateOrder                            = business.InvokeRpcGenerateOrder
	InvokeRpcUpdateTrack                              = business.InvokeRpcUpdateTrack
	GenerateRequestAO                                 = track.GenerateRequestAO
	InvokeRPCTrackFindTrackList                       = track.InvokeRPCTrackFindTrackList
	InvokeRpcTrackFindContentNoListBySerialNumberList = track.InvokeRpcTrackFindContentNoListBySerialNumberList
	InvokeRpcTrackFindHistoryTrackList                = track.InvokeRpcTrackFindHistoryTrackList
	InvokeRpcTrackFindOngoingTrackList                = track.InvokeRpcTrackFindOngoingTrackList
	InvokeRpcTrackFindStationList                     = track.InvokeRpcTrackFindStationList
	InvokeRpcTrackFinishTrack                         = track.InvokeRpcTrackFinishTrack
	InvokeRpcTrackQueryDataInTimeScope                = track.InvokeRpcTrackQueryDataInTimeScope
	InvokeRpcTrackQueryFieldList                      = track.InvokeRpcTrackQueryFieldList
	InvokeRpcTrackSendSyncRequest                     = track.InvokeRpcTrackSendSyncRequest
	InvokeRpcTrackSyncTrack                           = track.InvokeRpcTrackSyncTrack
	InvokeQuerySupplierBySpNoList                     = wecom.InvokeQuerySupplierBySpNoList
	InvokeRPCFindClaimedCustomer                      = wecom.InvokeRPCFindClaimedCustomer
	InvokeRPCGenerateReportRecord                     = wecom.InvokeRPCGenerateReportRecord
	InvokeRPCGetNeedNotifyCustomerList                = wecom.InvokeRPCGetNeedNotifyCustomerList
	InvokeRPCJudgeTodayIsWorkday                      = wecom.InvokeRPCJudgeTodayIsWorkday
	InvokeRPCListCustomerPointRecord                  = wecom.InvokeRPCListCustomerPointRecord
	InvokeRPCListUserByCond                           = wecom.InvokeRPCListUserByCond
	InvokeRPCQueryHalfDayLeaveMap                     = wecom.InvokeRPCQueryHalfDayLeaveMap
	InvokeRPCSaveCustomerPointRecord                  = wecom.InvokeRPCSaveCustomerPointRecord
	InvokeRpcAddCustomer                              = wecom.InvokeRpcAddCustomer
	InvokeRpcFindAccountUserList                      = wecom.InvokeRpcFindAccountUserList
	InvokeRpcFindCustomerCouldBeSelectedList          = wecom.InvokeRpcFindCustomerCouldBeSelectedList
	InvokeRpcFindCustomerGroupList                    = wecom.InvokeRpcFindCustomerGroupList
	InvokeRpcFindCustomerList                         = wecom.InvokeRpcFindCustomerList
	InvokeRpcFlushAccountUser                         = wecom.InvokeRpcFlushAccountUser
	InvokeRpcFlushContactWay                          = wecom.InvokeRpcFlushContactWay
	InvokeRpcFlushCustomer                            = wecom.InvokeRpcFlushCustomer
	InvokeRpcFlushCustomerGroup                       = wecom.InvokeRpcFlushCustomerGroup
	InvokeRpcGetAccountTagList                        = wecom.InvokeRpcGetAccountTagList
	InvokeRpcGetAccountUserList                       = wecom.InvokeRpcGetAccountUserList
	InvokeRpcGetAccountUserListByOpenIDList           = wecom.InvokeRpcGetAccountUserListByOpenIDList
	InvokeRpcGetAllUserList                           = wecom.InvokeRpcGetAllUserList
	InvokeRpcGetApprovalDetail                        = wecom.InvokeRpcGetApprovalDetail
	InvokeRpcGetApprovalList                          = wecom.InvokeRpcGetApprovalList
	InvokeRpcGetBusinessCustomer                      = wecom.InvokeRpcGetBusinessCustomer
	InvokeRpcGetCheckInData                           = wecom.InvokeRpcGetCheckInData
	InvokeRpcGetCustomerGroup                         = wecom.InvokeRpcGetCustomerGroup
	InvokeRpcGetCustomerGroupChat                     = wecom.InvokeRpcGetCustomerGroupChat
	InvokeRpcGetCustomerGroupChatByChatID             = wecom.InvokeRpcGetCustomerGroupChatByChatID
	InvokeRpcGetCustomerGroupList                     = wecom.InvokeRpcGetCustomerGroupList
	InvokeRpcGetCustomerInRobotChat                   = wecom.InvokeRpcGetCustomerInRobotChat
	InvokeRpcGetCustomerPointRecord                   = wecom.InvokeRpcGetCustomerPointRecord
	InvokeRpcGetDepartmentTree                        = wecom.InvokeRpcGetDepartmentTree
	InvokeRpcGetFindEnumMap                           = wecom.InvokeRpcGetFindEnumMap
	InvokeRpcGetOauthCallBackURL                      = wecom.InvokeRpcGetOauthCallBackURL
	InvokeRpcGetRoleByUserID                          = wecom.InvokeRpcGetRoleByUserID
	InvokeRpcGetRoleList                              = wecom.InvokeRpcGetRoleList
	InvokeRpcGetReportDetail                          = wecom.InvokeRpcGetReportDetail
	InvokeRpcGetReportList                            = wecom.InvokeRpcGetReportList
	InvokeRpcGetUserByRoleName                        = wecom.InvokeRpcGetUserByRoleName
	InvokeRpcGetUserDetailByUserID                    = wecom.InvokeRpcGetUserDetailByUserID
	InvokeRpcGetUserDetailList                        = wecom.InvokeRpcGetUserDetailList
	InvokeRpcGetUserIdentity                          = wecom.InvokeRpcGetUserIdentity
	InvokeRpcGetUserInfo                              = wecom.InvokeRpcGetUserInfo
	InvokeRpcGetUserList                              = wecom.InvokeRpcGetUserList
	InvokeRpcModifyAccountUserRemark                  = wecom.InvokeRpcModifyAccountUserRemark
	InvokeRpcQueryCustomerGroupByCond                 = wecom.InvokeRpcQueryCustomerGroupByCond
	InvokeRpcQueryNameByType                          = wecom.InvokeRpcQueryNameByType
	InvokeRpcSendAccountProcMsg                       = wecom.InvokeRpcSendAccountProcMsg
	InvokeRpcSendAccountTrackMsg                      = wecom.InvokeRpcSendAccountTrackMsg
	InvokeRpcSendTextMsg                              = wecom.InvokeRpcSendTextMsg
	InvokeRpcSendWarnMsg                              = wecom.InvokeRpcSendWarnMsg
	InvokeRpcSyncAcademy                              = wecom.InvokeRpcSyncAcademy
	InvokeRpcSyncAccountUser                          = wecom.InvokeRpcSyncAccountUser
	InvokeRpcSyncContactWayList                       = wecom.InvokeRpcSyncContactWayList
	InvokeRpcSyncCourse                               = wecom.InvokeRpcSyncCourse
	InvokeRpcSyncCustomerGroup                        = wecom.InvokeRpcSyncCustomerGroup
	InvokeRpcSyncInternshipPlan                       = wecom.InvokeRpcSyncInternshipPlan
	InvokeRpcSyncTraining                             = wecom.InvokeRpcSyncTraining
	InvokeRpcUpdateCustomerRemark                     = wecom.InvokeRpcUpdateCustomerRemark
	InvokeRpcUpdateSupplier                           = wecom.InvokeRpcUpdateSupplier
	InvokeRpcWecomFindBusinessCustomerList            = wecom.InvokeRpcWecomFindBusinessCustomerList
	InvokeRpcWecomFindSupplierList                    = wecom.InvokeRpcWecomFindSupplierList
	InvokeRpcWecomQueryFieldList                      = wecom.InvokeRpcWecomQueryFieldList
	InvokeRpcWecomSyncBusinessCustomer                = wecom.InvokeRpcWecomSyncBusinessCustomer
	InvokeRpcWecomSyncSupplier                        = wecom.InvokeRpcWecomSyncSupplier
	InvokeRpcWecomUpdateBusinessCustomer              = wecom.InvokeRpcWecomUpdateBusinessCustomer
	InvokeWecomRPCCalculateUserIndicator              = wecom.InvokeWecomRPCCalculateUserIndicator
	InvokeWecomRPCQueryIndicatorCount                 = wecom.InvokeWecomRPCQueryIndicatorCount
	InvokeWecomRPCQueryIndicatorCountInBatch          = wecom.InvokeWecomRPCQueryIndicatorCountInBatch
	InvokeWecomRPCQueryIndicatorDetail                = wecom.InvokeWecomRPCQueryIndicatorDetail
	InvokeWecomRpcQueryDataInTimeScope                = wecom.InvokeWecomRpcQueryDataInTimeScope
	InvokeWecomRpcUpdateUser                          = wecom.InvokeWecomRpcUpdateUser
)
