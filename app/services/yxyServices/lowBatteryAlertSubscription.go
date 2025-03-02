package yxyServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"

	"gorm.io/gorm"
)

func SubscribeLowBatteryAlert(Subscription models.LowBatteryAlertSubscription) error {
	if err := database.DB.Where("user_id = ? AND campus = ?", Subscription.UserID, Subscription.Campus).
		First(&models.LowBatteryAlertSubscription{}).Error; err == gorm.ErrRecordNotFound {
		Subscription.Count = 1
		if e := database.DB.Create(&Subscription).Error; e != nil {
			return e
		}
		return nil
	} else if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"threshold": Subscription.Threshold,
		"count":     gorm.Expr("count + 1"),
	}
	err := database.DB.Model(&models.LowBatteryAlertSubscription{}).
		Where("user_id = ? AND campus = ?", Subscription.UserID, Subscription.Campus).
		Updates(updates).Error
	return err
}

func GetOrCreateLowBatteryAlertSubscription(userID int, campus string) (*models.LowBatteryAlertSubscription, error) {
	var subscription models.LowBatteryAlertSubscription
	if err := database.DB.Where(models.LowBatteryAlertSubscription{
		UserID: userID,
		Campus: campus,
	}).Attrs(models.LowBatteryAlertSubscription{
		Threshold: 20,
		Count:     0,
	}).FirstOrCreate(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}
