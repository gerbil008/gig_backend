package _db_api

import (
    "log"
    "github.com/bmatsuo/lmdb-go/lmdb"
    "os"
)

func Db_init() *lmdb.Env{
    env, err := lmdb.NewEnv()
    if err != nil {
        log.Fatal(err)
    }
    err = env.SetMaxDBs(1)
    if err != nil {
        log.Fatal(err)
    }
    return env
}

func Db_open(env *lmdb.Env, name string) int{
    os.Mkdir(name, 0755);
    err := env.Open(name, 0, 0664)
    if err != nil {
        log.Fatal(err)
    }
    return 0;
}

func Db_write(env *lmdb.Env, name string, value string, key string) int{
    var dbi lmdb.DBI;
    err := env.Update(func(txn *lmdb.Txn) (err error) {
        dbi, err = txn.OpenDBI(name, lmdb.Create)
        if err != nil {
            return
        }
        err = txn.Put(dbi, []byte(key), []byte(value), 0)
        return 
    })
    if err != nil {
        log.Fatal(err)
        return 1;
    }
    return 0;
}

func Db_read(env *lmdb.Env, key string) string{
    var dbi lmdb.DBI;
    var result string;
    err := env.View(func(txn *lmdb.Txn) error {
        val, err := txn.Get(dbi, []byte(key))
        if err != nil {
            return err
        }
        result = string(val);
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }

    return result;

}



