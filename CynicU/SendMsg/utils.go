package SendMsg


const (
	AckLength = 16
)

func IsEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	l := len(a)
	for i := 0 ; i < l ; i ++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}