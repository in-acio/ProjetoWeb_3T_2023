package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VoteInput struct {
	ItemId uint `json:"item_id"`
	Value  uint `json:"value"`
}