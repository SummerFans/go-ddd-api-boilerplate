package postgres

import (
	"database/sql"
	"go-ddd-api-boilerplate/domain/entity"
	"go-ddd-api-boilerplate/domain/repository"

	_ "github.com/lib/pq"
)

type goodsRepository struct {
	db *sql.DB
}

func (r *goodsRepository) Save(goods *entity.Goods) (int64, error) {
	_, err := r.db.Exec("INSERT INTO table (id) VALUES (?)", "1")

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func NewGoodsRepository(
	db *sql.DB,
) repository.GoodsRepository {
	return &goodsRepository{
		db: db,
	}
}
