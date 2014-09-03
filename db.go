package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/coopernurse/gorp"
    "gopkg.in/mgo.v2"
)

func initDb() *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish

    //db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
    db, err := sql.Open("mysql", "root:1@unix(/var/run/mysqld/mysqld.sock)/mydb")
    checkErr(err, "sql.Open failed")

    // construct a gorp DbMap
    // dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

    dbmap.AddTableWithName(Company{}, "company").SetKeys(true, "Id")
    dbmap.AddTableWithName(Feedback{}, "feedbacks").SetKeys(true, "Id")
    dbmap.AddTableWithName(FeedbackForm{}, "feedbacks").SetKeys(true, "Id")
	// first check
	var rs []string
	dbmap.Select(&rs,   "select company from myReport")
	var c Company
	dbmap.SelectOne(&c,   "select * from company where id = ? ", 1)
	fmt.Println("com1 ---- ", c)
//	var alls []Report
//	dbmap.Select(&alls,   "select *  from myReport order by id")
//	for x, r := range alls {
//		fmt.Printf("a --- %d,  %v\n", x, r)
//	}
//	fmt.Println(alls)

	for x, r := range rs {
		fmt.Printf("b --- %d,  %v\n", x, r)
	}
	fmt.Println("----- print table")
	// -----
	return dbmap

}

func initMgo() *mgo.Session{
    session, err := mgo.Dial("127.0.0.1")
    if err != nil { panic(err) }
    session.SetMode(mgo.Monotonic, true)
	return session
}

// this struct is for initdb test try operation, so put it here
type Company struct {
	Id			int64  `form:"Id"`
	Name		string `form:"Name"`
	Person		string `form:"Person"`
	Phone		string `form:"Phone"`
	Time		string
	Operator	string
	Info		string `form:"Info"`
	Usedomain	string `form:"Usedomain"`
}
