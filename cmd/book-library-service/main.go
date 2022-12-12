package main

import (
	"book-library-service/internal/book-library-service/book"
	"book-library-service/internal/book-library-service/book/db"
	"book-library-service/internal/book-library-service/config"
	"book-library-service/internal/book-library-service/handlers"
	"book-library-service/pkg/client/mysqldb"
	"book-library-service/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMySql := cfg.MySql

	mySqlClient, err := mysqldb.NewClient(cfgMySql.Host, cfgMySql.Port, cfgMySql.Username,
		cfgMySql.Password, cfgMySql.Database, cfgMySql.Net)
	if err != nil {
		panic(err)
	}
	mySqlRepository := db.NewRepositoryM(mySqlClient, logger)
	serviceMySql := book.NewService(mySqlRepository)

	//cfgMongo := cfg.MongoDb
	//mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username,
	//	cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	//if err != nil {
	//	panic(err)
	//}
	//
	//repository := db.NewRepository(mongoDBClient, cfgMongo.Collection, logger)
	//
	//service := book.NewService(repository)

	logger.Info("add cors settings")
	corsSettings := handlers.CorsSettings()

	logger.Info("register book handler")
	handler := book.NewHandler(logger, serviceMySql)

	handler.Register(router)
	h := corsSettings.Handler(router)

	start(h, cfg)
}

func start(h http.Handler, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}
	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      h,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
