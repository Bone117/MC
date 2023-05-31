package model

import "server/global"

type ReportWithUsername struct {
	Report
	Username string `json:"username"`
}
type SignWithPhone struct {
	Sign
	Phone string `json:"phone"`
}

type Sign struct {
	global.MODEL
	UserId             uint               `json:"userId" gorm:"not null;comment:用户id"`             // 用户id
	JieCiId            uint               `json:"jieCiId"`                                         // 届次
	WorkName           string             `json:"workName" gorm:"not null;comment:作品名称"`           // 作品名称
	WorkFileTypeId     uint               `json:"workFileTypeId" gorm:"not null;comment:作品类型"`     // 作品类型
	NickName           string             `json:"nickName" gorm:"comment:第一作者"`                    // 第一作者
	Username           string             `json:"username" gorm:"comment:学号"`                      // 学号
	OtherAuthor        string             `json:"otherAuthor" gorm:"comment:其他作者"`                 // 其他作者
	WorkAdviser        string             `json:"workAdviser" gorm:"comment:指导老师"`                 // 指导老师
	WorkSoftware       string             `json:"workSoftware" gorm:"comment:平台"`                  // 平台
	Status             uint               `json:"status" gorm:"default:0;comment:审核状态" `           // 1.已通过 2.已拒绝 3.待发布 4.已发布 5.审核中 10其它->待审核
	WorkDesc           string             `json:"workDesc" gorm:"type:text;not null;comment:作品简介"` // 作品简介
	CoverUrl           string             `json:"coverUrl"`                                        // 缩略图
	RejReason          string             `json:"rejReason" gorm:"type:text;;comment:不通过理由"`       // 审核不通过理由
	Likes              uint               `json:"likes" gorm:"default:0;comment:点赞数" `
	Files              []File             `gorm:"foreignkey:SignId"` // 定义关联关系
	Evaluates          []Evaluate         `gorm:"foreignkey:SignId;references:ID"`
	ReportWithUsername ReportWithUsername `json:"reportWithUsername" gorm:"-"`
}
