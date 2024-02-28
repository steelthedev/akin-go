package connection

import (
	"log"

	"github.com/steelthedev/akin-go/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&pkg.Project{})

	return db
}
