package models

type Lesson struct {
	ID          int    `json:"ID"`
	Username    string `json:"username"`
	Campus      string `json:"campus"`
	LessonName  string `json:"lessonName"`
	LessonPlace string `json:"lessonPlace"`
	Sections    string `json:"sections"`
	Week        string `json:"week"`
	Weekday     string `json:"weekday"`
	Term        string `json:"term"`
	Year        string `json:"year"`
}
