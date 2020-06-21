package logic

import (
	. "github.com/takahiko-tanaka-10antz/behaviortree/behaviortree"
	"github.com/takahiko-tanaka-10antz/behaviortree/models"
)

type EnemyLogic struct {
	Enemy models.Enemy
}

func NewEnemyLogic(hp int) *EnemyLogic {
	return &EnemyLogic{
		Enemy: *models.NewEnemy(hp),
	}
}

func (e *EnemyLogic) Dead() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			if e.Enemy.HP <= 0 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}

func (e *EnemyLogic) Alive() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			if e.Enemy.HP >= 0 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}
