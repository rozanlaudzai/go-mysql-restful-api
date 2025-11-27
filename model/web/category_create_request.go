package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=200"`
}
