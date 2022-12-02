package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/poli/domain"
	"strings"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type poliService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &poliService{
		qry: repo,
	}
}

// AddPoli implements domain.Service
func (ps *poliService) AddPoli(newPoli domain.PoliCore) (domain.PoliCore, error) {
	res, err := ps.qry.Add(newPoli)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.PoliCore{}, errors.New("rejected from database")
		}

		return domain.PoliCore{}, errors.New("some problem on database")
	}
	return res, nil
}

// DeletePoli implements domain.Service
func (ps *poliService) DeletePoli(id uint) (domain.PoliCore, error) {
	res, err := ps.qry.Delete(id)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return domain.PoliCore{}, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return domain.PoliCore{}, errors.New(config.DATABASE_ERROR)
	}
	return res, nil
}

// GetAllPoli implements domain.Service
func (ps *poliService) GetAllPoli() ([]domain.PoliCore, error) {
	res, err := ps.qry.GetAll()
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	// if len(res) == 0 {
	// 	log.Info("no data")
	// 	return nil, errors.New(config.DATA_NOTFOUND)
	// }

	return res, nil
}

// UpdatePoli implements domain.Service
func (ps *poliService) UpdatePoli(upPoli domain.PoliCore, poliID uint) (domain.PoliCore, error) {
	res, err := ps.qry.Update(upPoli, poliID)
	if err != nil {
		return domain.PoliCore{}, errors.New(config.DATABASE_ERROR)
	}

	return res, nil
}
