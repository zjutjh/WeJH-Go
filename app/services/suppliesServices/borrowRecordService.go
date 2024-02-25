package suppliesServices

import (
	"time"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/utils"
	"wejh-go/config/database"

	"gorm.io/gorm"
)

func GetBorrowRecord(campus uint8, status int, studentID string) ([]models.BorrowRecord, error) {
	var borrowRecord []models.BorrowRecord
	var result *gorm.DB
	if status == 1 {
		result = database.DB.Where(models.BorrowRecord{
			Campus:    campus,
			Status:    1,
			StudentID: studentID,
		}).Or(models.BorrowRecord{
			Campus:    campus,
			Status:    2,
			StudentID: studentID,
		}).Order("apply_time desc").Find(&borrowRecord)
	} else {
		result = database.DB.Where(models.BorrowRecord{
			Campus:    campus,
			Status:    status,
			StudentID: studentID,
		}).Order("borrow_time").Order("apply_time desc").Find(&borrowRecord)
	}
	return borrowRecord, result.Error
}

func GetBorrowRecordByApplyData(suppliesID int, studentID string) (models.BorrowRecord, error) {
	var borrowRecord models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		SuppliesID: suppliesID,
		StudentID:  studentID,
		Status:     1,
	}).Or(models.BorrowRecord{
		SuppliesID: suppliesID,
		StudentID:  studentID,
		Status:     3,
	}).First(&borrowRecord)
	if result.Error != nil {
		return borrowRecord, result.Error
	}
	aesDecryptContact(&borrowRecord)
	return borrowRecord, nil
}

func CreateBorrowRecord(borrowRecord models.BorrowRecord) error {
	aesEncryptContact(&borrowRecord)
	result := database.DB.Create(&borrowRecord)
	return result.Error
}

func GetBorrowRecordByBorrowID(borrowID int) (models.BorrowRecord, error) {
	var borrowRecord models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		ID: borrowID,
	}).First(&borrowRecord)
	if result.Error != nil {
		return borrowRecord, result.Error
	}
	aesDecryptContact(&borrowRecord)
	return borrowRecord, nil
}

func DeleteRecord(RecordID int) error {
	result := database.DB.Delete(models.BorrowRecord{ID: RecordID})
	return result.Error
}

func GetRecordByAdmin(pageNum, pageSize, status, choice, id int, campus uint8, studentID, name, spec string) ([]models.BorrowRecord, *int64, error) {
	var record []models.BorrowRecord
	var num int64
	query := database.DB.Table("borrow_records").
		Select("borrow_records.*, supplies.*").
		Joins("JOIN supplies ON borrow_records.supplies_id = supplies.id").
		Where(models.BorrowRecord{
			ID:        id,
			Campus:    campus,
			StudentID: studentID,
		}).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("status").Debug()

	if name != "" {
		query = query.Where("supplies.name = ?", name)
	}
	if spec != "" {
		query = query.Where("supplies.spec = ?", spec)
	}
	if status != 0 {
		query = query.Where(models.BorrowRecord{Status: status})
	} else {
		switch choice {
		case 0:
			query = query.Where("status IN ?", []int{1, 2, 3, 4}).Order("apply_time desc")
		case 1:
			query = query.Where("status IN ?", []int{1, 2}).Order("apply_time desc")
		case 2:
			query = query.Where("status IN ?", []int{3, 4}).Order("CASE WHEN DATE_ADD(borrow_time, INTERVAL 7 DAY) > NOW() THEN 1 ELSE 0 END, borrow_time").Order("apply_time desc")
		}
	}

	result := query.Find(&record)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	result = query.Count(&num)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	for i := range record {
		aesDecryptContact(&record[i])
	}
	return record, &num, nil
}

func GetBorrowRecordByOthers(studentID, college string, supplies_id int, count uint, campus uint8, time time.Time) (models.BorrowRecord, error) {
	var borrowRecord models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		StudentID:  studentID,
		College:    college,
		SuppliesID: supplies_id,
		Campus:     campus,
		Count:      count,
		ApplyTime:  time,
	}).First(&borrowRecord)
	return borrowRecord, result.Error
}

func PassBorrow(id int, sid int, num uint) error {
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Updates(models.BorrowRecord{Status: 3, BorrowTime: time.Now()})
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Model(models.Supplies{}).Unscoped().Where(models.Supplies{ID: sid}).Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", num), "borrowed": gorm.Expr("borrowed + ?", num)})
	return result.Error
}

func RejectBorrow(id int) error {
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Update("status", 2)
	return result.Error
}

func CancelRejectBorrow(id int) error {
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Update("status", 1)
	return result.Error
}

