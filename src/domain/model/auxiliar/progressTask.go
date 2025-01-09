package auxiliar

import "gorm.io/gorm"

type ProgressTask struct {
	gorm.Model
	ID int64 `gorm:"primaryKey;autoIncrement:true"`
	TaskId    int 
	ProgresId int
}
