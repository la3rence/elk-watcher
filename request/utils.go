package request

import (
	"encoding/json"
	dingtalk "github.com/Lonor/dingtalkbot-sdk"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
 封装 HTTP 客户端，仅限 JSON Body 请求
*/
func Request(url string, data string, method string, headers map[string]string) (bodyText string) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// map 判断 key 是否存在
	_, existBasicHeader := headers["username"]
	if existBasicHeader {
		req.SetBasicAuth(headers["username"], headers["password"])
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Printf("[Error] Request [%d] [%s]\n", resp.StatusCode, err)
		return ""
	} else {
		return string(bodyByte)
	}
}

func PostDingTalk(msg string) {
	bot := dingtalk.NewDingBot(os.Getenv("DINGTALK_TOKEN"), os.Getenv("DINGTALK_SECRET"))
	_ = bot.SendSimpleText(msg)
}

func GetLogCount() (value int) {
	elasticURL := os.Getenv("ELASTIC_URL")
	headers := map[string]string{
		"username": os.Getenv("ELASTIC_USERNAME"),
		"password": os.Getenv("ELASTIC_PASSWORD"),
	}
	// ES API JSON data: 这里尝试请求距离现在 15 秒以内的日志
	data := `{
	"query": {
    	"bool": {
      		"filter": {
        		"range": {
          			"@timestamp": {
            			"gt": "now-15s"
						}
        			}
				}
			}
		}
	}`
	responseString := Request(elasticURL, data, "GET", headers)
	if responseString == "" {
		return -1
	}
	esResponse := ElasticSearchResponse{}
	elasticSearchResponse := parseJSON(responseString, esResponse)
	return elasticSearchResponse.Hits.Total.Value
}

func parseJSON(text string, class ElasticSearchResponse) ElasticSearchResponse {
	_ = json.Unmarshal([]byte(text), &class)
	return class
}

type ElasticSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore interface{}   `json:"max_score"`
		Hits     []interface{} `json:"hits"`
	} `json:"hits"`
}
