package utils

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node     *snowflake.Node
	nodeOnce sync.Once
)

// UUIDGenerate generates a unique snowflake ID
func UUIDGenerate(nodeID int64) (string, error) {
	var err error
	nodeOnce.Do(func() {
		node, err = snowflake.NewNode(nodeID % 1024)
	})
	if err != nil {
		return "", err
	}
	uuid := node.Generate().String()
	return uuid, nil
}
