package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/karim-w/kddd/infra/config"
	_ "github.com/lib/pq"
)

func GetDBCtx(c config.ConfigManger) *sqlx.DB {
	dsn := c.GetConfig().Database.DSN
	if dsn == "" {
		err := fmt.Errorf("missing database dsn variable from enviroment")
		panic(err)
	} else {
		if db, err := sqlx.Connect("postgres", dsn); err != nil {
			panic(err)
		} else {
			db.MustExec(schema)
			return db
		}
	}
}

func GetTestDBCtx(dsn string) *sqlx.DB {
	if dsn == "" {
		err := fmt.Errorf("missing database dsn variable from enviroment")
		panic(err)
	} else {
		if db, err := sqlx.Connect("postgres", dsn); err != nil {
			panic(err)
		} else {
			db.MustExec(schema)
			return db
		}
	}
}
