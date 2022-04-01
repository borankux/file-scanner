package database

type File struct {
	Name string `json:"name" gorm:"primary_key"`
	Size uint64 `json:"size"`
}
