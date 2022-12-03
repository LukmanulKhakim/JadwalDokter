package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/jadwal/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type jadwalService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &jadwalService{
		qry: repo,
	}
}

// AddJadwal implements domain.Service
func (js *jadwalService) AddJadwal(newJadwal domain.JadwalCore) (domain.JadwalCore, error) {
	res, err := js.qry.Add(newJadwal)
	if err != nil {
		return domain.JadwalCore{}, errors.New("some problem on database")
	}

	return res, nil
}

// DeleteJadwal implements domain.Service
func (js *jadwalService) DeleteJadwal(id uint) (domain.JadwalCore, error) {
	res, err := js.qry.Delete(id)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return domain.JadwalCore{}, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return domain.JadwalCore{}, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

// GetJadwal implements domain.Service
func (js *jadwalService) GetJadwal(hari string) ([]domain.JadwalCore, error) {
	res, err := js.qry.GetAll(hari)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}

	return res, nil
}
