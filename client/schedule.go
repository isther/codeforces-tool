package client

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ID struct {
	string
}

type IDList struct {
	ID []ID
}

type ContestInfo struct {
	ID          string
	Name        string
	Start       string
	Length      string
	BeforeStart string
	Link        string
}

var (
	baseCodeforcesContestsUrl = "http://codeforces.com/contests/"
	listLimit                 = 3
)

func (list IDList) GetContestList() (contestList []*ContestInfo) {
	list.getContestsID()
	// fmt.Println(list)

	for _, v := range list.ID {
		contestList = append(contestList, v.getContestInfoByID())
	}
	return contestList
}

func (list *IDList) getContestsID() {
	url := baseCodeforcesContestsUrl

	getDocument(url).Find("#pageContent .contestList .datatable table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i > listLimit {
			return
		}

		id, _ := s.Attr("data-contestid")
		list.ID = append(list.ID, ID{string: id})
	})
}

func (id ID) getContestInfoByID() (contest *ContestInfo) {
	url := baseCodeforcesContestsUrl + id.string

	getDocument(url).Find("#pageContent .contestList .datatable table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 || i > 1 {
			return
		}

		text := s.Find("td").Text()
		temp := strings.Split(text, "\n")

		pos := 0
		flag := 0
		cnt := 0
		info := []string{}
		for i, v := range temp {
			v = strings.TrimSpace(v)
			if v != "" {
				info = append(info, v)
				cnt++
			}

			if len(v) == len("Dec/27/2021 17:35") && flag == 0 && i != 0 {
				pos = cnt - 1
				flag = 1
			}
		}

		contest = &ContestInfo{
			ID:          id.string,
			Name:        info[0],
			Start:       info[pos],
			Length:      info[pos+1],
			BeforeStart: info[pos+2],
			Link:        baseCodeforcesContestsUrl + id.string,
		}
		// fmt.Println(contest)
	})
	return contest
}

func getDocument(url string) *goquery.Document {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
