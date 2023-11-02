package bootstrap

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	asset "github.com/lbemi/lbemi/apps/asset/entity"
	cloud "github.com/lbemi/lbemi/apps/cloud/entity"
	logsys "github.com/lbemi/lbemi/apps/log/entity"
	"github.com/lbemi/lbemi/apps/system/entity"
	"github.com/lbemi/lbemi/pkg/config"
	"github.com/lbemi/lbemi/pkg/global"
)

func InitializeDB(c *config.Config) *gorm.DB {
	switch c.Driver {
	case "mysql":
		return initMysqlGorm(c)
	default:
		return initMysqlGorm(c)
	}
}

func initMysqlGorm(c *config.Config) *gorm.DB {
	dataConfig := c.Database
	if dataConfig.Database == "" {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dataConfig.User,
		dataConfig.Password,
		dataConfig.Host,
		dataConfig.Port,
		dataConfig.Database,
		dataConfig.Charset,
	)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用自动创建外键约束
		Logger:                                   getGormLogger(c),
	})

	if err != nil {
		global.Logger.Errorf("mysql connect failed. err: %v", err)
		os.Exit(-13)
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dataConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dataConfig.MaxOpenConns)
	if dataConfig.IsInitialize {
		migration(db)
	}
	return db
}

func migration(db *gorm.DB) {
	entities := []interface{}{
		&entity.Menu{},
		&entity.User{},
		&entity.Role{},
		&entity.RoleMenu{},
		&entity.UserRole{},
		&config.Rule{},
		&asset.Host{},
		&asset.Group{},
		&asset.HostGroup{},
		&asset.HostAccount{},
		&asset.Account{},
		&cloud.Cluster{},
		&entity.UserResource{},
		&logsys.LogLogin{},
		&logsys.LogOperator{},
	}
	global.Logger.Info("Initializing database ...")
	if err := db.AutoMigrate(entities...); err != nil {
		global.Logger.Errorf("Failed to initialize database. err: %v", err)
		return
	}
}

func getGormLogger(c *config.Config) logger.Interface {
	logMode := getLoggerMode(c.LogMode)
	cnf := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: true,
		Colorful:                  !c.EnableFileLogWrite,
	}
	return logger.New(getGormLogWriter(c), cnf)
}

func getLoggerMode(logMode string) logger.LogLevel {
	switch logMode {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

func getGormLogWriter(c *config.Config) logger.Writer {
	logConfig := c.Log

	var writer io.Writer
	if c.Database.EnableFileLogWrite {
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(logConfig.RootDir, c.LogFilename),
			MaxSize:    logConfig.MaxSize,
			MaxAge:     logConfig.MaxAge,
			MaxBackups: logConfig.MaxBackup,
			Compress:   logConfig.Compress,
		}
	} else {
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}
