package models

type SkillA struct {
	Rate   float64
	Damage int
}

type SkillB struct {
	Rate   float64
	Damage int
}

func NewSkillA(r float64, d int) *SkillA {
	return &SkillA{
		Rate:   r,
		Damage: d,
	}
}

func NewSkillB(r float64, d int) *SkillB {
	return &SkillB{
		Rate:   r,
		Damage: d,
	}
}
