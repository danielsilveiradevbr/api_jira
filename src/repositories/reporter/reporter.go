package reporter

import (
	reporterDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	reporterModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func SalvaReporter(db *gorm.DB, reporterDTO *reporterDTO.Reporter) error {
	reporter := reporterModel.NewReporter(reporterDTO)
	result := db.First(&reporter, "email = ?", reporter.Email)
	if result.RowsAffected == 0 {
		res := db.Create(&reporter)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&reporter)
	}
	return nil
}
