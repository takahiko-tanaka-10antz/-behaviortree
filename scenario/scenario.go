package scenario

import (
	"fmt"

	. "github.com/takahiko-tanaka-10antz/behaviortree/behaviortree"
)

func Start() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Println("出発")
			return Success, nil
		},
	)
}

func END1() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Println("End1")
			return Success, nil
		},
	)
}

func END2() Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			fmt.Println("End2")
			return Success, nil
		},
	)
}
