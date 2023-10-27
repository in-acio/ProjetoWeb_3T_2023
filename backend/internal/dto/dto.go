package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name        string `json:"name"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VoteInput struct {
	ItemId uint `json:"item_id"`
	Value  uint `json:"value"`
}