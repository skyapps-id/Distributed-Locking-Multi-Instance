package repository

import (
	"context"
	"distributed_locking/driver"
	"distributed_locking/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Promo interface {
	GetPromoByCode(ctx context.Context, code string) (result entity.Promo, err error)
	DecrimentPromoByCode(ctx context.Context, code string) (err error)
	DecrimentPromoByCode2(ctx context.Context, code string) (err error)
}

type promo struct {
	db  *gorm.DB
	rds driver.RedisClient
}

func NewPromoRepository(db *gorm.DB, rds driver.RedisClient) Promo {
	if db == nil {
		panic("db is nil")
	}

	return &promo{
		db:  db,
		rds: rds,
	}
}

func (r *promo) GetPromoByCode(ctx context.Context, code string) (result entity.Promo, err error) {
	err = r.db.WithContext(ctx).Table("promos").Where("promos.code = ?", code).Take(&result).Error
	if err != nil {
		return
	}

	return
}

func (r *promo) DecrimentPromoByCode(ctx context.Context, code string) (err error) {
	tx := r.db.Begin()

	var result entity.Promo
	err = tx.WithContext(ctx).Table("promos").Clauses(clause.Locking{Strength: "UPDATE"}).Where("promos.code = ?", code).Take(&result).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if result.Qouta == 0 {
		tx.Rollback()
		err = errors.New("coupon expired")
		return
	}

	err = tx.WithContext(ctx).Table("promos").Where("promos.code = ?", code).Update("qouta", result.Qouta-1).Error
	if err != nil {
		return
	}

	tx.Commit()

	return
}

func (r *promo) DecrimentPromoByCode2(ctx context.Context, code string) (err error) {
	rs := r.rds.Redsync()
	mutex := rs.NewMutex("promo-lock-" + code)
	mutex.Lock()
	defer mutex.Unlock()

	tx := r.db.Begin()

	var result entity.Promo
	err = tx.WithContext(ctx).Table("promos").Where("promos.code = ?", code).Take(&result).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if result.Qouta == 0 {
		tx.Rollback()
		err = errors.New("coupon expired")
		return
	}

	err = tx.WithContext(ctx).Table("promos").Where("promos.code = ?", code).Update("qouta", result.Qouta-1).Error
	if err != nil {
		return
	}

	tx.Commit()

	return
}
