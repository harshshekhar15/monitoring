package main

// Add new client to Hosts array
func addHost(newHost string) {
	if len(Hosts) != 0 {
		for _, host := range Hosts {
			if host == newHost {
				return
			}
		}
	}
	Hosts = append(Hosts, newHost)
}
