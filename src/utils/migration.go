package utils

import (
	m "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&m.Pedidos{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&m.PedProds{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&m.PedidosLog{})
}
