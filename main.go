package main

import (
	"fmt"
	"gig/_db_api"
	"github.com/bmatsuo/lmdb-go/lmdb"
)


func main(){
	var env *lmdb.Env;
	fmt.Println("HAllo");
	env = _db_api.Db_init();
	_db_api.Db_open(env, "testdb");
	_db_api.Db_write(env, "testdb", "moin", "servus");
	fmt.Println(_db_api.Db_read(env, "moin"));

}