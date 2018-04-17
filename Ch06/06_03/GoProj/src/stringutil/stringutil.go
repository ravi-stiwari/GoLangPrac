package stringutil

//FullName firstname lastname
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

//FullNameNakedReturn firstname lastname
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
