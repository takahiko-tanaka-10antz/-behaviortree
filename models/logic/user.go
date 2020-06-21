package logic

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/takahiko-tanaka-10antz/behaviortree/behaviortree"
	"github.com/takahiko-tanaka-10antz/behaviortree/models"
)

type UserLogic struct {
	User models.User
}

func NewUserLogic(hp int, sa, sb float64) *UserLogic {
	return &UserLogic{
		User: *models.NewUser(hp, sa, sb),
	}
}

func (u *UserLogic) IsHP100OrMore() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			if u.User.HP >= 100 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}

func (u *UserLogic) GetCloser() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Println("敵に寄る")
			return Success, nil
		},
	)
}

func (u *UserLogic) CallAFriends(f string) Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Printf("フレンド%sを呼ぶ\n", f)
			return Success, nil
		},
	)
}

func (u *UserLogic) ActivateSkills(e *models.Enemy, typ string) Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			switch typ {
			case "A":
				e.HP -= u.User.SkillA.Damage
			case "B":
				e.HP -= u.User.SkillB.Damage
			}

			fmt.Printf("スキル%s発動\n", typ)
			return Success, nil
		},
	)
}

func (u *UserLogic) GetSkills(typ string) Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			rand.Seed(time.Now().UnixNano())
			r := rand.Float64()

			var sr float64
			switch typ {
			case "A":
				sr = u.User.SkillA.Rate
			case "B":
				sr = u.User.SkillB.Rate
			}

			if r > sr {
				return Failure, nil
			}
			return Success, nil
		},
	)
}
