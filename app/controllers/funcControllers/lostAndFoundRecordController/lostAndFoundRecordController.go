package lostAndFoundRecordController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image/jpeg"
	"image/png"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/services/lostAndFoundRecordServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type LostAndFoundForm struct {
	ID      int         `json:"id"`
	Type    bool        `json:"type"`
	Campus  string      `json:"campus"`
	Kind    string      `json:"kind"`
	Img1    interface{} `json:"img1"`
	Img2    interface{} `json:"img2"`
	Img3    interface{} `json:"img3"`
	Content string      `json:"content"`
}

func GetRecords(c *gin.Context) {

	campus, _ := url.QueryUnescape(c.Query("campus"))
	kind, _ := url.QueryUnescape(c.Query("kind"))
	pageNumRaw := c.Query("page_num")
	pageSizeRaw := c.Query("page_size")
	pageNum, err := strconv.Atoi(pageNumRaw)
	pageSize, err := strconv.Atoi(pageSizeRaw)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	var lostAndFoundRecords []models.LostAndFoundRecord
	if kind == "全部" {
		lostAndFoundRecords, err = lostAndFoundRecordServices.GetAllKindRecord(campus, pageNum, pageSize)
	} else {
		lostAndFoundRecords, err = lostAndFoundRecordServices.GetRecord(campus, kind, pageNum, pageSize)
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	var totalPageNum *int64
	if kind == "全部" {
		totalPageNum, err = lostAndFoundRecordServices.GetRecordAllKindTotalPageNum(campus)
	} else {
		totalPageNum, err = lostAndFoundRecordServices.GetRecordTotalPageNum(campus, kind)
	}
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"data":           lostAndFoundRecords,
		"total_page_num": math.Ceil(float64(*totalPageNum) / float64(pageSize)),
	})
}

func GetRecordsByAdmin(c *gin.Context) {
	pageNumRaw := c.Query("page_num")
	pageSizeRaw := c.Query("page_size")
	pageNum, err := strconv.Atoi(pageNumRaw)
	pageSize, err := strconv.Atoi(pageSizeRaw)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	var lostAndFoundRecords []models.LostAndFoundRecord
	var totalPageNum *int64
	publisher := getPublisher(c)
	if *publisher == "Admin" {
		lostAndFoundRecords, err = lostAndFoundRecordServices.GetRecordBySuperAdmin(pageNum, pageSize)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		totalPageNum, err = lostAndFoundRecordServices.GetRecordTotalPageNumBySuperAdmin()
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else {
		lostAndFoundRecords, err = lostAndFoundRecordServices.GetRecordByAdmin(*publisher, pageNum, pageSize)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
		totalPageNum, err = lostAndFoundRecordServices.GetRecordTotalPageNumByAdmin(*publisher)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, gin.H{
		"data":           lostAndFoundRecords,
		"total_page_num": math.Ceil(float64(*totalPageNum) / float64(pageSize)),
	})
}

func GetKindList(c *gin.Context) {
	kinds, err := lostAndFoundRecordServices.GetKindList()
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, kinds)
}

func InsertRecord(c *gin.Context) {
	var postForm LostAndFoundForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	publisher := getPublisher(c)
	record := models.LostAndFoundRecord{
		Type:        postForm.Type,
		Campus:      postForm.Campus,
		Kind:        postForm.Kind,
		PublishTime: time.Now(),
		IsProcessed: false,
		Publisher:   *publisher,
		Content:     postForm.Content,
	}
	if postForm.Img1 != nil {
		record.Img1 = postForm.Img1.(string)
	}
	if postForm.Img2 != nil {
		record.Img2 = postForm.Img2.(string)
	}
	if postForm.Img3 != nil {
		record.Img3 = postForm.Img3.(string)
	}
	err = lostAndFoundRecordServices.CreateRecord(record)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func UpdateRecord(c *gin.Context) {
	var postForm LostAndFoundForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	publisher := getPublisher(c)
	record, err := lostAndFoundRecordServices.GetRecordById(postForm.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if *publisher != record.Publisher {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}
	err = lostAndFoundRecordServices.UpdateRecord(postForm.ID, models.LostAndFoundRecord{
		Type:        postForm.Type,
		Campus:      postForm.Campus,
		Kind:        postForm.Kind,
		Content:     postForm.Content,
		Publisher:   *publisher,
		IsProcessed: false,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func getPublisher(c *gin.Context) *string {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return nil
	}
	var publisher string
	if user.Username == "zhforu" {
		publisher = "For You 朝晖校区"
	} else if user.Username == "pfforu" {
		publisher = "For You 屏峰校区"
	} else if user.Username == "mgsforu" {
		publisher = "For You 莫干山校区"
	} else if user.Username == "zhstuac" {
		publisher = "朝晖学生事务大厅"
	} else if user.Username == "pfstuac" {
		publisher = "屏峰学生事务大厅"
	} else if user.Username == "mgsstuac" {
		publisher = "莫干山学生事务大厅"
	} else if user.Type == models.Admin {
		publisher = "Admin"
	}
	return &publisher
}

func UploadImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	img := form.File["img"][0]
	imgName := img.Filename
	_ = c.SaveUploadedFile(img, "./tmp/"+imgName)
	file, _ := os.Open("./tmp/" + imgName)
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	contentType := http.DetectContentType(buffer)
	file.Close()
	file, _ = os.Open("./tmp/" + imgName)
	defer file.Close()
	if contentType == "image/png" {
		newTypeName := "./tmp/" + strings.TrimSuffix(img.Filename, path.Ext(path.Base(img.Filename))) + ".png"
		_ = os.Rename("./tmp/"+imgName, newTypeName)
		imgNew, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		out, err := os.Create("./tmp/" + strings.TrimSuffix(img.Filename, path.Ext(path.Base(img.Filename))) + ".jpg")
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		defer out.Close()
		err = jpeg.Encode(out, imgNew, &jpeg.Options{Quality: 95})
		if err != nil {
			fmt.Println(err)
			_ = c.AbortWithError(200, apiException.ImgTypeError)
			return
		}
		_ = os.Remove(newTypeName)
		imgName = strings.TrimSuffix(img.Filename, path.Ext(path.Base(img.Filename))) + ".jpg"
		file.Close()
		file, _ = os.Open("./tmp/" + imgName)
	}
	imgNew, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(200, apiException.ImgTypeError)
		return
	}
	fileName := uuid.NewString() + ".jpg"
	output, _ := os.Create("./img/" + fileName)
	defer output.Close()
	err = jpeg.Encode(output, imgNew, &jpeg.Options{Quality: 40})
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(200, apiException.ImgTypeError)
		return
	}
	_ = os.Remove("./tmp/" + imgName)
	utils.JsonSuccessResponse(c, config.GetWebpUrlKey()+fileName)
}

func DeleteRecord(c *gin.Context) {
	lostIdRaw := c.Query("lost_id")
	lostId, err := strconv.Atoi(lostIdRaw)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	publisher := getPublisher(c)
	record, err := lostAndFoundRecordServices.GetRecordById(lostId)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if *publisher != record.Publisher {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}
	err = lostAndFoundRecordServices.UpdateRecord(lostId, models.LostAndFoundRecord{
		Type:        record.Type,
		Campus:      record.Campus,
		Kind:        record.Kind,
		Content:     record.Content,
		Publisher:   *publisher,
		IsProcessed: true,
	})
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img1, config.GetWebpUrlKey()))
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img2, config.GetWebpUrlKey()))
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img3, config.GetWebpUrlKey()))
	utils.JsonSuccessResponse(c, nil)
}
