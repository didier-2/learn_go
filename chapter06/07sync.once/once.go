package main

import "sync"

type rank struct {
	standard []string
}

var globalRank = &rank{}
var once sync.Once = sync.Once{}

func initGlobalRankStandard(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})
}

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct {
}
type DBFactory interface {
	GetConnection() *Conn
}

func initMysqlfac(connStr string) DBFactory {
	return &MysqlDBFactory{}
}

type MysqlDBFactory struct {
	once sync.Once
}

func (MysqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMysqlfac("")
	})

	//todo
	return nil
}

func main() {
	//standard := []string{"asia"}
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		initGlobalRankStandard(standard)
	//	}()
	//}
	//connStr:="xxxxxxxxxx"
}
