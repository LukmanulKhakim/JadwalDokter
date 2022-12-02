package repository

import (
	"jadwaldokter/features/jadwal/domain"

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
func (rq *repoQuery) Add(newJadwal domain.JadwalCore) (domain.JadwalCore, error) {
	var cnv Jadwal
	cnv = FromDomain(newJadwal)
	if err := rq.db.Select("hari", "jam", "dokter_id", "poli_id").Create(&cnv).Error; err != nil {
		log.Error("Error on insert user", err.Error())
		return domain.JadwalCore{}, err
	}
	// selesai dari DB
	newJadwal = ToDomain(cnv)
	return newJadwal, nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(id uint) (domain.JadwalCore, error) {
	var resQry Jadwal

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
func (rq *repoQuery) GetAll(hari string) ([]domain.JadwalCore, error) {
	var resQry []Jadwal

	if err := rq.db.Table("jadwals").Select("jadwals.id", "jadwals.hari", "jadwals.jam", "dokters.nama_dokter", "polis.nama_poli").Joins("join dokters on dokters.id = jadwals.dokter_id").Joins("join polis on polis.id = jadwals.poli_id").Scan(&resQry).Where(" hari LIKE ? OR jam LIKE ? OR dokters.nama_dokter LIKE ? OR polis.nama_poli LIKE ?",
		"%"+hari+"%", "%"+hari+"%", "%"+hari+"%", "%"+hari+"%").Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	res := ToDomainArray(resQry)
	return res, nil
}
