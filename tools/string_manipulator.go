package tools

//IsNillOrEmpty function to envaluater if a strin is empty
func IsNillOrEmpty(value string) bool {
	if value == "" || len(value) == 0 {
		return true
	}
	return false
}
