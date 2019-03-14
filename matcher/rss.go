package matcher

import (
	"encoding/xml"
	"github.com/go19/search"
	"log"
	"regexp"
	"errors"
	"net/http"
	"fmt"
)

type (

	item struct {
		XMLName xml.Name `xml:"item"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		GUId string `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}


	image struct {
		XMLName xml.Name `xml:"image"`
		URL string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		PubDate string `xml:"pubDate"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image image	`xml:"image"`
		Item []item `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}

)

type rssMatcher struct{}

func (matcher rssMatcher)Search(feed *search.Feed,searchTerm string) ([] *search.Result,error) {

	var results []*search.Result
	log.Printf(feed.Type,feed.Name,feed.URI)

	document, err := matcher.Retrieve(feed)
	if(err != nil){
		return nil,err
	}
	for _,channelItem := range  document.Channel.Item {


		matched, err := regexp.MatchString(searchTerm,channelItem.Title)
		if(err != nil){
			return nil,err
		}
		if(matched){
			results = append(results, &search.Result{
				Filed:"Title",
				Content:channelItem.Title,
			})
		}
		matched,err = regexp.MatchString(searchTerm,channelItem.Description)
		if(err != nil){
			return nil,err
		}
		if(matched){
			results = append(results,&search.Result{
				Filed:"Description",
				Content:channelItem.Description,
			})
		}
	}
	return results,nil
}

func init()  {
	var rssm rssMatcher
	search.Register("rss",rssm)
}

func (match rssMatcher)Retrieve(feed *search.Feed)(*rssDocument,error)  {

	if(feed.URI == ""){
		return nil,errors.New("feed uri not provided")
	}

	resp, err := http.Get(feed.URI)
	if(err != nil){
		return nil,err
	}
	defer resp.Body.Close()
	if(resp.StatusCode != 200){
		return nil, fmt.Errorf("http status code:",resp.StatusCode)
	}

	var rssDoc rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&rssDoc)

	return &rssDoc,err
}

