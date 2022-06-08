package dbdriver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
	"sync"
)

// ConnectConfig 连接配置
type ConnectConfig struct {
	// 连接名，确定唯一，默认default
	Name      string
	Driver    string
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	Loc       string
}

var connectPool map[string]*gorm.DB

var mux sync.Mutex

func init() {
	connectPool = make(map[string]*gorm.DB, 10)
}

// GetConnect 获取链接
func GetConnect(name string) *gorm.DB {
	if name == "" {
		name = "default"
	}
	if db, ok := connectPool[name]; ok {
		return db
	}
	return nil
}

func OpenConnect(config ConnectConfig) *gorm.DB {
	if config.Name == "" {
		config.Name = "default"
	}
	mux.Lock()
	defer mux.Unlock()
	if db, ok := connectPool[config.Name]; ok {
		return db
	}

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
		url.QueryEscape(config.Loc))

	connect, err := gorm.Open(config.Driver, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	connectPool[config.Name] = connect
	return connect
}

// CloseConnects 关闭链接
func CloseConnects() {
	if len(connectPool) == 0 {
		return
	}
	for s := range connectPool {
		err := connectPool[s].Close()
		if err != nil {
			continue
		}
	}
}

// CloseConnect 关闭某个链接
func CloseConnect(name string) {
	if name == "" {
		name = "default"
	}
	mux.Lock()
	defer mux.Unlock()
	if connect, ok := connectPool[name]; ok {
		connect.DB()
		delete(connectPool, name)
	}
}