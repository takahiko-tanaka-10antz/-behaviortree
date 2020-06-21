package models

// スキルAの構造体
type SkillA struct {
	Rate   float64
	Damage int
}

// スキルBの構造体
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
