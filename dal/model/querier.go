package model

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error)
}
