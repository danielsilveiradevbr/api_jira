package cliente

import (
	clienteDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/cliente"
	clienteModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/cliente"
	"gorm.io/gorm"
)

func SalvaCliente(db *gorm.DB, clientes []string) (*clienteModel.Cliente, error) {
	var cliente = &clienteModel.Cliente{}
	for _, value := range clientes {
		var clienteDto = &clienteDto.Cliente{
			Name: value,
		}
		cliente = clienteModel.NewCliente(clienteDto)
		result := db.First(&cliente, "Nome = ?", cliente.Nome)
		if result.RowsAffected == 0 {
			res := db.Create(&cliente)
			if res.Error != nil {
				return nil, res.Error
			}
		} else {
			db.Save(&cliente)
		}
	}

	return cliente, nil
}
