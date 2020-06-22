package behaviortree

// 並列実行した子供が全て成功したらSuccessを返す
func ParallelNode(children []Node) (Status, error) {
	status := make(chan Status, len(children))
	err := make(chan error, len(children))
	done := make(chan interface{})

	// SequencerNodeを並列に実行する
	go func() {
		s, e := SequencerNode(children)
		if e != nil {
			err <- e
			status <- Failure
		}
		if s == Failure {
			err <- nil
			status <- Failure
		}

		err <- nil
		status <- Success

		close(status)
		close(err)
		close(done)
	}()

	for {
		select {
		case err := <-err:
			if err != nil {
				return Failure, err
			}
		case status := <-status:
			if status == Failure {
				return Failure, nil
			}
		case <-done:
			return Success, nil
		}
	}
}
