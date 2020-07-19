package registry

// 服务
type Service struct {
	// 服务名
	Name string `json:"name"`
	// 节点列表
	Nodes []*Node `json:"nodes"`
}

// 单个节点
type Node struct {
	Id   string `json:"id"`
	IP   string `json:"ip"`
	Port int    `json:"port"`
	// 权重，用于加权轮询
	Weight int `json:"weight"`
}
