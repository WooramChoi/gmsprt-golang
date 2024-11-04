package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gmsprt-golang/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Run(config *Config) {
	// Logger
	logger, err := setupLogger(config)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	// DBPool
	db, err := setupDBPool(config, gorm.Config{
		Logger: gormLogger.New(
			logger,
			gormLogger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  gormLogger.Warn,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	// GIN
	gin.DefaultWriter = logger.Writer()
	r := gin.New()
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.DBMiddleware(db))

	// TODO set route
	// Service 객체 생성 시 dbPool 및 logger 를 주입. Service 단신으로 동작하도록 함
	// Handler 객체 생성 시 Service 를 주입. Handler 내에서는 파라미터 파싱 하여 서비스 실행 및 결과를 바로 반환
	// 결론적으로, middleware 는 필요 없던게?
	// -> Service 를 middleware 로 주입. 핸들러에선 필요한 Service 를 상황에 맞게 가져와서 사용

	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func setupDBPool(config *Config, db_config gorm.Config) (*gorm.DB, error) {

	var db *gorm.DB
	var err error
	serverConfig := *config

	dbType := serverConfig.Database.Type
	dbHost := serverConfig.Database.Host
	dbPort := serverConfig.Database.Port
	dbDbname := serverConfig.Database.Dbname
	dbUsername := serverConfig.Database.Username
	dbPassword := serverConfig.Database.Password

	switch dbType {
	case "mysql":
		fallthrough
	case "mariadb":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbUsername, dbPassword, dbHost, dbPort, dbDbname)
		db, err = gorm.Open(mysql.Open(dsn), &db_config)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul", dbHost, dbUsername, dbPassword, dbDbname, dbPort)
		db, err = gorm.Open(postgres.Open(dsn), &db_config)
	case "sqlite":
		fallthrough
	default:
		db, err = gorm.Open(sqlite.Open("gorm.db"), &db_config)
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 연결 풀 설정
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func setupLogger(config *Config) (*log.Logger, error) {
	fpLog, err := os.OpenFile("logging.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer fpLog.Close()

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)

	multiWriter := io.MultiWriter(os.Stdout, fpLog)
	logger.SetOutput(multiWriter)
	return logger, nil
}
