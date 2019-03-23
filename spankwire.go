package spankwire

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "http://www.spankwire.com/api/HubTrafficApiCall"
const APITimeout = 5

func GetVideoByID(ID string) SpankwireSingleVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"?data=getVideoById&output=json&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SpankwireSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result

}

func GetVideoEmbedCode(ID string) SpankwireEmbedCode {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"?data=getVideoEmbedCode&output=json&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SpankwireEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
