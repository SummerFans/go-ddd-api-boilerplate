package dependency

import (
	"database/sql"
	"fmt"
	"go-ddd-api-boilerplate/conf"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {

	pg_conf, _ := conf.GetDatabaseConfig()

	conninfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg_conf["host"], pg_conf["port"], pg_conf["user"], pg_conf["pwd"], pg_conf["database"])

	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(pg_conf["pool_min"].(int))
	db.SetMaxOpenConns(pg_conf["pool_max"].(int))

	if _, err := db.Exec("SELECT 1"); err != nil {
		return nil, fmt.Errorf("pg databases service is not available: %v", err.Error())
	}

	return db, nil
}

func NewRedisConnection() (*redis.Client, error) {

	redis_conf, _ := conf.GetRedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redis_conf["host"], redis_conf["port"]),
		Password: redis_conf["pwd"].(string),    // no password set
		DB:       redis_conf["db"].(int),        // use default DB
		PoolSize: redis_conf["pool_size"].(int), // 连接池大小
	})

	return rdb, nil
}

func Close(db interface{}) {
	switch connection := db.(type) {
	case *sql.DB:
		connection.Close()
	}
}
