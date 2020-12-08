package clase1

import "testing"

/*
ejecutar una funcion especifica
go test -v run  (nombre funcion)
go test -v run  (nombre funcion)$  final de la exp
*/
func TestAdd(t *testing.T) {
	want := 9
	got := Add(2, 3)
	if got != want {
		t.Logf("error : se esperab %d se obtuvo %d", want, got)
		t.Fail()
	}
}

func TestAddMultiple(t *testing.T) {
	want := 6
	got := AddMultiple(1, 2, 3)
	if got != want {
		t.Logf("error : se esperab %d se obtuvo %d", want, got)
		t.Fail()
	}
}
