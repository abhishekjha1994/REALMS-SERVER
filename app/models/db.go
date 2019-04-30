package models

import (
	"database/sql"

	eu "github.com/REALMS-SERVER/lib/errutil"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //for postgres
	"github.com/revel/revel"
)

var (
	// Dbm is db handle
	Dbm *gorp.DbMap
)

func init() {
	initDB()

}
func initDB() {
	var dbInfo string

	dbInfo = "host=chipmonkdb.ckidpyxivgoi.ap-south-1.rds.amazonaws.com user=raghavkm password=ch!pMonk " +
		"dbname=realms sslmode=disable"

	Db, err := sql.Open("postgres", dbInfo)
	if Db == nil || err != nil {
		revel.ERROR.Println("could not connect to postgres", dbInfo)
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: Db, Dialect: gorp.PostgresDialect{}}
}
func updateModels() {
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	t := Dbm.AddTableWithName(Reg{}, "users")
	t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Username": 50,
		"Name":     100,
	})

	err1 := Dbm.CreateTablesIfNotExists()
	eu.CheckErr(err1, "Error in creating Dynamic tables")

	Dbm.TraceOn("[gorp]", revel.INFO)
	err := Dbm.CreateTablesIfNotExists()
	eu.CheckErr(err, "Error in creating tables")
}
