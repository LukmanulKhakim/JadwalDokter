package helper

func Convert(day string) string {
	var res string
	if day == "Saturday" {
		res = "sabtu"
	} else if day == "Sunday" {
		res = "minggu"
	} else if day == "Monday" {
		res = "senin"
	} else if day == "Tuesday" {
		res = "selasa"
	} else if day == "Wednesday" {
		res = "rabu"
	} else if day == "Thursday" {
		res = "kamis"
	} else {
		res = "jumat"
	}
	return res
}
