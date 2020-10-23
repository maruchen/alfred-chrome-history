package main

import (
	"fmt"
	"strconv"
	"strings"

	aw "github.com/deanishe/awgo"
)

type WorkflowOptions struct {
	Profile string `env:"CHROME_PROFILE"`
	Limit   int    `env:"LIMIT"`
}

func runWithAlfred(wf *aw.Workflow) {
	opts := WorkflowOptions{}
	cfg := aw.NewConfig()
	if err := cfg.To(&opts); err != nil {
		panic(err)
	}

	args := wf.Args()
	query := strings.Join(args, " ")
	entries, err := queryHistory(opts.Profile, query, query, opts.Limit)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		title := fmt.Sprintf("%s [%s]", entry.Title, entry.LastVisitTime.Format("1月2日 15点"))
		// item := wf.NewItem(entry.Title)
		item := wf.NewItem(title)
		item.UID(strconv.Itoa(entry.ID))
		item.Subtitle(entry.URL)
		item.Arg(entry.URL)
		item.Valid(true)
	}
	wf.SendFeedback()
}
