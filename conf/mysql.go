package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Host            string `toml:"host" json:"host" env:"MYSQL_HOST"`                                 // Host：MySQL Server 的主机名
	Port            int    `toml:"port" json:"port" env:"MYSQL_PORT"`                                 // Port: MySQL Server 监听的端口
	Username        string `toml:"username" json:"username" env:"MYSQL_USERNAME"`                     // Username：连接的用户名
	Password        string `toml:"password" json:"password" env:"MYSQL_PASSWORD"`                     // Password：连接的密码
	Database        string `toml:"database" json:"database" env:"MYSQL_DATABASE"`                     // Database: 连接使用的数据库
	Charset         string `toml:"charset" json:"charset" env:"MYSQL_CHARSET"`                        // Charset：连接(客户端)解析字符使用的编码
	ParseTime       bool   `toml:"parseTime" json:"parseTime" env:"MYSQL_PARSETIME"`                  // ParseTime：是否解析时间
	MultiStatements bool   `toml:"multiStatements" json:"multiStatements" env:"MYSQL_MUTILSTATEMENT"` // MutilStatement：是否执行执行多条语句

	// MySQL 连接的高级配置
	MaxOpenConn int `toml:"maxOpenConn" json:"maxOpenConn" env:"MYSQL_MAX_OPEN_CONN"` // MaxOpenConn：最大连接数
	MaxIdleConn int `toml:"maxIdleConn" json:"maxIdleConn" env:"MYSQL_MAX_IDLE_CONN"` // MaxIdleConn: 最大闲置连接数
	MaxLifeTime int `toml:"maxLifeTime" json:"maxLifeTime" env:"MYSQL_MAX_LIFE_TIME"` // MaxLifeTime：连接最大持续时间(连接的生命周期)
	MaxIdleTime int `toml:"maxIdleTime" json:"maxIdleTime" env:"MYSQL_MAX_IDLE_TIME"` // MaxIdleTime：闲置连接的生命周期

	// lock 使用锁保护 db：业务可能会有多个 groutine 来读取 db，使用锁保护它只被初始化一次
	lock sync.Mutex
	db   *gorm.DB // db: 保存连接对象
}

// defaultMySQL 返回具有默认值 MySQL 配置对象
func defaultMySQL() *MySQL {
	return &MySQL{
		Host:            "localhost",
		Port:            3306,
		Username:        "vblog",
		Password:        "vblog",
		Database:        "vblog",
		Charset:         "utf8mb4",
		ParseTime:       true,
		MultiStatements: false,
		MaxOpenConn:     200,
		MaxLifeTime:     0,
		MaxIdleConn:     200,
		MaxIdleTime:     0,
	}
}

// String 实现 fmt.Stringer 接口：用于构建 MySQL的连接字符串
func (m *MySQL) String() string {

	// https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&multiStatements=%t",
		m.Username, m.Password, m.Host, m.Port, m.Database, m.Charset, m.ParseTime, m.MultiStatements)
}

// GetConnPool 创建连接对象：使用 database/sql 接口
func (m *MySQL) GetConnPool() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.String())
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", m.String(), err.Error())
	}

	// 对连接池进行配置
	db.SetMaxOpenConns(m.MaxOpenConn) // 最大的打开连接
	db.SetMaxIdleConns(m.MaxIdleConn) // 最大闲置连接
	// 设置连接的生命周期
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	// 设置闲置连接的生命周期
	if m.MaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	// ping 检测。检查 MySQL Server 是否可用
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", m.String(), err.Error())
	}
	return db, nil
}

// GORM 返回一个 GORM 的连接
func (m *MySQL) GORM() (*gorm.DB, error) {
	// 获取 gorm db，保证 m.db 只初始化一次
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		// 获取 database/sql 的连接对象
		c, err := m.GetConnPool()
		if err != nil {
			// 获取连接池失败，就没必要进行下去了
			return nil, err
		}

		// 使用 database/sql 的连接对象初始化 gorm db 对象
		m.db, err = gorm.Open(mysql.New(mysql.Config{
			Conn: c,
		}), &gorm.Config{
			// PrepareStmt：是否创建并缓存预编译语句，可以提高后续调用的性能
			PrepareStmt: true,
			// 对于写操作（insert update delete），为了保证数据的完整性，GORM 会启动一个默认的事务来执行这些SQL
			// SkipDefaultTransaction：是否跳过默认事务
			SkipDefaultTransaction: false,
			// CreateBatchSize：为了高效的批量插入数据，一次最大写入 200 个
			CreateBatchSize: 200,
		})
		if err != nil {
			return nil, err
		}
	}

	return m.db, nil
}
