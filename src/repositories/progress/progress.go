package progress

import (
	progressDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/progress"
	progressModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/progress"
	"gorm.io/gorm"
)

func SalvaProgress(db *gorm.DB, progressDTO *progressDto.Progress) (*progressModel.Progress, error) {
	progress := progressModel.NewProgress(progressDTO)
	result := db.First(&progress, "progress = ? and total = ? and percent = ?", progress.Progress, progress.Total, progress.Percent)
	if result.RowsAffected == 0 {
		res := db.Create(&progress)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&progress)
	}
	return progress, nil
}
