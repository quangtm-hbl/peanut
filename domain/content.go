package domain

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Thumbnail   string
	Name        string
	Media       string
	Description string
	Playtime    int
	Resolution  int
	ARwidth     int
	ARheight    int
	Fever       bool
	Ondemand    bool
}
