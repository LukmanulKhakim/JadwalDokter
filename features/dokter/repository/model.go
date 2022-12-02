package repository

import (
	"jadwaldokter/features/dokter/domain"

	"gorm.io/gorm"
)

type Poli struct {
	gorm.Model
	Nama_poli string
	Dokters   []Dokter `gorm:"foreignKey:Poli_ID"`
	Jadwals   []Jadwal `gorm:"foreignKey:Poli_ID"`
}

type Dokter struct {
	gorm.Model
	Nama_dokter string
	Poli_ID     uint
	Nama_poli   string   `gorm:"-:migration" gorm:"<-"`
	Jadwals     []Jadwal `gorm:"foreignKey:Dokter_ID"`
}

type Jadwal struct {
	gorm.Model
	Hari        string
	Jam         string
	Dokter_ID   uint
	Poli_ID     uint
	Nama_dokter string `gorm:"-:migration" gorm:"<-"`
	Nama_poli   string `gorm:"-:migration" gorm:"<-"`
}

func FromDomain(dp domain.DokterCore) Dokter {
	return Dokter{
		Model:       gorm.Model{ID: dp.ID},
		Nama_dokter: dp.Nama_dokter,
		Poli_ID:     dp.Poli_ID,
		Nama_poli:   dp.Nama_poli,
	}
}

func ToDomain(d Dokter) domain.DokterCore {
	return domain.DokterCore{
		ID:          d.ID,
		Nama_dokter: d.Nama_dokter,
		Poli_ID:     d.Poli_ID,
		Nama_poli:   d.Nama_poli,
	}
}

func ToDomainArray(ad []Dokter) []domain.DokterCore {
	var res []domain.DokterCore
	for _, val := range ad {
		res = append(res, domain.DokterCore{ID: val.ID, Nama_dokter: val.Nama_dokter, Poli_ID: val.Poli_ID, Nama_poli: val.Nama_poli})
	}
	return res
}
