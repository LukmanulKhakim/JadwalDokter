package delivery

import "jadwaldokter/features/poli/domain"

type AddFormat struct {
	Nama_poli string `json:"nama_poli" form:"nama_poli"`
}

type UpdateFormat struct {
	Nama_poli string `json:"nama_poli" form:"nama_poli"`
}

func ToDomain(i interface{}) domain.PoliCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.PoliCore{Nama_poli: cnv.Nama_poli}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.PoliCore{Nama_poli: cnv.Nama_poli}
	}
	return domain.PoliCore{}
}
