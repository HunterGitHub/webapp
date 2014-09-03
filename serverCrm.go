package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
//	"github.com/martini-contrib/auth"
	"net/http"
	"log"
//	"encoding/json"
	"fmt"
	"time"
//	"strings"
)

func mainCrm() {
	dbmap := initDb()

    k := martini.Classic()

//	k.Use(auth.Basic("admin", "1234"))

    k.Use(render.Renderer(render.Options{
		Directory: tplDir + "/html",
		Delims: render.Delims{"{[{", "}]}"},
    }))

    k.Get("/", func(r render.Render) {
		r.HTML(200, "crmFeedback", nil)
	})

	k.Get("/loginPage", func(r render.Render) {
		r.HTML(200, "login", nil)
	})

	k.Post("/login", binding.Bind(UserForm{}), func(u UserForm, r render.Render) {
		r.JSON(200, u)
	})

    k.Get("/json", func(r render.Render) {
		var f []Feedback
		_, err := dbmap.Select(&f, selectAll)
		checkErr(err, "select feedbacks err");  //fmt.Println("com ---- ", c)
//		b, _ := json.Marshal(f); fmt.Println("------- f s " , f, "----- f json", string(b))
		r.JSON(200, f)
	})

	k.Post("/savef", binding.Bind(FeedbackForm{}), func(f FeedbackForm, r render.Render) {
		fmt.Println("before --- ", f)
		if f.Id == 0 { 
			F := Feedback{f, ""}
			F.Time = time.Now().Format("2006-01-02 15:04:05")
			err := dbmap.Insert(&F); checkErr(err, "insert f err") 
			r.JSON(200, F)
		} else { 
			_, err := dbmap.Update(&f); checkErr(err, "Update f err") 
			r.Status(200)
		}
		fmt.Println("after  --- ", f)
	})

	k.Post("/delf", binding.Bind(Feedback{}), func(f Feedback, r render.Render) {
		fmt.Println(f)
		_, err := dbmap.Delete(&f); checkErr(err, "del f err")
		r.Status(200)
	})


	k.Use(martini.Static(pubDir))

	log.Fatal(http.ListenAndServe(":80", k))

//	run.Init()
	dbmap.Db.Close()
}

