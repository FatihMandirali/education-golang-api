package request

type ClassUpdateRequest struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	BranchId int    `json:"branchId" binding:"required"`
}
