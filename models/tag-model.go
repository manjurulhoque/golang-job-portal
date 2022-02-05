package models

type Tag struct {
	BaseModel
	Name string `json:"name"`
}

type TagInput struct {
	Name string `json:"name"`
}
