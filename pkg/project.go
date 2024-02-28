package pkg

type Project struct {
	ID           uint   `gorm:primary_key;index:idx_name,unique`
	ProjectTitle string `json:"project_title"`
	IsActive     bool   `json:is_active`
}
