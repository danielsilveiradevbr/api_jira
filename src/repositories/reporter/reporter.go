package reporter

import (
	reporterDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/reporter"
	reporterModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/user"
	"gorm.io/gorm"
)

func SalvaReporter(db *gorm.DB, reporterDTO *reporterDto.Reporter) (*reporterModel.User, error) {
	reporter := reporterModel.NewReporter(db, reporterDTO)
	result := db.First(&reporter, "email = ?", reporter.Email)
	if result.RowsAffected == 0 {
		res := db.Create(&reporter)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&reporter)
	}
	return reporter, nil
}
