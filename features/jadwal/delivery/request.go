package delivery

import "jadwaldokter/features/jadwal/domain"

type AddFormat struct {
	Hari      string `json:"hari" form:"hari"`
	Jam       string `json:"jam" form:"jam"`
	Dokter_ID uint   `json:"dokter_id" form:"dokter_id"`
	Poli_ID   uint   `json:"poli_id" form:"poli_id"`
}

func ToDomain(i interface{}) domain.JadwalCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.JadwalCore{Hari: cnv.Hari, Jam: cnv.Jam, Dokter_ID: cnv.Dokter_ID, Poli_ID: cnv.Poli_ID}
	}
	return domain.JadwalCore{}
}
