package conf

import (
	"fmt"
	"os"
	"strconv"
)

func GetDatabaseConfig() (map[string]interface{}, error) {

	var pgConf = make(map[string]interface{})
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pwd := os.Getenv("PG_PWD")
	database := os.Getenv("PG_DB")
	pmin := os.Getenv("PG_POOL_MIN")
	pmax := os.Getenv("PG_POOL_MAX")

	if (len(host) == 0) || (len(port) == 0) || (len(user) == 0) || (len(pwd) == 0) || (len(database) == 0) {
		return nil, fmt.Errorf("pg databases service is not available: env is not set")
	}

	pool_min, err := strconv.Atoi(pmin)
	if err != nil {
		return pgConf, err
	}

	pool_max, err := strconv.Atoi(pmax)
	if err != nil {
		return pgConf, err
	}

	pgConf["host"] = host
	pgConf["port"] = port
	pgConf["user"] = user
	pgConf["pwd"] = pwd
	pgConf["database"] = database
	pgConf["pool_min"] = pool_min
	pgConf["pool_max"] = pool_max

	return pgConf, nil
}

func GetRedisConfig() (map[string]interface{}, error) {

	var redisConf = make(map[string]interface{})
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pwd := os.Getenv("REDIS_PWD")
	db := os.Getenv("REDIS_DB")
	pool := os.Getenv("REDIS_POOL")

	if (len(host) == 0) || (len(port) == 0) || (len(pwd) == 0) {
		return nil, fmt.Errorf("redis service is not available: env is not set")
	}

	pool_size, err := strconv.Atoi(pool)
	if err != nil {
		return redisConf, err
	}

	db_num, err := strconv.Atoi(db)
	if err != nil {
		return redisConf, err
	}

	redisConf["host"] = host
	redisConf["port"] = port
	redisConf["pwd"] = pwd
	redisConf["db"] = db_num
	redisConf["pool_size"] = pool_size

	return redisConf, nil
}
