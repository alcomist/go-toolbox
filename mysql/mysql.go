package mysql

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type DB struct {
	*sqlx.DB
}

type dbInstance struct {
	mu sync.Mutex
	db map[string]*DB
}

var dbInst dbInstance

func init() {
	dbInst.db = make(map[string]*DB)
}

func MustGet(s string, c *mysql.Config) *DB {

	if c == nil {
		log.Fatal("no mysql.Config")
	}

	dbInst.mu.Lock()
	defer dbInst.mu.Unlock()

	db, ok := dbInst.db[s]
	if ok && db != nil {
		return db
	}

	dsn := c.FormatDSN()

	sqlxDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	dbInst.db[s] = &DB{sqlxDB}
	return dbInst.db[s]
}
