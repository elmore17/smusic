package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://spotify23.p.rapidapi.com/tracks/?ids=4WNcduiCmDNfmTEz7JvmLv"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "b41ae80150mshb5f1bd37a09df29p1d6e78jsn625e72274f71")
	req.Header.Add("X-RapidAPI-Host", "spotify23.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}