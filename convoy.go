package convoy

import (
	"database/sql"
	"fmt"
	"time"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Configuring sql.DB for Better Performance
// https://www.alexedwards.net/blog/configuring-sqldb
// https://making.pusher.com/production-ready-connection-pooling-in-go/
// https://github.com/golang/go/issues/9851
// https://charlesxu.io/go-opts/

const defaultPort int = 3306
const defaultMaxOpenConns int = 2
const defaultMaxIdleConns = 2

// defaultConnMaxLifeTime if opt is not specify duration is 5 minutes
const defaultConnMaxLifeTime time.Duration = time.Minute * time.Duration(5)

// Options defines a set of options we can pass to the builder
type Options struct {
	Host, User, Pass, DB string
	Port                 int
	// ConnMaxLifetime
	ConnMaxLifetime time.Duration
	// MaxOpenCons used when calling db.SetMaxOpenConns
	// by default there's no limit of the open connections
	// (in-use + idle)
	MaxOpenCons int
	// MaxIdleConns when is none a new connection has to be created from scratch
	// It should alwasy be less or equal than MaxOpenCons
	// It is recommended to be at least 1
	MaxIdleCons                int
	ParseTime, MultiStatements bool
	GormConfig                 *gorm.Config
}

func NewOps(from Options) Options {
	if from.Port == 0 {
		from.Port = defaultPort
	}
	if from.ConnMaxLifetime == 0 {
		from.ConnMaxLifetime = defaultConnMaxLifeTime
	}
	if from.MaxOpenCons == 0 {
		from.MaxOpenCons = defaultMaxOpenConns
	}
	if from.MaxIdleCons == 0 {
		from.MaxIdleCons = defaultMaxIdleConns
	}

	return from
}

func NewDSN(opt Options) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=%t&multiStatements=%t", opt.User, opt.Pass, opt.Host, opt.Port, opt.DB, opt.ParseTime, opt.MultiStatements)
}

func NewDB(opt Options) (*sql.DB, error) {
	dsn := NewDSN(opt)
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	con.SetConnMaxLifetime(opt.ConnMaxLifetime)
	con.SetMaxOpenConns(opt.MaxOpenCons)
	con.SetMaxIdleConns(opt.MaxIdleCons)
	return con, nil
}

func NewGormDB(opt Options) (*gorm.DB, error) {
	dsn := NewDSN(opt)
	con, err := gorm.Open(gmysql.Open(dsn), opt.GormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := con.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(opt.ConnMaxLifetime)
	sqlDB.SetMaxOpenConns(opt.MaxOpenCons)
	sqlDB.SetMaxIdleConns(opt.MaxIdleCons)
	return con, nil
}
