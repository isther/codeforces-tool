package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contest struct {
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int64  `json:"durationSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
}

type Response struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

var (
	baseUrl = "http://codeforces.com/api/contest.list"
)

func GetContest() []*Contest {
	contest, err := getContest(baseUrl + "?contest=true")
	if err != nil {
		return nil
	}
	return contest
}

func getContest(url string) ([]*Contest, error) {
	client := &http.Client{}
	// req, _ := http.NewRequest("GET", url, nil)
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return nil, fmt.Errorf("Status error")
	}
	defer resp.Body.Close()

	contestList, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Response
	_ = json.Unmarshal([]byte(contestList), &r)

	cnt := -1
	for i, v := range r.Result {
		if v.Phase == "FINISHED" {
			break
		}
		cnt = i
	}
	if cnt == -1 {
		return nil, fmt.Errorf("Don't have contest")
	}

	var contestlist []*Contest

	for i := 0; i < 3 && cnt >= 0; i++ {
		contestlist = append(contestlist, &r.Result[cnt])
		cnt--
	}

	return contestlist, nil
}
