package migrations

import (
	"github.com/xuyuntech/wechatshop-admin/pkg/models"
	"github.com/xuyuntech/wechatshop-admin/pkg/db"
)

func Migration() {

	AutoMigrate(&models.Product{}, &models.ProductImage{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
