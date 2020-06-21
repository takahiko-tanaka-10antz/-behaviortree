package models

type User struct {
	HP     int
	SkillA SkillA
	SkillB SkillB
}

func NewUser(hp int, sa, sb float64) *User {
	return &User{
		HP:     hp,
		SkillA: *NewSkillA(sa, 50),
		SkillB: *NewSkillB(sb, 60),
	}
}
