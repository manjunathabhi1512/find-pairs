package repository

import "gorm.io/gorm"

type CommonRepository interface{
}

type CommonRepo struct{
	DB *gorm.DB
}

func NewCommonRepo(Db *gorm.DB) CommonRepo{
	return CommonRepo{DB: Db}
}