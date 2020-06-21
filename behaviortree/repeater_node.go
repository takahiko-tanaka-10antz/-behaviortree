package behaviortree

import "errors"

// 指定した回数、子供の処理をする
func RepeaterNode(cnt int, node Node) Node {
	return NewBranch(
		func(children []Node) (Status, error) {
			for i := 0; i < cnt; i++ {
				tick, children := node()
				if tick == nil {
					return Failure, errors.New("behaviortree.RepeaterNode :: cannot tick a nil node")
				}

				status, err := tick(children)
				if err == nil && status == Failure {
					return Failure, nil
				}
			}
			return Success, nil
		},
	)
}