func ReturnBorrow(id int, sid int, num uint) error {
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Updates(models.BorrowRecord{Status: 4, ReturnTime: time.Now()})
	if result.Error != nil {
		return result.Error
	}
	var  supplies models.Supplies
	result = database.DB.Where(models.Supplies{ID: sid}).Unscoped().First(&supplies)
	if result.Error != nil {
		return result.Error
	}
	if supplies.Kind == "正装" {
		result = database.DB.Model(models.Supplies{}).Where(models.Supplies{ID: sid}).Unscoped().Updates(map[string]interface{}{"stock": gorm.Expr("stock + ?", num), "borrowed": gorm.Expr("borrowed - ?", num)})
	}
	return result.Error
}

func CancelBorrow(id int, sid int, num uint) error {
	var supplies models.Supplies
	result := database.DB.Where(models.Supplies{ID: sid}).Unscoped().First(&supplies)
	if result.Error != nil {
		return result.Error
	}
	if supplies.Kind == "正装" {
		result = database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Updates(map[string]interface{}{"status": 1, "borrow_time": nil})
		if result.Error != nil {
			return result.Error
		}
		result = database.DB.Model(models.Supplies{}).Where(models.Supplies{ID: sid}).Unscoped().Updates(map[string]interface{}{"stock": gorm.Expr("stock + ?", num), "borrowed": gorm.Expr("borrowed - ?", num)})
	} else {
		result = database.DB.Delete(models.BorrowRecord{ID: id})
	}
	return result.Error
}

func CancelReturnBorrow(id int, sid int, num uint) error {
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Updates(map[string]interface{}{"status": 3, "return_time": nil})
	if result.Error != nil {
		return result.Error
	}
	var supplies models.Supplies
	result = database.DB.Where(models.Supplies{ID: sid}).Unscoped().First(&supplies)
	if result.Error != nil {
		return result.Error
	}
	if supplies.Kind == "正装" {
		result = database.DB.Model(models.Supplies{}).Where(models.Supplies{ID: sid}).Unscoped().Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", num), "borrowed": gorm.Expr("borrowed + ?", num)})
	}
	return result.Error
}

func UpdateRecord(id int, name, gender, college, dormitory, contact string, suppliesid int, count uint) error {
	contact = utils.AesEncrypt(contact, config.GetEncryptKey())
	result := database.DB.Model(models.BorrowRecord{}).Where(models.BorrowRecord{ID: id}).Updates(models.BorrowRecord{
		Name:       name,
		Gender:     gender,
		College:    college,
		Dormitory:  dormitory,
		Contact:    contact,
		SuppliesID: suppliesid,
		Count:      count,
	})
	return result.Error
}

func GetSuppliesID(campus uint8, organization string, kind string, name string, spec string) (int, error) {
	var supplies *models.Supplies
	result := database.DB.Order("id DESC").Where(models.Supplies{
		Kind:         kind,
		Name:         name,
		Spec:         spec,
		Campus:       campus,
		Organization: organization,
	}).First(&supplies)
	return supplies.ID, result.Error
}

func GetExcelData(organization string, campus uint8) ([]models.BorrowRecord, error) {
	var supplies []models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		Campus:       campus,
		Organization: organization,
	}).Where("status = ? OR status = ?", 3, 4).Find(&supplies)
	for i := range supplies {
		aesDecryptContact(&supplies[i])
	}
	return supplies, result.Error
}

func GetALLExcelData(campus uint8) ([]models.BorrowRecord, error) {
	var supplies []models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		Campus: campus,
	}).Where("status = ? OR status = ?", 3, 4).Find(&supplies)
	for i := range supplies {
		aesDecryptContact(&supplies[i])
	}
	return supplies, result.Error
}

func GetBorrowRecordBySuppliesID(suppliesID int) (models.BorrowRecord, error) {
	var borrowRecord models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		SuppliesID: suppliesID,
	}).First(&borrowRecord)
	if result.Error != nil {
		return borrowRecord, result.Error
	}
	aesDecryptContact(&borrowRecord)
	return borrowRecord, nil
}

func GetALLBorrowRecordBySuppliesID(suppliesID int) ([]models.BorrowRecord, error) {
	var borrowRecord []models.BorrowRecord
	result := database.DB.Where(models.BorrowRecord{
		SuppliesID: suppliesID,
	}).Find(&borrowRecord)
	return borrowRecord, result.Error
}

func DeleteBorrowRecord(borrowRecord models.BorrowRecord) error {
	result := database.DB.Model(models.BorrowRecord{}).Delete(&borrowRecord)
	return result.Error
}

func aesEncryptContact(record *models.BorrowRecord) {
	key := config.GetEncryptKey()
	record.Contact = utils.AesEncrypt(record.Contact, key)
}

func aesDecryptContact(record *models.BorrowRecord) {
	key := config.GetEncryptKey()
	record.Contact = utils.AesDecrypt(record.Contact, key)
}
