package models

type ThemePermission struct {
	ID              int    `json:"id"`
	StudentID       string `json:"student_id"`
	CurrentThemeID  int    `json:"current_theme_id"`
	ThemePermission string `json:"theme_permission"`
}

type ThemePermissionData struct {
	ThemeIDs []int `json:"theme_id"`
}
