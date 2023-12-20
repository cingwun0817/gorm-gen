package main

// gorm gen configure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gen"

	"gen/dal/model"
)

const MySQLDSN = "root:admin@tcp(127.0.0.1:3306)/gen?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(connectDB(MySQLDSN))

	g.ApplyBasic(g.GenerateAllTable()...)

	g.ApplyInterface(func(model.Querier) {}, g.GenerateAllTable()...)

	g.Execute()
}
