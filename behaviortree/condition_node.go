package behaviortree

// 条件との合致を評価する関数が True を戻した場合、現在アクティブな子供を呼び出す
func ConditionNode(children []Node) (Status, error) {
	for _, c := range children {
		status, err := c.Tick()
		if err != nil {
			return Failure, err
		}
		if status == Failure {
			return Failure, nil
		}
	}

	return Success, nil
}
