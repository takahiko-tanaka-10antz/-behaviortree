package behaviortree

import (
	"errors"
)

// 全てのノードの親
func RootNode(node Node) error {
	tick, children := node()
	if tick == nil {
		return errors.New("behaviortree.RootNode :: cannot tick a nil node")
	}

	status, err := tick(children)
	if err == nil && status == Failure {
		return errors.New("behaviortree.RootNode :: If it fails, end the process")
	}

	return nil
}
