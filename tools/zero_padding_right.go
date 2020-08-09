package tools

// //RemoveZeroPaddingFromRight function to remove zero from data from the right
// func RemoveZeroPaddingFromRight(data string) string {
// 	rs := ""
// 	for i := 0; i < len(data); i++ {

// 		if string(data[i]) != "0" {

// 			for j := i; j < len(data); j++ {
// 				rs = rs + string(data[j])
// 			}
// 			return rs
// 		}
// 	}
// 	return rs
// }

//ZeroPaddingRight function to pad data from the right
func ZeroPaddingRight(expectedlenght int, data string) string {
	acturallenght := len(data)
	for i := acturallenght; i < expectedlenght; i++ {
		data = data + "0"
	}
	return data
}
