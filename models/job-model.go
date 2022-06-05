package models

type Job struct {
	BaseModel
	Title       string `json:"title" validate:"required"`
	Description string `gorm:"type:text" json:"description" validate:"required"`
	UserId      uint   `json:"user_id"`
	Salary      int    `json:"salary" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Filled      bool   `gorm:"default:false" json:"filled"`

	User User `gorm:"foreignkey:user_id" json:"user"`
	//Tags []Tag `gorm:"many2many:job_tags;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:tag_id;jointable_foreignkey:job_id;" json:"tags"`
	Tags []Tag `gorm:"many2many:job_tags;" json:"tags"`
}

type JobInput struct {
	BaseModel
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Salary      int    `json:"salary" validate:"required"`
	Location    string `json:"location" validate:"required"`

	Tags []int `json:"tags"`
}
