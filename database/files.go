package database

type File struct {
	Name string `json:"name" gorm:"primary_key"`
	Size int64  `json:"size"`
}
