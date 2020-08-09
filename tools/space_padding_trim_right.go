package tools

//Trim function to remove space from data
func Trim(data string) string {
	rs := ""
	for i := 0; i < len(data); i++ {

		if string(data[i]) != " " {
			rs = rs + string(data[i])
		}
	}
	return rs
}
