package domain

type JadwalCore struct {
	ID          uint
	Hari        string
	Jam         string
	Dokter_ID   uint
	Poli_ID     uint
	Nama_dokter string
	Nama_poli   string
}

type Repository interface {
	Add(newJadwal JadwalCore) (JadwalCore, error)
	GetAll(hari string) ([]JadwalCore, error)
	Delete(id uint) (JadwalCore, error)
}

type Service interface {
	AddJadwal(newJadwal JadwalCore) (JadwalCore, error)
	GetJadwal(hari string) ([]JadwalCore, error)
	DeleteJadwal(id uint) (JadwalCore, error)
}
