package tipoUser

import (
	tipoUserDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/tipoUser"
	tipoUserModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/tipoUser"
	"gorm.io/gorm"
)

func SalvaTipoUser(db *gorm.DB, tipoUserDto *tipoUserDto.TipoUser) (*tipoUserModel.TipoUser, error) {
	tipoUser := tipoUserModel.NewTipoUser(tipoUserDto)
	result := db.First(&tipoUser, "nome = ?", tipoUser.Nome)
	if result.RowsAffected == 0 {
		res := db.Create(&tipoUser)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&tipoUser)
	}
	return tipoUser, nil
}
