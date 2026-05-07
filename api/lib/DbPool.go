package lib

import (
	"database/sql"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

// 定义一个全局对象db
var Db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {

	// 设置format json
	log.SetFormatter(&log.TextFormatter{})
	// 设置输出警告级别
	log.SetLevel(log.InfoLevel)
	logfile, _ := os.OpenFile("./logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	log.SetOutput(logfile)

	var cfg *ini.File
	var DatabaseSetting = &Database{}
	cfg, _ = ini.Load("conf.ini")
	cfg.Section("database").MapTo(DatabaseSetting)

	// 支持 Render 等云端部署时使用环境变量
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = DatabaseSetting.Host
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = DatabaseSetting.User
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = DatabaseSetting.Password
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		name = DatabaseSetting.Name
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		// 如果是本地，默认从 conf.ini 里面可能自带端口 (如 127.0.0.1:3306)，就不需要再加端口了
		// 但如果是 TiDB Serverless, 通常需要 :4000
		port = ""
	} else {
		port = ":" + port
	}

	tlsConfig := "false"
	if os.Getenv("DB_TLS") == "true" {
		tlsConfig = "true" // 针对 TiDB
	}

	// 自动检测 TiDB Cloud 域名，强制开启 TLS
	if strings.Contains(host, "tidbcloud.com") {
		tlsConfig = "true"
	}

	// 构造 DSN：如果 host 本身已经带了端口 (比如 conf.ini 里的 127.0.0.1:3306)，就不再拼接 port
	dsn = user + ":" + password + "@tcp(" + host + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 只有在明确需要 tls 的时候才加上 &tls=true
	if tlsConfig == "true" {
		dsn += "&tls=true"
	}

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}
	Db.SetMaxOpenConns(1024)               //   设置连接数总数, 需要根据实际业务来测算, 应小于 mysql.max_connection (应该远远小于), 后续根据指标进行调整
	Db.SetMaxIdleConns(100)                //  设置最大空闲连接数, 该数值应该小于等于 SetMaxOpenConns 设置的值
	Db.SetConnMaxLifetime(0)               // 设置连接最大生命周期, 默认为 0(不限制), 我不建议设置该值, 只有当 mysql 服务器出现问题, 会导致连接报错, 恢复后可以自动恢复正常, 而我们配置了时间也不能卡住出问题的时间, 配置小还不如使用 SetConnMaxIdleTime 来解决
	Db.SetConnMaxIdleTime(4 * time.Second) // 设置空闲状态最大生命周期, 该值应小于 mysql.wait_timeout 的值, 以避免被服务端断开连接, 产生报错影响业务， 一般可以配置 1天。

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
