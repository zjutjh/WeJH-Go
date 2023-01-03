package lessonController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/lessonServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type lessonForm struct {
	ID          int    `json:"ID"`
	Campus      string `json:"campus" binding:"required"`
	LessonName  string `json:"lessonName" binding:"required"`
	LessonPlace string `json:"lessonPlace" binding:"required"`
	Sections    string `json:"sections" binding:"required"`
	Week        string `json:"week" binding:"required"`
	Weekday     string `json:"weekday" binding:"required"`
	Term        string `json:"term" binding:"required"`
	Year        string `json:"year" binding:"required"`
}

type fetchForm struct {
	Term string `json:"term" binding:"required"`
	Year string `json:"year" binding:"required"`
}

type deleteForm struct {
	ID int `json:"ID" binding:"required"`
}

func CreateLesson(c *gin.Context) {
	var postForm lessonForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	err = lessonServices.CreateLesson(models.Lesson{
		Username:    user.Username,
		Campus:      postForm.Campus,
		LessonPlace: postForm.LessonPlace,
		LessonName:  postForm.LessonName,
		Sections:    postForm.Sections,
		Week:        postForm.Week,
		Weekday:     postForm.Weekday,
		Term:        postForm.Term,
		Year:        postForm.Year,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func GetLesson(c *gin.Context) {
	var postForm fetchForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	lessons, err := lessonServices.GetLesson(user.Username, postForm.Term, postForm.Year)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, lessons)
}

func UpdateLesson(c *gin.Context) {
	var postForm lessonForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil || postForm.ID == 0 {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	err = lessonServices.UpdateLesson(postForm.ID, models.Lesson{
		Username:    user.Username,
		Campus:      postForm.Campus,
		LessonPlace: postForm.LessonPlace,
		LessonName:  postForm.LessonName,
		Sections:    postForm.Sections,
		Week:        postForm.Week,
		Weekday:     postForm.Weekday,
		Term:        postForm.Term,
		Year:        postForm.Year,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func DeleteLesson(c *gin.Context) {
	var postForm deleteForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = lessonServices.DeleteLesson(postForm.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
