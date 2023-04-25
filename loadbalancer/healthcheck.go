package loadbalancer

import (
	"fmt"

	"github.com/robfig/cron"
)

func startHealthCheck() {
	c := cron.New()

	for _, eachServer := range serverList {
		func(serverToCheck *server) {
			c.AddFunc("@every 5s", func() {
				fmt.Printf("Checking health: %s\n", serverToCheck.name)
				serverToCheck.checkHealth()
			})
		}(eachServer)
	}

	c.Start()
}
