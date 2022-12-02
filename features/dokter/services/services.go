package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/dokter/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type dokterService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &dokterService{
		qry: repo,
	}
}

// AddDokter implements domain.Service
func (ds *dokterService) AddDokter(newItem domain.DokterCore) (domain.DokterCore, error) {
	res, err := ds.qry.Add(newItem)
	if err != nil {
		return domain.DokterCore{}, errors.New("some problem on database")
	}

	return res, nil
}

// DeleteDokter implements domain.Service
func (ds *dokterService) DeleteDokter(ID uint) (domain.DokterCore, error) {
	res, err := ds.qry.Delete(ID)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return domain.DokterCore{}, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return domain.DokterCore{}, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

// GetDokter implements domain.Service
func (ds *dokterService) GetDokter() ([]domain.DokterCore, error) {
	res, err := ds.qry.GetAll()
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	return res, nil
}
