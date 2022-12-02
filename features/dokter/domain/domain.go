package domain

type DokterCore struct {
	ID          uint
	Nama_dokter string
	Poli_ID     uint
	Nama_poli   string
}

type Repository interface {
	Add(newItem DokterCore) (DokterCore, error)
	GetAll() ([]DokterCore, error)
	Delete(ID uint) (DokterCore, error)
}

type Service interface {
	AddDokter(newItem DokterCore) (DokterCore, error)
	GetDokter() ([]DokterCore, error)
	DeleteDokter(ID uint) (DokterCore, error)
}
