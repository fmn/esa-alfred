package main

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/pascalw/go-alfred"
	"github.com/yuichiro-h/go-esa"

	"os"
)

func searchCommand(c *cli.Context) {
	var query = strings.Join(c.Args(), " ")

	client := newEsaClient()
	teamName := os.Getenv(envNameEsaDefaultTeamName)

	response := alfred.NewResponse()
	res, err := client.GetTeamPosts(teamName, &esa.GetTeamPostsRequest{
		Q: esa.String(query),
	})

	if err != nil {
		response.AddItem(&alfred.AlfredResponseItem{
			Valid: true,
			Uid:   "error",
			Title: err.Error(),
			Icon:  "icon.png",
		})
		response.Print()
		return
	}

	posts := res.Posts

	for _, post := range posts {
		title := post.Name

		if post.Category != nil {
			title = *post.Category + "/" + title
		}

		response.AddItem(&alfred.AlfredResponseItem{
			Valid:    true,
			Uid:      post.URL,
			Title:    title,
			Subtitle: post.CreatedBy.ScreenName + " " + post.CreatedAt.Format("2006/01/02 15:04:05"),
			Arg:      post.URL,
			Icon:     "icon.png",
		})
	}

	response.Print()
}
