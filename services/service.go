package service

type Thing struct {
	X, Y, Z int
}

func Start() bool {
	return true
}

func Running() bool {
	return false
}

func NewThing(x int, y int, z int) (*Thing, error) {
	return &Thing{
		X: x,
		Y: y,
		Z: z,
	}, nil
}
