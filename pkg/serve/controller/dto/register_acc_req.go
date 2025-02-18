package dto

type RegisterAccReq struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required,email"`
	NikeName string `json:"nickName" xml:"nickName" form:"nickName" validate:"required"`
	Password string `json:"password" xml:"password" form:"password" validate:"required,min=6"`
}
