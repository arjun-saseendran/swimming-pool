package dto

type UserCreateInput struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type UserUpdateInput struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func NewUserCreateInput() *UserCreateInput {
	return &UserCreateInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}
