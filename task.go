package robot

import (
	"net/http"
	"strconv"
)

const (
	almHardLimit = "ALM-HARD-LIMIT"
	almSoftLimit = "ALM-SOFT-LIMIT"
)

type Task struct {
	HardLimit int
	SoftLimit int
	GraphID   string
	In        Elements
	Out       Elements
}

func NewTask(req http.Request) (Task, error) {
	s := req.Header.Get("ALM-HARD-LIMIT")
	hardLimit, _ := strconv.ParseInt(s, 10, 64)
	s = req.Header.Get("ALM-SOFT-LIMIT")
	softLimit, _ := strconv.ParseInt(s, 10, 64)
	graphID := req.Header.Get("ALM-GRAPH-ID")
	t := Task{
		HardLimit: int(hardLimit),
		SoftLimit: int(softLimit),
		GraphID:   graphID,
	}
	return t, nil
}
