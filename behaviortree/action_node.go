package behaviortree

import "errors"

// 実行処理
func ActionNode(children []Node) (Status, error) {
	if len(children) == 0 {
		return Failure, errors.New("behaviortree.ActionNode :: child node does not exist")
	}

	status, err := children[0].Tick()
	if err != nil {
		return Failure, err
	}
	return status, nil
}
