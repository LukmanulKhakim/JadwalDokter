package delivery

import "jadwaldokter/features/jadwal/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type AddResponse struct {
	ID        uint   `json:"id"`
	Hari      string `json:"hari"`
	Jam       string `json:"jam"`
	Dokter_ID uint   `json:"dokter_id"`
	Poli_ID   uint   `json:"poli_id"`
}

type GetResponse struct {
	ID          uint   `json:"id"`
	Hari        string `json:"hari"`
	Jam         string `json:"jam"`
	Nama_dokter string `json:"nama_dokter"`
	Nama_poli   string `json:"nama_poli"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := basic.(domain.JadwalCore)
		res = AddResponse{ID: cnv.ID, Hari: cnv.Hari, Jam: cnv.Jam, Dokter_ID: cnv.Dokter_ID, Poli_ID: cnv.Poli_ID}
	case "all":
		var arr []GetResponse
		cnv := basic.([]domain.JadwalCore)
		for _, val := range cnv {
			arr = append(arr, GetResponse{ID: val.ID, Hari: val.Hari, Jam: val.Jam, Nama_dokter: val.Nama_dokter, Nama_poli: val.Nama_poli})
		}
		res = arr
	}
	return res
}
