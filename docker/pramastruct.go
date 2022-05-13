package docker

type CreateStruct struct {
	Hostname   string `json:"hostname" binding:"required""`
	CPUShares  int64  `json:"cpushares" binding:"required"`
	CpusetCpus string `json:"cpusetcpus" binding:"required"`
	NanoCPUs   int64  `json:"nanocpus" binding:"required"`
	Memory     int64  `json:"memory" binding:"required"`
	MemorySwap int64  `json:"memoryswap" binding:"required"`
	MinNet     string `json:"min_net" binding:"required"`
	MaxNet     string `json:"max_net" binding:"required"`
	ProtNums   int    `json:"prot_nums" binding:"required"`
	DomainNums int    `json:"domain_nums" binding:"required"`
	Image      string `json:"image" binding:"required"`
	Ipprefix   string `json:"Ipprefix" binding:"required"`
}
type InfoStruct struct {
	//Id string `json:"id" binding:"required"`
	Id string `json:"id"`
}
