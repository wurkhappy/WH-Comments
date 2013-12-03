package main

import (
	"github.com/ant0ine/go-urlrouter"
	"github.com/wurkhappy/WH-Comments/handlers"
)

//order matters so most general should go towards the bottom
var router urlrouter.Router = urlrouter.Router{
	Routes: []urlrouter.Route{
		urlrouter.Route{
			PathExp: "/agreement/:agreementID/comments",
			Dest: map[string]interface{}{
				"POST": handlers.CreateComment,
				"GET":  handlers.GetComments,
			},
		},
		urlrouter.Route{
			PathExp: "/agreement/:agreementID/tags",
			Dest: map[string]interface{}{
				"GET":  handlers.GetTags,
			},
		},
	},
}
