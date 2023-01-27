package request

type BranchUpdateRequest struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
}
