package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/OzkrOssa/mikrotik-go"
	"github.com/OzkrOssa/rss/config"
	"github.com/OzkrOssa/rss/services"
)

func main() {
	// c := cron.New()
	// c.AddFunc("*/5 7-18 * * 1-6", start)
	// c.Start()

	// select {}
	start()
}
func start() {
	t := time.Now()
	var wg sync.WaitGroup

	attributes := []string{
		"ACTIVO",
		"CORTADO",
	}

	mikrotikHost, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	apiUsers, err := services.FetchUsersFromAPI(attributes)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Println(err)
	}
	for _, host := range mikrotikHost {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()

			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered from panic:", r)
				}
			}()
			log.Println("Connecting with: ", h)
			repo, err := mikrotik.NewMikrotikRepository(h, os.Getenv("MIKROTIK_API_USER"), os.Getenv("MIKROTIK_API_PASSWORD"))
			if err != nil {
				log.Println(err, h)
				return
			}

			secrets, err := repo.GetSecrets("", h)
			if err != nil {
				log.Println(err)
			}

			addressList := repo.GetAddressList("Moroso")
			for _, u := range apiUsers {
				for _, s := range secrets {
					comment := fmt.Sprintf("%s_%s", u.Cedula, u.NroContrato)
					addressListComment := fmt.Sprintf("%s %s_%s", u.Nombre, u.Apellido, u.NroContrato)
					if s["comment"] == comment && u.StatusContrato == "ACTIVO" {
						repo.RemoveSecretFromAddressList(addressList, s["remote-address"])
					} else if s["comment"] == comment && u.StatusContrato == "CORTADO" {
						repo.AddSecretToAddressList(s["remote-address"], addressListComment, "Moroso")
					}
				}
			}
		}(host)
	}
	wg.Wait()
	e := time.Since(t)
	fmt.Println(e)
}
