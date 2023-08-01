package repository

import (
	"context"
	"database/sql"
	"emansulaeman/belajar-golang-restful-api/model/domain"
)

// interface ini dibuat untuk contract category
type CategoryRepository interface {
	// parameter yang kedua sql.Tx untuk query yang jenisnya transactional
	// oarameter yang ketiga merupakan data yang aslinya
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category   // merupakan return valuenya
	Update(ctx context.Context, tx *sql.Tx, categroy domain.Category) domain.Category // merupakan return valuenya
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) // merupakan return valuenya
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category                         // slice berupa return valuenya
}
