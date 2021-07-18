package appcontainer

import (
	"github.com/golobby/container/v2"
	"gorm.io/gorm"
)

type AppContainer struct {
	DBOrm *gorm.DB
}

func Init(dbOrm *gorm.DB) error {
	return container.Singleton(func() AppContainer {
		return AppContainer{
			DBOrm: dbOrm,
		}
	})
}

func GetInstance() *AppContainer {
	var appContainer AppContainer
	_ = container.Bind(&appContainer)
	return &appContainer
}
