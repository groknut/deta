package parse

import (
	"encoding/json"
	"fmt"
	"strings"
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

// FlattenTree возвращает линейный список видимых узлов (с учётом Expanded).
func FlattenTree(root *JSONNode) []*JSONNode {
    var list []*JSONNode
    var walk func(n *JSONNode)
    walk = func(n *JSONNode) {
        list = append(list, n)
        if n.Expanded {
            for _, child := range n.Children {
                walk(child)
            }
        }
    }
    walk(root)
    return list
}

// FormatJSONNode возвращает строковое представление узла для отображения.
func FormatJSONNode(node *JSONNode) string {
    indent := strings.Repeat("  ", node.Depth)
    marker := "  "
    if len(node.Children) > 0 {
        if node.Expanded {
            marker = "▼ "
        } else {
            marker = "▶ "
        }
    }
    keyPart := indent + marker + node.Key
    if node.Value != nil {
        return keyPart + fmt.Sprintf(": %v", node.Value)
    }
    return keyPart
}
