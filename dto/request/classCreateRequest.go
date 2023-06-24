package request

type ClassCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branchId"`
	IsActive bool   `json:"isActive"`
}
