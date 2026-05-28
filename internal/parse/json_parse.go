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

// BuildTree рекурсивно заполняет узел детьми и значениями.
func BuildTree(node *JSONNode, val any, depth int) {
    switch v := val.(type) {
    case map[string]any:
        for k, elem := range v {
            child := &JSONNode{Key: k, Depth: depth + 1, Expanded: false}
            BuildTree(child, elem, depth+1)
            node.Children = append(node.Children, child)
        }
    case []any:
        for i, elem := range v {
            child := &JSONNode{Key: fmt.Sprintf("[%d]", i), Depth: depth + 1, Expanded: false}
            BuildTree(child, elem, depth+1)
            node.Children = append(node.Children, child)
        }
    default:
        node.Value = v
    }
}
