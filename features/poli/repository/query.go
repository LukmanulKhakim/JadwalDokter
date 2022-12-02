package repository

import (
	"jadwaldokter/features/poli/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}

// Add implements domain.Repository
func (rq *repoQuery) Add(newPoli domain.PoliCore) (domain.PoliCore, error) {
	var cnv Poli
	cnv = FromDomain(newPoli)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("Error on insert user", err.Error())
		return domain.PoliCore{}, err
	}
	// selesai dari DB
	newPoli = ToDomain(cnv)
	return newPoli, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(id uint) (domain.PoliCore, error) {
	var resQry Poli

	if err := rq.db.First(&resQry, "ID = ?", id).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}
	if err := rq.db.Unscoped().Delete(&resQry).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}

	res := ToDomain(resQry)
	return res, nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.PoliCore, error) {
	var resQry []Poli
	if err := rq.db.Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(upPoli domain.PoliCore, poliID uint) (domain.PoliCore, error) {
	var cnv Poli = FromDomain(upPoli)
	if err := rq.db.Table("users").Where("id = ?", poliID).Updates(&cnv).Error; err != nil {
		log.Error("error on updating user", err.Error())
		return domain.PoliCore{}, err
	}

	res := ToDomain(cnv)
	return res, nil
}
