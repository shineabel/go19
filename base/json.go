package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type (



	gResult struct {
		Title string `json:"title"`
		TitleNoFormating string `json:"titleNoFormating"`
		Content string `json:"content"`
		cacheURL string `json:"cacheUrl"`
		visibleURL string `json:"visibleUrl"`
		URL string `json:"url"`
		GSearchResultClass string `json:"GsearchResultClass"`
		unescapedURL string `json:"unescapedUrl"`

	}

	gResponse struct {

		ResponseData struct{
			Results []gResult `json:"results"`

		} `json:"responseData"`
	}
)

func main() {

	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)
	if(err != nil){
		fmt.Println("query failed:",err)
	}
	defer resp.Body.Close()
	var gr gResponse


	err = json.NewDecoder(resp.Body).Decode(&gr)
	if(err != nil){
		fmt.Println("error")
		return
	}
	fmt.Println(gr)

}
