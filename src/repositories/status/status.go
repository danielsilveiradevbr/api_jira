package status

import (
	statusDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/status"
	statusModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/status"
	statusCategoryRep "github.com/danielsilveiradevbr/api_jira/src/repositories/statusCategory"
	"gorm.io/gorm"
)

func SalvaStatus(db *gorm.DB, statusDTO *statusDto.Status) (*statusModel.Status, error) {
	statusCategory, err := statusCategoryRep.SalvaStatusCategory(db, &statusDTO.StatusCategory)
	if err != nil {
		return nil, err
	}
	status := statusModel.NewStatus(statusDTO)
	status.StatusCategoryId = statusCategory.ID
	result := db.First(&status, "descricao = ?", status.DESCRICAO)
	if result.RowsAffected == 0 {
		res := db.Create(&status)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&status)
	}
	return status, nil
}
