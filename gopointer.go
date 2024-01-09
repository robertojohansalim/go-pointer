package gopointer

func Pointerize[K any](tmp K) *K {
	return &tmp
}
