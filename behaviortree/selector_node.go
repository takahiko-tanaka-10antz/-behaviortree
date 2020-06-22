package behaviortree

// childrenの処理がFailureである限り処理し続ける
func SelectorNode(children []Node) (Status, error) {
	for _, c := range children {
		status, err := c.Tick()
		if err != nil {
			return Failure, err
		}
		if status == Success {
			return Success, nil
		}
	}

	return Failure, nil
}
