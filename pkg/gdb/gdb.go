package gdb

import (
	"fmt"
	"log"
)

// gorm v2 版本多连接, 使用需要预先配置 DefaultName ConnConfigs 两项
// 建议项目启动时即初始化所有连接:
// 1. 初始化默认连接 InitDef()
// 2. 初始化所有连接 InitAll()
// 需要访问默认连接 (Entry) , 可以使用以下方法
// 1. gdb.DB
// 2. gdb.Def()
// 3. gdb.Conn(gdb.DefaultName)
// 要访问 gorm 对象, 可以使用
// 1. gdb.Conn("connection name").GORM

var (
	// 默认连接配置key
	DefaultName = "default"
	// 连接配置数组
	ConnConfigs map[string]MySQLConf
)

// 数据库连接映射
var connections = map[string]*Entry{}

// 默认数据库查询对象
var DB *Entry

// init DefaultName setting connection
func InitDef() (*Entry, error) {
	entry, err := NewEntry(ConnConfigs[DefaultName])
	if err != nil {
		return entry, err
	}
	connections[DefaultName] = entry
	DB = entry
	return entry, nil
}

// init all config connections
func InitAll() error {
	for name := range ConnConfigs {
		if connections[name] == nil {
			entry, err := NewEntry(ConnConfigs[name])
			if err != nil {
				return err
			}
			connections[name] = entry
		}
	}
	return nil
}

func Close() {
	for k, entry := range connections {
		if err := entry.SQLDB.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
}

// get default connection
// connection can't create will be panic
func Def() *Entry {
	return Conn(DefaultName)
}

// 获取指定的连接, 当创建新连接失败时, 会抛出异常
// 推荐在程序运行时即执行 InitAll() 初始化所有连接
func Conn(name string) *Entry {
	var err error
	if conn, ok := connections[name]; ok {
		return conn
	}
	if c, ok := ConnConfigs[name]; ok {
		connections[name], err = NewEntry(c)
		if err != nil {
			panic(fmt.Sprintf("Connect mysql (%s: %s) failed: %+v", name, c.String(), err))
			return nil
		}
		return connections[name]
	}
	panic(fmt.Sprintf("Can't read mysql config: %s", name))
	return nil
}
