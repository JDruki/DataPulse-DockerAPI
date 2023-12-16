package dao

import (
	"Auto/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	Db  *sql.DB
	err error
)

// 数据库连接池缓存
var dbPool = make(map[string]*sql.DB)
var mu sync.Mutex

// 创建数据库连接
func createDBConnection(repoName, dsn string) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	db, ok := dbPool[repoName]
	if ok {
		return db, nil
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	// 确保连接正常
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, fmt.Errorf("数据库不健康: %w", err)
	}

	dbPool[repoName] = db

	return db, nil
}
func InitDatabase(repoName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MC.Name,
		config.Conf.MC.Password,
		config.Conf.MC.Host,
		repoName)

	db, err := createDBConnection(repoName, dsn)
	if err != nil {
		return err
	}
	Db = db
	fmt.Printf("数据库已连接: %s\n", repoName)
	return nil
}
