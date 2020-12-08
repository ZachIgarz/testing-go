package clase1

import "testing"

/*
ejecutar una funcion especifica
go test -v run  (nombre funcion)
go test -v run  (nombre funcion)$  final de la exp
go test -v muestra logs
*/
func TestAdd(t *testing.T) {
	want := 5
	t.Logf("want vale %d \n", want)
	got := Add(2, 3)
	t.Logf("got vale %d \n", got)
	if got != want {
		t.Errorf("error : se esperab %d se obtuvo %d", want, got)

	}
	t.Log("Termino la prueba de Add")
}

func TestAddMultiple(t *testing.T) {
	want := 6
	got := AddMultiple(1, 2, 3)
	if got != want {
		t.Errorf("error : se esperab %d se obtuvo %d", want, got)
	}
}
