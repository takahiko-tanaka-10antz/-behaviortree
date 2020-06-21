package behaviortree

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
