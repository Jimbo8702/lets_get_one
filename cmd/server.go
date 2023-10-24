package main

import (
	"log"
	"os"

	"github.com/Jimbo8702/lets_get_one/api"
	"github.com/Jimbo8702/lets_get_one/db"
	"github.com/Jimbo8702/lets_get_one/render"
	"github.com/Jimbo8702/lets_get_one/session"
	"github.com/Jimbo8702/lets_get_one/util"
	"github.com/Jimbo8702/lets_get_one/validator"
	"github.com/sagikazarmark/slog-shim"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	dsn := util.BuildDSN(config)
	db, err := db.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	var (
		vd 		= validator.NewPlaygroundValidator("params")
		ls 		= slog.New(slog.NewTextHandler(os.Stderr, nil))
		sess 	= session.New(config)
		render  = render.New(sess, config)
		app 	= api.New(db, ls, vd, sess, render)
	)


 	app.Init()

	app.ListenAndServe(config.Port)
}

//TODO: add a new session constructor to inject into api
// add renderer contrusctor to inject into api

