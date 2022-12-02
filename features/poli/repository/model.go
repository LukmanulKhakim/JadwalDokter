package repository

import (
	"jadwaldokter/features/poli/domain"

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

func FromDomain(dp domain.PoliCore) Poli {
	return Poli{
		Model:     gorm.Model{ID: dp.ID},
		Nama_poli: dp.Nama_poli,
	}
}

func ToDomain(p Poli) domain.PoliCore {
	return domain.PoliCore{
		ID:        p.ID,
		Nama_poli: p.Nama_poli,
	}
}

func ToDomainArray(ap []Poli) []domain.PoliCore {
	var res []domain.PoliCore
	for _, val := range ap {
		res = append(res, domain.PoliCore{ID: val.ID, Nama_poli: val.Nama_poli})
	}

	return res
}
