package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/OzkrOssa/rss/models"
)

func FetchUsersFromAPI(clientStatus []string) ([]*models.Client, error) {
	log.Println("Fetching API.....")
	var wg sync.WaitGroup
	var saeplusUsers []*models.Client

	dataChan := make(chan []*models.Client, len(clientStatus))
	errorChan := make(chan error, len(clientStatus))

	for _, attr := range clientStatus {
		wg.Add(1)
		go func(att string) {
			defer wg.Done()
			req, err := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf(
					"%s?estatus_contrato=%s",
					os.Getenv("API_SAEPLUS_BASE_URL"), att,
				),
				nil,
			)
			if err != nil {
				errorChan <- fmt.Errorf("error al realizar la solicitud: %s", err)
				return
			}

			req.Header.Set(
				"Api-Token",
				os.Getenv("API_SAEPLUS_TOKEN"),
			)
			req.Header.Set("Api-Connect", "red_planet")

			client := http.Client{}
			res, getErr := client.Do(req)
			if getErr != nil {
				errorChan <- fmt.Errorf("error al realizar la solicitud HTTP: %s", getErr)
				return
			}

			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				errorChan <- fmt.Errorf("error al leer la respuesta HTTP: %s", err)
				return
			}

			var response models.Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				errorChan <- fmt.Errorf("error al deserializar el JSON: %s", err)
				return
			}
			dataChan <- response.Data.Info
		}(attr)
	}

	go func() {
		wg.Wait()
		close(dataChan)
		close(errorChan)
	}()

	for data := range dataChan {
		// Procesar los datos recibidos del canal dataChan
		for i := range data {
			abonado, err := strconv.Atoi(data[i].NroContrato)
			if err != nil {
				fmt.Println("Error al convertir a entero:", err)
			}
			data[i].NroContrato = strconv.Itoa(abonado)
		}
		saeplusUsers = append(saeplusUsers, data...)
	}

	// Manejar los errores recibidos del canal errorChan
	for err := range errorChan {
		return nil, err
	}
	return saeplusUsers, nil
}
