package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
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
		fmt.Println("No hay casas nuevas publicadas hoy :(")
		return
	}

	t := time.Now()
	hoy := fmt.Sprintf("%02d-%02d-%02d",
		t.Day(), t.Month(), t.Year())

	fmt.Println("Estas son las casas encontradas para vos el día de hoy!", hoy)
	for _, rec := range result.Results {
		fmt.Println("Link = ", rec.Permalink, "-> Price = USD ", rec.Price)
	}
}
