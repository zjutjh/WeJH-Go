package suppliesController

import (
	"math"
	"os"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/suppliesServices"
	"wejh-go/app/utils"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetBorrowRecordData struct {
	Campus uint8 `form:"campus" binding:"required"` // 校区 1:朝晖 2:屏峰 3: 莫干山
	Status int   `form:"status" binding:"required"` // 借用状态 1:未审核 2:被驳回 3:借用中 4:已归还
}

type GetBorrowRecordResponse struct {
	ID         int       `json:"id"`
	Count      uint      `json:"count"`
	Status     int       `json:"status"`
	Campus     uint8     `json:"campus"`
	Name       string    `json:"name"`
	Img        string    `json:"img"`
	Kind       string    `json:"kind"`
	Spec       string    `json:"spec"`
	ApplyTime  time.Time `json:"apply_time"`
	BorrowTime time.Time `json:"borrow_time"`
	ReturnTime time.Time `json:"return_time"`
}

// 获取借用记录
func GetBorrowRecord(c *gin.Context) {
	// 判断是否登录
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	// 获取参数
	var data GetBorrowRecordData
	err = c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断并获取借用记录
	var Record []models.BorrowRecord
	Record, err = suppliesServices.GetBorrowRecord(data.Campus, data.Status, user.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 输出借用记录
	borrowRecord := make([]GetBorrowRecordResponse, 0)
	for i := range Record {
		newRecord := GetBorrowRecordResponse{
			ID:         Record[i].ID,
			Count:      Record[i].Count,
			Status:     Record[i].Status,
			Campus:     Record[i].Campus,
			ApplyTime:  Record[i].ApplyTime,
			BorrowTime: Record[i].BorrowTime,
			ReturnTime: Record[i].ReturnTime,
		}
		borrowRecord = append(borrowRecord, newRecord)
		supplies, err := suppliesServices.GetALLSuppliesById(Record[i].SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		borrowRecord[i].Img = supplies.Img
		borrowRecord[i].Name = supplies.Name
		borrowRecord[i].Kind = supplies.Kind
		borrowRecord[i].Spec = supplies.Spec
	}
	utils.JsonSuccessResponse(c, borrowRecord)
}

type ApplyBorrowData struct {
	SuppliesID int  `json:"supplies_id" binding:"required"`
	Count      uint `json:"count" binding:"required"`
}

// 申请借用
func ApplyBorrow(c *gin.Context) {
	// 判断是否登录
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	// 判断个人信息是否存在
	personalinfo, err := suppliesServices.GetPersonalInfoByStudentID(user.StudentID)
	if err == gorm.ErrRecordNotFound {
		_ = c.AbortWithError(200, apiException.PersonalInfoNotFill)
		return
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 获取参数
	var data ApplyBorrowData
	err = c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断物资是否存在
	var supplies models.Supplies
	supplies, err = suppliesServices.GetALLSuppliesById(data.SuppliesID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if data.Count > supplies.Stock {
		_ = c.AbortWithError(200, apiException.StockNotEnough)
		return
	}
	// 判断是否已经申请过
	_, err = suppliesServices.GetBorrowRecordByApplyData(data.SuppliesID, user.StudentID)
	if err == nil {
		_ = c.AbortWithError(200, apiException.RecordAlreadyExisted)
		return
	} else if err != gorm.ErrRecordNotFound {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 申请借用
	err = suppliesServices.CreateBorrowRecord(models.BorrowRecord{
		Name:         personalinfo.Name,
		Gender:       personalinfo.Gender,
		StudentID:    personalinfo.StudentID,
		College:      personalinfo.College,
		Dormitory:    personalinfo.Dormitory,
		Contact:      personalinfo.Contact,
		Campus:       supplies.Campus,
		SuppliesID:   data.SuppliesID,
		Count:        data.Count,
		Organization: supplies.Organization,
		Status:       1,
		ApplyTime:    time.Now(),
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type DeleteRecordData struct {
	BorrowID int `form:"borrow_id" binding:"required"`
}

// 取消借用
func CancelBorrow(c *gin.Context) {
	// 判断是否登录
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	// 获取参数
	var data DeleteRecordData
	err = c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断借用记录是否存在
	var borrowRecord models.BorrowRecord
	borrowRecord, err = suppliesServices.GetBorrowRecordByBorrowID(data.BorrowID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否有权限
	if borrowRecord.StudentID != user.StudentID {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否待审核
	if borrowRecord.Status != 1 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 取消借用
	err = suppliesServices.DeleteRecord(data.BorrowID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type GetSuppliesRecordByAdminData struct {
	PageNum      int    `form:"page_num" binding:"required"`
	PageSize     int    `form:"page_size" binding:"required"`
	Campus       uint8  `form:"campus"`                           // 校区 1:朝晖 2:屏峰 3: 莫干山
	ID           int    `form:"id" `                              // 申请id
	StudentID    string `form:"student_id" `                      // 学号
	SuppliesName string `form:"supplies_name"`                    // 物资名称
	Spec         string `form:"spec"`                             // 规格
	Status       int    `form:"status" binding:"oneof=0 1 2 3 4"` // 状态 1:未审核 2:被驳回 3:借用中 4:已归还
	Choice       int    `form:"choice" binding:"oneof=0 1 2"`     // 1:审批 2:归还清点
}
type GetSuppliesRecordByAdminResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	StudentID    string    `json:"student_id"`
	Gender       string    `json:"gender"`
	College      string    `json:"college"`
	Dormitory    string    `json:"dormitory"`
	Contact      string    `json:"contact"`
	SuppliesID   int       `json:"supplies_id"`
	SuppliesName string    `json:"supplies_name"`
	Kind         string    `json:"kind"`
	Spec         string    `json:"spec"`
	Count        uint      `json:"count"`
	ApplyTime    time.Time `json:"apply_time"`
	BorrowTime   time.Time `json:"borrow_time"`
	ReturnTime   time.Time `json:"return_time"`
	Status       int       `json:"status"`
}

// 管理员获取获取学生申请/归还记录
func GetSuppliesRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data GetSuppliesRecordByAdminData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断choice和status是否合法
	validStatus := map[int][]int{
		1: {1, 2, 0},
		2: {3, 4, 0},
		0: {0, 1, 2, 3, 4},
	}
	valid := false
	for _, status := range validStatus[data.Choice] {
		if data.Status == status {
			valid = true
			break
		}
	}
	if !valid {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 获取记录和总页数
	var records []models.BorrowRecord
	var totalPageNum *int64
	records, totalPageNum, err = suppliesServices.GetRecordByAdmin(data.PageNum, data.PageSize, data.Status, data.Choice, data.ID, data.Campus, data.StudentID, data.SuppliesName, data.Spec)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	//输出记录
	response := make([]GetSuppliesRecordByAdminResponse, 0)
	for i := range records {
		newRecord := GetSuppliesRecordByAdminResponse{
			Gender:     records[i].Gender,
			StudentID:  records[i].StudentID,
			Status:     records[i].Status,
			Count:      records[i].Count,
			SuppliesID: records[i].SuppliesID,
			College:    records[i].College,
			Dormitory:  records[i].Dormitory,
			Contact:    records[i].Contact,
			ApplyTime:  records[i].ApplyTime,
			BorrowTime: records[i].BorrowTime,
			ReturnTime: records[i].ReturnTime,
		}
		response = append(response, newRecord)
		supplies, err := suppliesServices.GetALLSuppliesById(records[i].SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		record, err := suppliesServices.GetBorrowRecordByOthers(records[i].StudentID, records[i].College, records[i].SuppliesID, records[i].Count, records[i].Campus, records[i].ApplyTime)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		response[i].ID = record.ID
		response[i].Name = record.Name
		response[i].SuppliesName = supplies.Name
		response[i].Kind = supplies.Kind
		response[i].Spec = supplies.Spec

	}
	utils.JsonSuccessResponse(c, gin.H{
		"total_page_num": math.Ceil(float64(*totalPageNum) / float64(data.PageSize)),
		"data":           response,
	})
}

type CheckRecordData struct {
	SuppliesCheck int `json:"supplies_check" binding:"required,oneof=1 2"` // 1:通过 2:驳回
	ID            int `json:"id" binding:"required"`                       // 申请id
}

// 管理员审批
func CheckRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data CheckRecordData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否存在
	record, err := suppliesServices.GetBorrowRecordByBorrowID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否已经待审核
	if record.Status == 2 {
		_ = c.AbortWithError(200, apiException.RecordRejected)
		return
	} else if record.Status != 1 && record.Status != 2 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 审批
	if data.SuppliesCheck == 1 {
		//查询物资是否充足
		var supplies models.Supplies
		supplies, err = suppliesServices.GetALLSuppliesById(record.SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		if supplies.Stock < record.Count {
			_ = c.AbortWithError(200, apiException.StockNotEnough)
			return
		}
		err = suppliesServices.PassBorrow(data.ID, record.SuppliesID, record.Count)
	} else if data.SuppliesCheck == 2 {
		err = suppliesServices.RejectBorrow(data.ID)
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// 管理员取消驳回
type CancelRejectlDate struct {
	ID int `json:"id" binding:"required"`
}

func CancelRejectRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data CancelRejectlDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否存在
	record, err := suppliesServices.GetBorrowRecordByBorrowID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否已经驳回
	if record.Status != 2 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 取消驳回
	err = suppliesServices.CancelRejectBorrow(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type ReturnRecordData struct {
	ID             int `json:"id" binding:"required"`
	SuppliesReturn int `json:"supplies_return" binding:"required,oneof=1 2"` // 1:确认归还 2:取消借出
}

// 管理员归还清点
func ReturnRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data ReturnRecordData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否存在
	record, err := suppliesServices.GetBorrowRecordByBorrowID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否已经审批
	if record.Status != 3 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 归还清点
	if data.SuppliesReturn == 1 {
		err = suppliesServices.ReturnBorrow(data.ID, record.SuppliesID, record.Count)
	} else if data.SuppliesReturn == 2 {
		err = suppliesServices.CancelBorrow(data.ID, record.SuppliesID, record.Count)
	} else {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type CancelReturnDate struct {
	ID int `json:"id" binding:"required"`
}

// 管理员取消确认归还
func CancelReturnRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data CancelReturnDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否存在
	record, err := suppliesServices.GetBorrowRecordByBorrowID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断是否已经物资已经确认归还
	if record.Status != 4 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 取消确认归还
	err = suppliesServices.CancelReturnBorrow(data.ID, record.SuppliesID, record.Count)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type UpdateRecordData struct {
	ID        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	College   string `json:"college" binding:"required"`
	Dormitory string `json:"dormitory" binding:"required"`
	Contact   string `json:"contact" binding:"required"`
	SupplisID int    `json:"supplies_id" binding:"required"`
	Count     uint   `json:"count" binding:"required"`
}

// 管理员修改记录
func UpdateRecordByAdmin(c *gin.Context) {
	// 获取参数
	var data UpdateRecordData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	// 判断鉴权
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin && user.Type != models.StudentAffairsCenter {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否存在
	record, err := suppliesServices.GetBorrowRecordByBorrowID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断记录是否待审核
	if record.Status != 1 {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 判断物资是否重复借用
	if record.SuppliesID != data.SupplisID {
		_, err = suppliesServices.GetBorrowRecordByApplyData(data.SupplisID, record.StudentID)
		if err == nil {
			_ = c.AbortWithError(200, apiException.RecordAlreadyExisted)
			return
		} else if err != gorm.ErrRecordNotFound {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	//查询物资是否充足
	var supplies models.Supplies
	supplies, err = suppliesServices.GetALLSuppliesById(data.SupplisID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if supplies.Stock < data.Count {
		_ = c.AbortWithError(200, apiException.StockNotEnough)
		return
	}
	// 修改记录
	err = suppliesServices.UpdateRecord(data.ID, data.Name, data.Gender, data.College, data.Dormitory, data.Contact, data.SupplisID, data.Count)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type SuppliesRecordForm struct {
	StudentID    string `json:"student_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	College      string `json:"college" binding:"required"`
	Dormitory    string `json:"dormitory" binding:"required"`
	Contact      string `json:"contact" binding:"required"`
	Kind         string `json:"kind" binding:"required"`
	SuppliesName string `json:"supplies_name" binding:"required"`
	Spec         string `json:"spec" binding:"required"`
	Campus       uint8  `json:"campus" binding:"required"` //校区 1:朝晖 2:屏峰 3: 莫干山
	Count        uint   `json:"count" binding:"required"`
}

// 导入
func InsertSuppliesRecord(c *gin.Context) {
	var postForm SuppliesRecordForm
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

	_, err = suppliesServices.GetPersonalInfoByStudentID(postForm.StudentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = suppliesServices.SavePersonalInfo(models.PersonalInfo{
				Name:      postForm.Name,
				Gender:    postForm.Gender,
				StudentID: postForm.StudentID,
				College:   postForm.College,
				Dormitory: postForm.Dormitory,
				Contact:   postForm.Contact,
			})
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		} else {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	var flag bool
	if postForm.Kind != "正装" {
		flag, err = suppliesServices.CheckSupplies(postForm.SuppliesName, postForm.Kind, postForm.Spec, *publisher, postForm.Campus)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		if flag {
			err = suppliesServices.CreateSupplies(models.Supplies{
				Name:         postForm.SuppliesName,
				Kind:         postForm.Kind,
				Spec:         postForm.Spec,
				Campus:       postForm.Campus,
				Organization: *publisher,
			})
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
	} else {
		checkExist, err := suppliesServices.CheckSupplies(postForm.SuppliesName, postForm.Kind, postForm.Spec, *publisher, postForm.Campus)
		if checkExist == true && err == nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		} else if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		flag, err = suppliesServices.CheckSuppliesStock(postForm.SuppliesName, postForm.Kind, postForm.Spec, *publisher, postForm.Campus, postForm.Count)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		if flag {
			suppliesId, err := suppliesServices.GetSuppliesID(postForm.Campus, *publisher, postForm.Kind, postForm.SuppliesName, postForm.Spec)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
			err = suppliesServices.PassSuppliesRecord(suppliesId, postForm.Count)
			if err != nil {
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		} else if !flag {
			_ = c.AbortWithError(200, apiException.StockNotEnough)
			return
		}
	}

	suppliesId, err := suppliesServices.GetSuppliesID(postForm.Campus, *publisher, postForm.Kind, postForm.SuppliesName, postForm.Spec)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = suppliesServices.CreateBorrowRecord(models.BorrowRecord{
		Name:         postForm.Name,
		Gender:       postForm.Gender,
		StudentID:    postForm.StudentID,
		College:      postForm.College,
		Dormitory:    postForm.Dormitory,
		Contact:      postForm.Contact,
		Campus:       postForm.Campus,
		SuppliesID:   suppliesId,
		Count:        postForm.Count,
		Organization: *publisher,
		Status:       3,
		ApplyTime:    time.Now(),
		BorrowTime:   time.Now(),
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type GetExcelData struct {
	Campus uint8 `form:"campus" binding:"required"`
}

// 导出为excel文件
func ExportSuppliesRecord(c *gin.Context) {
	var data GetExcelData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := getIdentity(c)
	var ExcelData []models.BorrowRecord
	if *publisher == "学生事务大厅" {
		ExcelData, err = suppliesServices.GetExcelData(*publisher, data.Campus)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else if *publisher == "Admin" {
		ExcelData, err = suppliesServices.GetALLExcelData(data.Campus)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}

	//将数字转换为文字
	campusMap := map[uint8]string{
		1: "朝晖",
		2: "屏峰",
		3: "莫干山",
	}
	statusMap := map[int]string{
		3: "借用中",
		4: "已归还",
	}

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//设置列宽
	if err := streamWriter.SetColWidth(1, 20, 15); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//单独将学院这列加宽
	if err := streamWriter.SetColWidth(5, 5, 25); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//将时间三列加宽
	if err := streamWriter.SetColWidth(14, 16, 25); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//设置样式
	styleID, err := f.NewStyle(&excelize.Style{
		//居中对齐
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		//颜色
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#DFEBF6"}, Pattern: 1},
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{Value: "学生借用记录表", StyleID: styleID},
	}, excelize.RowOpts{Height: 30, Hidden: false}); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//合并单元格
	if err := streamWriter.MergeCell("A1", "P1"); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//设置列名
	header := []interface{}{}
	for _, cell := range []string{
		"申请ID", "姓名", "性别", "学号", "学院", "寝室号", "联系方式", "校区", "物资名称", "种类", "规格", "借用数量", "状态", "申请时间", "借用时间", "归还时间",
	} {
		header = append(header, cell)
	}
	if err := streamWriter.SetRow("A2", header); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	//批量导入数据
	for rowID, record := range ExcelData {
		campus, _ := campusMap[record.Campus]
		status, _ := statusMap[record.Status]
		result, err := suppliesServices.GetALLSuppliesById(record.SuppliesID)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		applyTime := ""
		if !record.ApplyTime.IsZero() {
			applyTime = record.ApplyTime.Format("2006-01-02 15:04:05")
		}
		borrowTime := ""
		if !record.BorrowTime.IsZero() {
			borrowTime = record.BorrowTime.Format("2006-01-02 15:04:05")
		}
		returnTime := ""
		if !record.ReturnTime.IsZero() {
			returnTime = record.ReturnTime.Format("2006-01-02 15:04:05")
		}
		row := []interface{}{
			record.ID, record.Name, record.Gender, record.StudentID, record.College, record.Dormitory,
			record.Contact, campus, result.Name, result.Kind, result.Spec, record.Count,
			status, applyTime, borrowTime, returnTime,
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID+3)
		if err := streamWriter.SetRow(cell, row); err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	//关闭
	if err := streamWriter.Flush(); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	// 保存Excel文件
	fileName := uuid.NewString() + ".xlsx"
	filePath := "./files/" + fileName
	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		err := os.Mkdir("./files", 0755)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	if err := f.SaveAs(filePath); err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, config.GetFileUrlKey()+fileName)
}
