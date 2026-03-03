package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

type Server struct {
	ID        int
	Hostname  string
	IsOffline bool
	CPUUsage  float64
	RAMUsage  float64
}

func main() {
	fmt.Println("Hello, World!")

	servers := []Server{
		Server{
			ID:       1290412,
			Hostname: "db-master",
			CPUUsage: 50.0,
			RAMUsage: 54.2,
		},
		Server{
			ID:       9523942,
			Hostname: "db-worker",
			CPUUsage: 85.0,
			RAMUsage: 66.2,
		},
		Server{
			ID:       9304318,
			Hostname: "control-plane",
			CPUUsage: 20.0,
			RAMUsage: 15.5,
		},
		Server{
			ID:       1485923,
			Hostname: "worker-node-1",
			CPUUsage: 78.0,
			RAMUsage: 89.0,
		},
		Server{
			ID:       6758342,
			Hostname: "worker-node-2",
			CPUUsage: 79.0,
			RAMUsage: 82.5,
		},
	}

	ResetOverLoad(servers)

	fmt.Println("Array after fix:")
	pp.Println(servers)

	fmt.Println("Add new server:")
	AddNewServer(&servers)

	pp.Println(servers)

}

func ResetOverLoad(list []Server) {
	for index := range list {
		if list[index].CPUUsage > 80 || list[index].RAMUsage > 85 {
			fmt.Printf("server number: %v CPU is overloaded, starting reload....\n", index)
			time.Sleep(1 * time.Second)
			list[index].CPUUsage = 0
			list[index].RAMUsage = 0
			fmt.Println("Server was rebooted.")
			pp.Println(list[index])
		}
	}
}

func AddNewServer(list *[]Server) {
	var ID int
	var Hostname string
	fmt.Println("Adding new server...")
	fmt.Print("Type name for new server: ")
	fmt.Scan(&Hostname)
	fmt.Print("Type server ID: ")
	fmt.Scan(&ID)
	*list = append(*list, Server{
		ID:       ID,
		Hostname: Hostname,
	})
}
