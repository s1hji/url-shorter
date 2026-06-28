package generator

import "testing"

func TestLen(t *testing.T) {
	length, err := Rand_generate()
	if err != nil {
		t.Error(err)
	}
	if len(length) != 10 {
		t.Errorf("длина сгенерируемой ссылки не равна 10, получили длину: %d ", len(length))
	}
}
