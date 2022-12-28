package models

import "gorm.io/gorm"

type Upload struct {
	gorm.Model
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FileSize int64  `json:"fileSize"`
}
