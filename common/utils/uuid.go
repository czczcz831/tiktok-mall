package utils

import (
	"github.com/bwmarrin/snowflake"
)

func UUIDGenerate(nodeID int64) (string, error) {
	node, err := snowflake.NewNode(nodeID % 1024)
	if err != nil {
		return "", err
	}
	uuid := node.Generate().String()
	return uuid, nil
}
