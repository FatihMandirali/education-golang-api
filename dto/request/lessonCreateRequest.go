package request

type LessonCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
