package models

// Member model
type Member struct {
	ID           string `gorm:"column:mem_id;primaryKey" json:"id"`
	SubClassID   string `gorm:"column:sub_class_id" json:"subclassId"`
	SubClassType string `gorm:"column:sub_class_type" json:"subclassType"`
}

func (member *Member) TableName() string {
	return "member"
}
