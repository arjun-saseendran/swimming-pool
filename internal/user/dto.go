package user

type InputUser struct {
	Name   string `json:"name" binding:"required,min=2,max=50"`
	Mobile string `json:"mobile" binding:"required,min=10,max=15"`
}

type UpdateInputUser struct {
	Name   string `json:"name" binding:"omitempty,min=2,max=50"`
	Mobile string `json:"mobile" binding:"omitempty,min=10,max=15"`
}

func NewInputUser() *InputUser {
	return &InputUser{}
}

func NewUpdateUser() *UpdateInputUser {
	return &UpdateInputUser{}
}
