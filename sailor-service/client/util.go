package client

type BinaryInteger interface {
	int | int32
}

func IntToArray[T BinaryInteger](dataKind T) []T {
	var val T
	array := make([]T, 0, 6)
	for i := 0; i < 6 && dataKind >= 1<<i; i++ {
		val = dataKind & (1 << i)
		if val > 0 {
			array = append(array, val)
		}
	}
	return array
}

func ArrayToInt[T BinaryInteger](ds []T) T {
	var val T
	for _, d := range ds {
		val += d
	}
	return val
}
