package delivery

import "jadwaldokter/features/dokter/domain"

type AddFormat struct {
	Nama_dokter string `json:"nama_dokter" form:"nama_dokter"`
	Poli_ID     uint   `json:"poli_id" form:"poli_id"`
}

func ToDomain(i interface{}) domain.DokterCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.DokterCore{Nama_dokter: cnv.Nama_dokter, Poli_ID: cnv.Poli_ID}
	}
	return domain.DokterCore{}
}
