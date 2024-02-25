package suppliesServices

import (
	"gorm.io/gorm"
	"os"
	"strings"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetSupplies(campus uint8) ([]models.Supplies, error) {
	var supplies []models.Supplies
	result := database.DB.Where(models.Supplies{
		Kind:   "正装",
		Campus: campus,
	}).Find(&supplies)
	return supplies, result.Error
}

func CreateSupplies(record models.Supplies) error {
	result := database.DB.Create(&record)
	return result.Error
}

func CheckSupplies(suppliesName, kind, spec, organization string, campus uint8) (bool, error) {
	var flag bool
	var supplies models.Supplies
	result := database.DB.Where(models.Supplies{
		Name:         suppliesName,
		Kind:         kind,
		Spec:         spec,
		Organization: organization,
		Campus:       campus,
	}).First(&supplies)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			flag = true
			return flag, nil
		} else {
			flag = false
			return flag, result.Error
		}
	} else {
		flag = false
		return flag, nil
	}
}

func CheckSuppliesToRemoveImg(suppliesId int, suppliesName, kind, organization string, campus uint8) (bool, error) {
	var flag bool
	var supplies models.Supplies
	result := database.DB.Where(models.Supplies{
		Name:         suppliesName,
		Kind:         kind,
		Organization: organization,
		Campus:       campus,
	}).Not("id = ?", suppliesId).First(&supplies)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			flag = true
			return flag, nil
		} else {
			flag = false
			return flag, result.Error
		}
	} else {
		flag = false
		return flag, nil
	}
}

func CheckSuppliesStock(suppliesName, kind, spec, organization string, campus uint8, num uint) (bool, error) {
	var flag bool
	var supplies models.Supplies
	result := database.DB.Where(models.Supplies{
		Name:         suppliesName,
		Kind:         kind,
		Spec:         spec,
		Organization: organization,
		Campus:       campus,
	}).First(&supplies)
	if result.Error != nil {
		return false, result.Error
	}
	if num <= supplies.Stock {
		flag = true
		return flag, nil
	} else {
		flag = false
		return flag, nil
	}
}

func PassSuppliesRecord(id int, num uint) error {
	result := database.DB.Model(models.Supplies{}).Where(models.Supplies{ID: id}).Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", num), "borrowed": gorm.Expr("borrowed + ?", num)})
	return result.Error
}

func UpdateStockByInsertSpec(id int, num uint) error {
	result := database.DB.Model(models.Supplies{}).Where(models.Supplies{ID: id}).Updates(map[string]interface{}{"stock": gorm.Expr("stock + ?", num)})
	return result.Error
}

func GetSuppliesById(id int) (models.Supplies, error) {
	var record models.Supplies
	result := database.DB.Where(models.Supplies{
		ID: id,
	}).First(&record)
	if result.Error != nil {
		return models.Supplies{}, result.Error
	}
	return record, nil
}

func GetALLSuppliesById(id int) (models.Supplies, error) {
	var record models.Supplies
	result := database.DB.Unscoped().Where(models.Supplies{
		ID: id,
	}).First(&record)
	return record, result.Error
}

func RemoveImg(record models.Supplies, img string) {
	if record.Img != "" && record.Img != img {
		_ = os.Remove("./img/" + strings.TrimPrefix(record.Img, config.GetWebpUrlKey()))
	}
}

func UpdateSupplies(id int, record models.Supplies) error {
	result := database.DB.Model(models.Supplies{}).Select("*").
		Omit("id", "kind", "organization", "borrowed").
		Where(&models.Supplies{ID: id}).Updates(&record)
	return result.Error
}

func DeleteSupplies(record models.Supplies) error {
	result := database.DB.Delete(&record)
	return result.Error
}

func CompletedDeleteSupplies(record models.Supplies) error {
	result := database.DB.Unscoped().Delete(&record)
	return result.Error
}

func GetSuppliesByPublisher(campus uint8, organization string) ([]models.Supplies, error) {
	var supplies []models.Supplies
	result := database.DB.Where(models.Supplies{
		Campus:       campus,
		Kind:         "正装",
		Organization: organization,
	}).Find(&supplies)
	return supplies, result.Error
}

func GetSuppliesByAdmin(campus uint8) ([]models.Supplies, error) {
	var supplies []models.Supplies
	result := database.DB.Unscoped().Where(models.Supplies{
		Campus: campus,
		Kind:   "正装",
	}).Find(&supplies)
	return supplies, result.Error
}
