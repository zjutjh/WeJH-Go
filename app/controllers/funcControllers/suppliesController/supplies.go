package suppliesController

import (
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/services/suppliesServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type GetSuppliesDate struct {
	Campus uint8 `form:"campus" binding:"required"` // 校区 1:朝晖 2:屏峰 3: 莫干山
}

// 获取物资列表
func GetSuppliesList(c *gin.Context) {
	// 获取参数,判断校区
	var data GetSuppliesDate
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	//读取各校区的物资
	var supplies []models.Supplies
	supplies, err = suppliesServices.GetSupplies(data.Campus)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 输出物资列表
	groupedSupplies := make(map[string][]models.Supplies)
	for _, supply := range supplies {
		groupedSupplies[supply.Name] = append(groupedSupplies[supply.Name], supply)
	}
	var response []map[string]interface{}
	for name, supplies := range groupedSupplies {
		specs := make([]map[string]interface{}, 0)
		img := ""
		for i, supply := range supplies {
			spec := make(map[string]interface{})
			spec["id"] = supply.ID
			spec["spec"] = supply.Spec
			spec["stock"] = supply.Stock
			if i == 0 {
				img = supply.Img
			}
			specs = append(specs, spec)
		}
		item := make(map[string]interface{})
		item["name"] = name
		item["img"] = img
		item["specs"] = specs
		response = append(response, item)
	}

	utils.JsonSuccessResponse(c, response)
}

type SpecInfo struct {
	ID    int    `json:"id"`
	Spec  string `json:"spec" binding:"required"`
	Stock uint   `json:"stock" binding:"required"`
}

type SuppliesForm struct {
	Name   string      `json:"name" binding:"required"`   // 物资名称
	Campus uint8       `json:"campus" binding:"required"` // 校区 1:朝晖 2:屏峰 3: 莫干山
	Img    interface{} `json:"img" binding:"required"`
	Specs  []SpecInfo  `json:"specs" binding:"required"`
}

// 发布正装信息
func InsertSupplies(c *gin.Context) {
	var postForm SuppliesForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := getIdentity(c)
	if *publisher != "学生事务大厅" && *publisher != "Admin" {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	for _, spec := range postForm.Specs {
		record := models.Supplies{
			Name:         postForm.Name,
			Kind:         "正装",
			Campus:       postForm.Campus,
			Organization: *publisher,
			Spec:         spec.Spec,
			Stock:        spec.Stock,
		}

		if postForm.Img != nil {
			record.Img = postForm.Img.(string)
		}

		err := suppliesServices.CreateSupplies(record)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}

// 修改正装信息
func UpdateSupplies(c *gin.Context) {
	var postForm SuppliesForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := getIdentity(c)
	if *publisher != "学生事务大厅" && *publisher != "Admin" {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	for _, spec := range postForm.Specs {
		if spec.ID == 0 {
			checkExist, err := suppliesServices.CheckSupplies(postForm.Name, "正装", spec.Spec, *publisher, postForm.Campus)
			if checkExist == true && err == nil {
				record := models.Supplies{
					Name:         postForm.Name,
					Kind:         "正装",
					Campus:       postForm.Campus,
					Organization: *publisher,
					Spec:         spec.Spec,
					Stock:        spec.Stock,
					Img:          postForm.Img.(string),
				}
				err = suppliesServices.CreateSupplies(record)
				if err != nil {
					_ = c.AbortWithError(200, apiException.ServerError)
					return
				}
			} else if checkExist == false && err == nil {
				suppliesId, err := suppliesServices.GetSuppliesID(postForm.Campus, *publisher, "正装", postForm.Name, spec.Spec)
				if err != nil {
					_ = c.AbortWithError(200, apiException.ServerError)
					return
				}
				err = suppliesServices.UpdateStockByInsertSpec(suppliesId, spec.Stock)
				if err != nil {
					_ = c.AbortWithError(200, apiException.ServerError)
					return
				}
			} else {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		} else {
			record, err := suppliesServices.GetSuppliesById(spec.ID)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}

			var img string
			if postForm.Img != nil {
				img = postForm.Img.(string)
			}

			// 移除更新后不需要的图片
			suppliesServices.RemoveImg(record, img)

			err = suppliesServices.UpdateSupplies(spec.ID, models.Supplies{
				Name:   postForm.Name,
				Campus: postForm.Campus,
				Spec:   spec.Spec,
				Stock:  spec.Stock,
				Img:    img,
			})
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
	}

	utils.JsonSuccessResponse(c, nil)
}

// 删除正装信息
func DeleteSupplies(c *gin.Context) {
	SuppliesIDRaw := c.Query("id")
	SuppliesID, err := strconv.Atoi(SuppliesIDRaw)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := getIdentity(c)
	if *publisher != "学生事务大厅" && *publisher != "Admin" {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	supplies, err := suppliesServices.GetSuppliesById(SuppliesID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	checkSuppliesExist, err := suppliesServices.CheckSuppliesToRemoveImg(SuppliesID, supplies.Name, supplies.Kind, *publisher, supplies.Campus)
	if checkSuppliesExist == true && err == nil {
		borrowRecords, err := suppliesServices.GetALLBorrowRecordBySuppliesID(SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		deleteImg := false
		for _, record := range borrowRecords {
			if record.Status == 3 || record.Status == 4 {
				deleteImg = true
				break
			}
		}
		if !deleteImg {
			_ = os.Remove("./img/" + strings.TrimPrefix(supplies.Img, config.GetWebpUrlKey()))
		}
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	_, err = suppliesServices.GetBorrowRecordBySuppliesID(SuppliesID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		} else {
			err = suppliesServices.CompletedDeleteSupplies(supplies)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
	} else {
		allBorrowRecord, err := suppliesServices.GetALLBorrowRecordBySuppliesID(SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		flag := false
		for _, record := range allBorrowRecord {
			if record.Status == 3 || record.Status == 4 {
				err = suppliesServices.DeleteSupplies(supplies)
				if err != nil {
					_ = c.AbortWithError(200, apiException.ServerError)
					return
				}
				flag = true
				break
			}
		}
		if !flag {
			err = suppliesServices.CompletedDeleteSupplies(supplies)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
		for _, record := range allBorrowRecord {
			if record.Status == 1 || record.Status == 2 {
				err = suppliesServices.DeleteBorrowRecord(record)
				if err != nil {
					_ = c.AbortWithError(200, apiException.ServerError)
					return
				}
			}
		}
	}

	utils.JsonSuccessResponse(c, nil)
}

func GetSuppliesByAdmin(c *gin.Context) {
	var data GetSuppliesDate
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := getIdentity(c)
	var supplies []models.Supplies
	if *publisher == "学生事务大厅" {
		supplies, err = suppliesServices.GetSuppliesByPublisher(data.Campus, *publisher)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else if *publisher == "Admin" {
		supplies, err = suppliesServices.GetSuppliesByAdmin(data.Campus)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	groupedSupplies := make(map[string][]models.Supplies)
	for _, supply := range supplies {
		groupedSupplies[supply.Name] = append(groupedSupplies[supply.Name], supply)
	}
	var response []map[string]interface{}
	for name, supplies := range groupedSupplies {
		specs := make([]map[string]interface{}, 0)
		img := ""
		var campus uint8
		for i, supply := range supplies {
			spec := make(map[string]interface{})
			spec["id"] = supply.ID
			spec["spec"] = supply.Spec
			spec["stock"] = supply.Stock
			spec["borrowed"] = supply.Borrowed
			if i == 0 {
				img = supply.Img
				campus = supply.Campus
			}
			specs = append(specs, spec)
		}
		item := make(map[string]interface{})
		item["name"] = name
		item["img"] = img
		item["campus"] = campus
		item["specs"] = specs
		response = append(response, item)
	}

	utils.JsonSuccessResponse(c, response)
}
