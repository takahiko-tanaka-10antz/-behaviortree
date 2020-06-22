package models

// スキルAの構造体
type Skill1 struct {
	Rate   float64
	Damage int
}

// スキルBの構造体
type Skill2 struct {
	Rate   float64
	Damage int
}

func NewSkill1(r float64, d int) *Skill1 {
	return &Skill1{
		Rate:   r,
		Damage: d,
	}
}

func NewSkill2(r float64, d int) *Skill2 {
	return &Skill2{
		Rate:   r,
		Damage: d,
	}
}
