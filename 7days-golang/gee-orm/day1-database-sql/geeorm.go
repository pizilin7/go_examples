package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"

	_ "github.com/mattn/go-sqlite3"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return e, nil
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database Success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
