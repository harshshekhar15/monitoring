package types

type Server struct {
	IP                   string `json:"ip"`
	PercentageCPUUsed    int    `json:"percentage_cpu_used"`
	PercentageMemoryUsed int    `json:"percentage_memory_used"`
}

type Usage struct {
	IP        string `json:"ip"`
	MaxCPU    int    `json:"max_cpu"`
	MaxMemory int    `json:"max_memory"`
}
