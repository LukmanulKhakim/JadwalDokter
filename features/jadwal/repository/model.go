package repository

import (
	"jadwaldokter/features/jadwal/domain"

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

func FromDomain(dp domain.JadwalCore) Jadwal {
	return Jadwal{
		Model:       gorm.Model{ID: dp.ID},
		Hari:        dp.Hari,
		Jam:         dp.Jam,
		Dokter_ID:   dp.Dokter_ID,
		Poli_ID:     dp.Poli_ID,
		Nama_dokter: dp.Nama_dokter,
		Nama_poli:   dp.Nama_poli,
	}
}

func ToDomain(j Jadwal) domain.JadwalCore {
	return domain.JadwalCore{
		ID:          j.ID,
		Hari:        j.Hari,
		Jam:         j.Jam,
		Poli_ID:     j.Poli_ID,
		Dokter_ID:   j.Dokter_ID,
		Nama_dokter: j.Nama_dokter,
		Nama_poli:   j.Nama_poli,
	}
}

func ToDomainArray(aj []Jadwal) []domain.JadwalCore {
	var res []domain.JadwalCore
	for _, val := range aj {
		res = append(res, domain.JadwalCore{
			ID:          val.ID,
			Hari:        val.Hari,
			Jam:         val.Jam,
			Poli_ID:     val.Poli_ID,
			Dokter_ID:   val.Dokter_ID,
			Nama_dokter: val.Nama_dokter,
			Nama_poli:   val.Nama_poli,
		})
	}
	return res
}
