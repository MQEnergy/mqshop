package user

type IndexReq struct {
	Page  int `form:"page" json:"page" validate:"required"`
	Limit int `form:"limit" json:"limit" validate:"required"`
}

type LoginReq struct {
	Account  string `form:"account" json:"account" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
