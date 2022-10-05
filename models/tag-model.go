package models

type Tag struct {
	BaseModel
	Name string `json:"name"`

	Jobs []Job `gorm:"many2many:job_tags;" json:"jobs"`
}

type TagJob struct {
	BaseModel
	Name string `json:"name"`
}

func (TagJob) TableName() string {
	return "tags"
}

type TagInput struct {
	Name string `json:"name"`
}

func (TagInput) TableName() string {
	return "tags"
}
