package models

type Theme struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	IsDarkMode  bool   `json:"is_dark_mode"`
	ThemeConfig string `json:"theme_config"`
}

type BarIcon struct {
	HomeIcon             string `json:"home_icon"`
	FunctionIcon         string `json:"function_icon"`
	MyIcon               string `json:"my_icon"`
	SelectedHomeIcon     string `json:"selected_home_icon"`
	SelectedFunctionIcon string `json:"selected_function_icon"`
	SelectedMyIcon       string `json:"selected_my_icon"`
}

type AppListIcon struct {
	ClassIcon         string `json:"class_icon"`
	GradeIcon         string `json:"grade_icon"`
	ExamIcon          string `json:"exam_icon"`
	FreeClassroomIcon string `json:"free_classroom_icon"`
	SchoolcardIcon    string `json:"schoolcard_icon"`
	LendIcon          string `json:"lend_icon"`
	ElectricityIcon   string `json:"electricity_icon"`
	SchoolbusIcon     string `json:"schoolbus_icon"`
	ClothIcon         string `json:"cloth_icon"`
}

type BaseColor struct {
	Base500 string `json:"base_500"`
	Base600 string `json:"base_600"`
	Base700 string `json:"base_700"`
}

type ThemeConfig struct {
	BarIcon            BarIcon     `json:"bar_icon"`
	AppListIcon        AppListIcon `json:"applist_icon"`
	AppListDarkIcon    AppListIcon `json:"applist_dark_icon"`
	BackgroundImg      string      `json:"background_img"`
	BackgroundColor    string      `json:"background_color"`
	BackgroundPosition string      `json:"background_position"`
	BaseColor          BaseColor   `json:"base_color"`
	SelectionImg       string      `json:"selection_img"`
}

type FormatTheme struct {
	ThemeID     int         `json:"theme_id"`
	Name        string      `json:"name"`
	IsDarkMode  bool        `json:"is_dark_mode"`
	ThemeConfig ThemeConfig `json:"theme_config"`
}
