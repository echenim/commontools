package tools

//TrimLeft function to remove space from data from the left
func TrimLeft(data string) string {
	rs := ""
	for i := 0; i < len(data); i++ {

		if string(data[i]) != " " {

			for j := i; j < len(data); j++ {
				rs = rs + string(data[j])
			}
			return rs
		}
	}
	return rs
}

//SpacePaddingLeft function to pad data with space from the left
func SpacePaddingLeft(expectedlenght int, data string) string {
	acturallenght := len(data)
	for i := acturallenght; i < expectedlenght; i++ {
		data = " " + data
	}
	return data
}
