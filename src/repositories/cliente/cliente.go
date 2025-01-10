package cliente

import (
	clienteDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/cliente"
	clienteModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/cliente"
	"gorm.io/gorm"
)

func SalvaCliente(db *gorm.DB, clientes []string) error {
	for _, value := range clientes {
		var clienteDto = &clienteDto.Cliente{
			Name: value,
		}
		var cliente = clienteModel.NewCliente(clienteDto)
		result := db.First(&cliente, "Nome = ?", cliente.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&cliente)
			if res.Error != nil {
				return res.Error
			}
		} else {
			db.Save(&cliente)
		}
	}

	return nil
}
