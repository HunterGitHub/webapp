package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"net/http"
	"time"
	"log"
	"fmt"
//	"strings"
)

func mainA() {
	dbmap := initDb()

    k := martini.Classic()
    k.Use(render.Renderer(render.Options{
		Directory: tplDir,
		Delims: render.Delims{"{[{", "}]}"},
    }))
    k.Get("/", func(r render.Render) {
		r.HTML(200, "angularIndex", nil)
	})
	k.Use(martini.Static(pubDir))
	log.Fatal(http.ListenAndServe(":80", k))


    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
		Directory: tplDir,
    }))

    m.Get("/default", func(r render.Render) {
        r.HTML(200, "pesys", nil)
	})

    m.Get("/companyPage", func(r render.Render) {
		var c []Company
		_, err := dbmap.Select(&c, selectAll)
		checkErr(err, "select company err");  //fmt.Println("com ---- ", c)
		amap := map[string]interface{}{"company": c}
        r.HTML(200, "companyPage", amap)
	})

    m.Get("/modifyPage/:id", func(args martini.Params, r render.Render) {
		var c Company
		err := dbmap.SelectOne(&c, selectOne, args["id"])
		checkErr(err, "select one company err")
		amap := map[string]interface{}{"company": c}
		r.HTML(200, "modifyPage", amap)
	})

    m.Get("/addPage", func(args martini.Params, r render.Render) {
		r.HTML(200, "modifyPage", nil)
	})

    m.Get("/delPage", func(args martini.Params, r render.Render) {
		r.HTML(200, "delPage", nil)
//		var param = ids.Dataids;  fmt.Println("json ---- ", param)
//		arr := strings.Split(param, "@")
//		amap := map[string]interface{}{"ids":arr};  //fmt.Println(msg)
	})

    m.Post("/del", binding.Bind(DelForm{}), func(df DelForm, args martini.Params, r render.Render) {
		var param = df.Dataids; fmt.Println("del json ---- ", param)
		r.Status(200)
	})

    m.Post("/add", binding.Bind(Company{}), func(c Company, args martini.Params, r render.Render) {
		c.Time = time.Now().String()
		c.Operator = "a sb"
		fmt.Println(c)
		err := dbmap.Insert(&c)
		checkErr(err, "insert one company err")
		r.Status(200)
	})

    m.Post("/modify", binding.Bind(Company{}), func(c Company, args martini.Params, r render.Render) {
		c.Operator = "another sb"
		fmt.Println(c)
		_, err := dbmap.Db.Exec(updateOne, c.Name, c.Person, c.Phone, c.Operator, c.Usedomain, c.Info, c.Id)
		checkErr(err, "update one company err")
		r.Status(200)
	})

    m.Get("/", func(r render.Render) {
		r.HTML(200, "angularIndex", nil)
	})
	// test part ------------------------
    m.Get("/index", func(r render.Render) {
        tplVar := map[string]interface{}{"metatitle": "this is my custom title", "start": "index"}
        r.HTML(200, "index", tplVar)
	})

    m.Get("/show", func(r render.Render) {
        tplVar := map[string]interface{}{"metatitle": "this is my custom title", "show": "show"}
        r.HTML(200, "show", tplVar)
	})

    m.Get("/all", func(r render.Render) {
		tplVar := []map[string]interface{}{{"metatitle": "this is my custom title", "start": "sssshit"}, {"1123":"fddd"}}
        r.JSON(200, tplVar)
	})

	m.Use(martini.Static(pubDir))

	log.Fatal(http.ListenAndServe(":8080", m))
//	run.Init()
	dbmap.Db.Close()
}

type DelForm struct {
	Dataids		string	`form:"dataids"`
}

func checkErr(err error, msg string) {
    if err != nil {
        fmt.Println(msg, err)
    }
}
