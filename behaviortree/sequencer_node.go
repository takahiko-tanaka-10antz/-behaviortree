package behaviortree

// childrenの処理がSuccessである限り処理し続ける
func SequencerNode(children []Node) (Status, error) {
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
