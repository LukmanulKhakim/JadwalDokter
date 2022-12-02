package repository

import (
	"jadwaldokter/features/dokter/domain"

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
func (rq *repoQuery) Add(newItem domain.DokterCore) (domain.DokterCore, error) {
	var cnv Dokter
	cnv = FromDomain(newItem)
	if err := rq.db.Select("nama_dokter", "poli_id").Create(&cnv).Error; err != nil {
		log.Error("Error on insert user", err.Error())
		return domain.DokterCore{}, err
	}
	// selesai dari DB
	newItem = ToDomain(cnv)
	return newItem, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(ID uint) (domain.DokterCore, error) {
	var resQry Dokter

	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}
	if err := rq.db.Delete(&resQry).Error; err != nil {
		log.Error(err.Error())
		return ToDomain(resQry), err
	}

	res := ToDomain(resQry)
	return res, nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.DokterCore, error) {
	var resQry []Dokter
	// if err := rq.db.Find(&resQry).Error; err != nil {
	// 	log.Error("Error on All user", err.Error())
	// 	return nil, err
	// }
	if err := rq.db.Table("dokters").Select("dokters.id", "dokters.nama_dokter", "polis.nama_poli").Joins("join polis on polis.id = dokters.poli_id").Scan(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
