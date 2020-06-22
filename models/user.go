package models

// プレイするユーザーの構造体
type User struct {
	HP     int
	Skill1 Skill1
	Skill2 Skill2
}

func NewUser(hp int, sa, sb float64) *User {
	return &User{
		HP:     hp,
		Skill1: *NewSkill1(sa, 50),
		Skill2: *NewSkill2(sb, 60),
	}
}
