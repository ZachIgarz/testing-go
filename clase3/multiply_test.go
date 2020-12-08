package clase3

import "testing"
/*

go test -run multiply/2x3
*/
func TestMultiply(t *testing.T) {

	table := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"2x1", 2, 1, 2},
		{"2x2", 2, 2, 4},
		{"2x3", 2, 3, 6},
		{"2x4", 2, 4, 8},
		{"2x5", 2, 5, 10},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			got := Multiply(v.a, v.b)
			if got != v.want {
				t.Errorf("se esperaba %d y se obtuvo %d", v.want, got)
			}
		})
	}
}
