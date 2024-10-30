package models

type Applicant struct {
	BaseModel
	Comment string `gorm:"type:text" json:"comment"`
	Status  int    `gorm:"default:1" json:"status"`
	UserId  uint   `json:"user_id"`
	JobId   uint   `json:"job_id"`

	Job Job `json:"-"`
}

type ApplicantInput struct {
	BaseModel
	Comment string `json:"comment"`
	Status  int    `json:"status"`
	JobId   uint   `json:"job_id"`
}
