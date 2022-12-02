package delivery

import "jadwaldokter/features/poli/domain"

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
	Nama_poli string `json:"nama_poli"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "reg":
		cnv := basic.(domain.PoliCore)
		res = AddResponse{ID: cnv.ID, Nama_poli: cnv.Nama_poli}
	case "all":
		var arr []AddResponse
		cnv := basic.([]domain.PoliCore)
		for _, val := range cnv {
			arr = append(arr, AddResponse{ID: val.ID, Nama_poli: val.Nama_poli})
		}
		res = arr
	}
	return res
}
