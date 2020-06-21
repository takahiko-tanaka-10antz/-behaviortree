package models

type Enemy struct {
	HP int
}

func NewEnemy(hp int) *Enemy {
	return &Enemy{
		HP: hp,
	}
}
