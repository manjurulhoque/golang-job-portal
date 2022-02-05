package models

type Tag struct {
	BaseModel
	Name string `json:"name"`

	Jobs []Job `gorm:"many2many:job_tags;" json:"jobs"`
}

type TagInput struct {
	Name string `json:"name"`
}
