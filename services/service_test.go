package service_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	service "github.com/pertrai1/power-go-tests/services"
)

func TestService(t *testing.T) {
	t.Parallel()
	want := true
	got := service.Running()
	t.Log("got is ", got)
	service.Start()
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestServiceOpenFile(t *testing.T) {
	_, err := os.Open("testing.txt")
	if err != nil {
		t.Fatal(err)
	}
}

// func TestNewThing(t *testing.T) {
// 	t.Parallel()
// 	x, y, z := 1, 2, 3
// 	got, err := service.NewThing(x, y, z)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if got.X != x {
// 		t.Errorf("want X: %v, got X: %v", x, got.X)
// 	}
// 	if got.Y != y {
// 		t.Errorf("want Y: %v, got Y: %v", y, got.Y)
// 	}
// 	if got.Z != z {
// 		t.Errorf("want Z: %v, got Z: %v", z, got.Z)
// 	}
// }

func TestNewThing(t *testing.T) {
	t.Parallel()
	x, y, z := 1, 2, 3
	want := &service.Thing{
		X: x,
		Y: y,
		Z: z,
	}
	got, err := service.NewThing(x, y, z)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
