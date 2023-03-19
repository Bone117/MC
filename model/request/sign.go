package request

type SignRequest struct {
	JieCiId        uint   `json:"jieCiId"`
	WorkName       string `json:"workName"`              // 作品名称
	WorkFileTypeId uint   `json:"workFileTypeId"`        // 作品类型
	OtherAuthor    string `json:"otherAuthor,omitempty"` // 其他作者
	WorkAdviser    string `json:"workAdviser,omitempty"` // 指导老师
	WorkSoftware   string `json:"workSoftware"`          // 平台
	WorkDesc       string `json:"workDesc"`              // 作品简介
	MajorId        uint   `json:"majorId"`               // 专业id
	GradeName      string `json:"gradeName"`             // 班级
}

type UpdateSignRequest struct {
	WorkName       string `json:"workName"`              // 作品名称
	WorkFileTypeId uint   `json:"workFileTypeId"`        // 作品类型
	OtherAuthor    string `json:"otherAuthor,omitempty"` // 其他作者
	WorkAdviser    string `json:"workAdviser,omitempty"` // 指导老师
	WorkSoftware   string `json:"workSoftware"`          // 平台
	WorkDesc       string `json:"workDesc"`              // 作品简介
}
