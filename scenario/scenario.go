package scenario

import (
	"fmt"

	. "github.com/takahiko-tanaka-10antz/behaviortree/behaviortree"
	"github.com/takahiko-tanaka-10antz/behaviortree/models"
)

func Start() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Println("出発")

			return Success, nil
		},
	)
}

func End(enemy *models.Enemy) Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			switch {
			case enemy.HP <= 0:
				fmt.Println("End1")
			case enemy.HP > 0:
				fmt.Println("End2")
			}

			return Success, nil
		},
	)
}
