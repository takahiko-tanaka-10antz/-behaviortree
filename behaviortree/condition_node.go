package behaviortree

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
