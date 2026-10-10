package business

// ProcessFilterInAgent 是 Agent 侧查询流程实例列表的过滤条件，作为
// FindOngoingProcessInfoInAgent / FindFinishedProcessInfoInAgent 的
// Data（JSON）载荷。用户身份由服务端从 gRPC metadata 中的认证信息
// 解析，不在此传递，防止越权冒充。
type ProcessFilterInAgent struct {
	SerialNumber string `json:"serialNumber"`
	G3Number     string `json:"g3Number"`
	Title        string `json:"title"`
	CurrentPage  int    `json:"currentPage"`
	PageSize     int    `json:"pageSize"`
}
