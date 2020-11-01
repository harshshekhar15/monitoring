package main

// getMaxCPU returns max CPU usage of a client
func getMaxCPU(host string) int {
	maxCPU := 0
	for _, svr := range Servers {
		if svr.IP == host {
			if svr.PercentageCPUUsed > maxCPU {
				maxCPU = svr.PercentageCPUUsed
			}
		}
	}
	return maxCPU
}

// getMaxMemory returns max memory usage of a client
func getMaxMemory(host string) int {
	maxMemory := 0
	for _, svr := range Servers {
		if svr.IP == host {
			if svr.PercentageMemoryUsed > maxMemory {
				maxMemory = svr.PercentageMemoryUsed
			}
		}
	}
	return maxMemory
}
