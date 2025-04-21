package storage

type Tool struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func UpdateTool(id string, tool Tool) error {
	// 这里实现更新逻辑
	return nil
}