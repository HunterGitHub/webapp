package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "gopkg.in/mgo.v2/bson"
)

func main() {
    dbmap := initDb()
    mgodb := initMgo()
    defer dbmap.Db.Close()
    defer mgodb.Close()
    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    auth := r.Group("/", gin.BasicAuth(gin.Accounts{
        "foo":    "baz",
        "austin": "1234",
        "lena":   "hello2",
        "manu":   "4321",
    }))
	fmt.Println(auth)
    r.GET("/js/:jsfile", func(c *gin.Context) {
        file := c.Params.ByName("jsfile")
        fmt.Println(file)
        c.File(pubDir + "/js/" + file)
//      c.String(200, pubDir + "/js/" + file)
    })
    r.GET("/css/:cssfile", func(c *gin.Context) {
        file := c.Params.ByName("cssfile")
        c.File(pubDir + "/css/" + file)
    })
    r.GET("/", func(c *gin.Context) {
        r.LoadHTMLGlob(tplFiles)
        c.HTML(200, "futureMntr.tmpl", nil)
    })
    r.GET("/mntr", func(c *gin.Context) {
        r.LoadHTMLGlob(tplFiles)
        c.HTML(200, "mntr.tmpl", nil)
    })
    r.GET("/re", func(c *gin.Context) {
        r.LoadHTMLGlob(tplFiles)
        c.HTML(200, "resume.tmpl", nil)
    })
    r.GET("/hostOps", func(c *gin.Context) {
        r.LoadHTMLGlob(tplFiles)
        c.HTML(200, "hostOps.tmpl", nil)
	})
    r.GET("/loginPage", func(c *gin.Context) {
        c.HTML(200, "login.tmpl", nil)
    })

    r.GET("/json", func(c *gin.Context) {
        var f []Feedback
        _, err := dbmap.Select(&f, selectAll)
        checkErr(err, "select feedbacks err");  //fmt.Println("com ---- ", c)
        c.JSON(200, f)
    })
    r.POST("/hostInfo", func(c *gin.Context) {
        var hl HostsList
        var his []HostInfo
        if c.Bind(&hl) {
            fmt.Println("the hostlist struct ---- ", hl)
            for _, v := range hl.Param {
                his = append(his, HostInfo{v, "okok", "system"})
            }
            fmt.Println("the hostInfo struct ---- ", his)
        }
        c.JSON(200, his)
    })
    r.POST("/MntrList", func(c *gin.Context) {
        type Param struct { Group string `json:"group"` }
        var g Param
        if c.Bind(&g) { fmt.Println("the group struct ---- ", g) }
        var hosts = list[g.Group]
        var res []HostInfo
        for _, v := range hosts {
            if v == "114.215.137.156" {
                var info = HostMntr(v)
                res = append(res, HostInfo{v, info, "system"})
            } else {
                res = append(res, HostInfo{v, "load 4 cpu 75% mem 60% disk 80%", "system"})
            }
        }
        fmt.Println("the group struct ---- ", res)
        c.JSON(200, res)
    })
    r.POST("/resumeCmt", func(c *gin.Context) {
        type Param struct { Cmt string `json:"cmt" bson:"cmt"` }
        var cmt Param
        if c.Bind(&cmt) { fmt.Println("the cmt is --- ", cmt) }
        coll := mgodb.DB("resume").C("cmt")
//      err := coll.Remove(nil)
        err := coll.Insert(&cmt)
//      err = coll.Find(bson.M{}).One(&cmt)
        if err != nil { panic(err) }
        c.JSON(200, cmt)
    })
    r.GET("/showCmt", func(c *gin.Context) {
        type Obj struct { Cmt string `json:"cmt" bson:"cmt"` }
        var cmt []Obj
        coll := mgodb.DB("resume").C("cmt")
        err := coll.Find(bson.M{}).All(&cmt)
        if err != nil { panic(err) }
        c.JSON(200, cmt)
    })
    r.Run(":80")

//  r.POST("/aaa", func(c *gin.Context) {
//      var aa AA
//      c.Bind(&aa)
//      fmt.Println("aa ---- ", aa)
//      c.JSON(200, "{aaa: 'fdfd'}")
//  })
//  type AA struct { Param   []string `json:"hosts"` }

}

//      c.JSON(200, `[{"host": "` + g.Group + `", "info": "oooh", "option": "system"}]`)
