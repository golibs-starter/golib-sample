package testing

import (
	"fmt"
	"gitlab.com/golibs-starter/golib/log"
	"gorm.io/gorm"
)

func TruncateTablesInvoker(tables ...string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		TruncateTables(db, tables...)
	}
}

func TruncateTables(db *gorm.DB, tables ...string) {
	for _, table := range tables {
		TruncateTable(db, table)
	}
}

func TruncateTableInvoker(table string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		TruncateTable(db, table)
	}
}

func TruncateTable(db *gorm.DB, table string) {
	if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", table)).Error; err != nil {
		log.Fatalf("Could not truncate table %s: %v", table, err)
	}
}

func SeedInvoker(model interface{}) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		Seed(db, model)
	}
}

func Seed(db *gorm.DB, model interface{}) {
	if err := db.Create(model).Error; err != nil {
		log.Fatalf("Could not create seed data, model: %v, err: %v", model, err)
	}
}
