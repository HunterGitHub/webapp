package main

import (
	"os"
)

const (

//	updateOne = "update company set name=?, person=?, phone=?, operator=?, usedomain=?, info=? where id=?"
//	selectAll = "select * from company"
//	selectOne = "select * from company where id=?"
	tableName = "feedbacks"

	updateOne = "update " + tableName + " set name=?, person=?, phone=?, operator=?, usedomain=?, info=? where id=?"
	selectAll = "select * from " + tableName
	selectOne = "select * from " + tableName + " where id=?"
)

//  var root	 = "/root/workspace/src/github.com/webapp"
var root		= os.Getenv("webroot")
var tplFiles	= root + "/static/html/*"
var pubDir		= root + "/static"
var tplDir		= root + "/static"

var db = map[string][]HostInfo{
		"local": {
				HostInfo{"114", "ok", "check"},
				HostInfo{"114.215.137.156", "ok", "system check"},
		},
}

var list = map[string][]string {
	"local": {
		"114.215.137.156",
		"127.0.0.1",
	},
}
