package client

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// ProblemTypes problem types
var ProblemTypes = [...]string{
	"contest",
}

// Info information
type Info struct {
	ProblemType  string `json:"problem_type"`
	ContestID    string `json:"contest_id"`
	GroupID      string `json:"group_id"`
	ProblemID    string `json:"problem_id"`
	SubmissionID string `json:"submission_id"`
	RootPath     string
}

// ErrorNeedProblemID error
const ErrorNeedProblemID = "You have to specify the Problem ID"

// ErrorNeedContestID error
const ErrorNeedContestID = "You have to specify the Contest ID"

// ErrorNeedSubmissionID error
const ErrorNeedSubmissionID = "You have to specify the Submission ID"

// ErrorUnknownType error
const ErrorUnknownType = "Unknown Type"


func (info *Info) errorContest() (string, error) {
	return "", errors.New(ErrorNeedContestID)
}

// Hint hint text
func (info *Info) Hint() string {
	text := strings.ToUpper(info.ProblemType)
	if info.GroupID != "" {
		text = text + " " + info.GroupID
	}
	if info.ProblemType != "acmsguru" && info.ContestID != "" {
		if info.ProblemType != "group" {
			text = text + " " + info.ContestID
		} else {
			text = text + ", contest " + info.ContestID
		}
	}
	if info.ProblemID != "" {
		text = text + ", problem " + info.ProblemID
	}
	if info.SubmissionID != "" {
		text = text + ", submission " + info.SubmissionID
	}
	return text
}

// Path path
func (info *Info) Path() string {
	path := info.RootPath
	if info.GroupID != "" {
		path = filepath.Join(path, info.GroupID)
	}
	if info.ProblemType != "acmsguru" && info.ContestID != "" {
		path = filepath.Join(path, info.ContestID)
	}
	if info.ProblemID != "" {
		path = filepath.Join(path, strings.ToLower(info.ProblemID))
	}
	return path
}

/*******************************
********************************
		Generate link
********************************
********************************/

// ProblemSetURL parse problem set url
func (info *Info) ProblemSetURL(host string) (string, error) {
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v", info.ContestID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// ProblemURL parse problem url
func (info *Info) ProblemURL(host string) (string, error) {
	if info.ProblemID == "" {
		return "", errors.New(ErrorNeedProblemID)
	}
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v/problem/%v", info.ContestID, info.ProblemID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// MySubmissionURL parse submission url
func (info *Info) MySubmissionURL(host string) (string, error) {
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v/my", info.ContestID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// SubmissionURL parse submission url
func (info *Info) SubmissionURL(host string) (string, error) {
	if info.SubmissionID == "" {
		return "", errors.New(ErrorNeedSubmissionID)
	}
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v/submission/%v", info.ContestID, info.SubmissionID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// StandingsURL parse standings url
func (info *Info) StandingsURL(host string) (string, error) {
	if info.ContestID == "" {
		return info.errorContest()
	}
	switch info.ProblemType {
	case "contest":
		return fmt.Sprintf(host+"/contest/%v/standings", info.ContestID), nil
	}
	return "", errors.New(ErrorUnknownType)
}

// SubmitURL submit url
func (info *Info) SubmitURL(host string) (string, error) {
	URL, err := info.ProblemSetURL(host)
	if err != nil {
		return "", err
	}
	return URL + "/submit", nil
}

// OpenURL open url
func (info *Info) OpenURL(host string) (string, error) {
	switch info.ProblemType {
	case "contest":
		if info.ContestID == "" {
			return host + "/contests", nil
		} else if info.ProblemID == "" {
			return fmt.Sprintf(host+"/contest/%v", info.ContestID), nil
		}
		return fmt.Sprintf(host+"/contest/%v/problem/%v", info.ContestID, info.ProblemID), nil
	}
	return "", errors.New("Hmmm I don't know what you want to do~")
}
