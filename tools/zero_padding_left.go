package tools

//RemoveZeroPaddingFromLeft function to remove zero from data from the left
func RemoveZeroPaddingFromLeft(data string) string {
	rs := ""
	for i := 0; i < len(data); i++ {

		if string(data[i]) != "0" {

			for j := i; j < len(data); j++ {
				rs = rs + string(data[j])
			}
			return rs
		}
	}
	return rs
}

//ZeroPaddingLeft function to pad data from the left
func ZeroPaddingLeft(expectedlenght int, data string) string {
	acturallenght := len(data)
	for i := acturallenght; i < expectedlenght; i++ {
		data = "0" + data
	}
	return data
}
