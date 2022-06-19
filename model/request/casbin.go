package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/resetPassword", Method: "POST"},
		{Path: "/user/setSelfInfo", Method: "POST"},

		{Path: "/notice/getNotice", Method: "GET"},
		{Path: "/notice/getNoticeList", Method: "POST"},
	}
}
