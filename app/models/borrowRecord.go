package models

import "time"

type BorrowRecord struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Gender       string    `json:"gender"`
	StudentID    string    `json:"student_id"`
	College      string    `json:"college"`   // 学院
	Dormitory    string    `json:"dormitory"` // 寝室号
	Contact      string    `json:"contact"`   //联系方式
	Campus       uint8     `json:"campus"`    // 校区 1:朝晖 2:屏峰 3: 莫干山
	SuppliesID   int       `json:"supplies_id"`
	Count        uint      `json:"count"`
	Organization string    `json:"organization"`
	ApplyTime    time.Time `json:"apply_time" gorm:"type:timestamp"`
	BorrowTime   time.Time `json:"borrow_time" gorm:"type:timestamp;default:null;"`
	ReturnTime   time.Time `json:"return_time" gorm:"type:timestamp;default:null;"`
	Status       int       `json:"status"` // 借用状态 1:未审核 2:被驳回 3:借用中 4:已归还
}
