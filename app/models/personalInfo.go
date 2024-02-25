package models

type PersonalInfo struct {
	ID        int    `json:"id" gorm:"autoIncrement;unique"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	StudentID string `gorm:"primaryKey" json:"student_id"`
	College   string `json:"college"`   // 学院
	Dormitory string `json:"dormitory"` // 寝室号
	Contact   string `json:"contact"`   //联系方式
}
