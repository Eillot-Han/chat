package websocket
//进程池
type Process struct {
	Pools map[string]*Pool
}

func NewProcess() *Process {
	return &Process{
		Pools: make(map[string]*Pool),
	}
}
