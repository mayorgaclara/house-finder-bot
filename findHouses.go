package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FindHouses() string {
	res, err := http.Get("https://api.mercadolibre.com/sites/MLA/search?item_location=lat:-37.122878_-37.066198,lon:-56.885710_-56.827878&category=MLA1468&since=today")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	if result.Paging.Total == 0 {
		return "No hay casas nuevas publicadas hoy :("
	}

	t := time.Now()
	hoy := fmt.Sprintf("%02d-%02d-%02d",
		t.Day(), t.Month(), t.Year())

	response := "Estas son las casas encontradas para vos el día de hoy! " + hoy + "\n" + "\n"
	for _, rec := range result.Results {
		if rec.Price <= 170000 {
			response += "🏠 CASA: " + rec.Permalink + "-> Precio: USD " + strconv.Itoa(rec.Price) + "\n" + "\n"
		}
	}
	if len(response) > 100 {
		return response
	} else {
		return "No encontré casas publicadas hoy a ese valor :("
	}
}
