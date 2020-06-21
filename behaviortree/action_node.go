package behaviortree

// 実行処理
func ActionNode(children []Node) (Status, error) {
	status, err := children[0].Tick()
	if err != nil {
		return Failure, err
	}
	return status, nil
}
