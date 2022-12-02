package domain

type PoliCore struct {
	ID        uint
	Nama_poli string
}

type Repository interface {
	Add(newPoli PoliCore) (PoliCore, error)
	GetAll() ([]PoliCore, error)
	Get(id uint) (PoliCore, error)
	Update(upPoli PoliCore, poliID uint) (PoliCore, error)
	Delete(id uint) (PoliCore, error)
}

type Service interface {
	AddPoli(newPoli PoliCore) (PoliCore, error)
	GetAllPoli() ([]PoliCore, error)
	GetPoli(id uint) (PoliCore, error)
	UpdatePoli(upPoli PoliCore, poliID uint) (PoliCore, error)
	DeletePoli(id uint) (PoliCore, error)
}
