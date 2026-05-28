package parse

import (
	"encoding/json"
	"fmt"
)

type JSONNode struct {
	Key      string
	Value    any
	Depth    int
	Expanded bool
	Children []*JSONNode
}

// парсим JSON-данные и возвращаем корень дерева
func ParseJSON(data []byte) (*JSONNode, error) {
    var raw any
    if err := json.Unmarshal(data, &raw); err != nil {
        return nil, err
    }
    root := &JSONNode{Key: "root", Depth: 0, Expanded: true}
    BuildTree(root, raw, 0)
    return root, nil
}
