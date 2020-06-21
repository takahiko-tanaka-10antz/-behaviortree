package models

// 敵の構造体
type Enemy struct {
	HP int
}

func NewEnemy(hp int) *Enemy {
	return &Enemy{
		HP: hp,
	}
}
