package GroopieTracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/res"
	_var "main/res/var"
	"net/http"
	"os"
)

func InitApi() {
	a, err := http.Get(res.UrlApiArtists)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(a.Body)

	artisteData, err := io.ReadAll(a.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(artisteData, &_var.AData)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	r, err := http.Get(res.UrlApiRelation)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	relationData, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(relationData, &_var.RData)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

}
