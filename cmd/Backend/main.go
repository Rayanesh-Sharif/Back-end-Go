package main

import (
	"RayaneshBackend/api"
	"RayaneshBackend/internal/database"
	"RayaneshBackend/pkg/session"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var err error
	apiData := new(api.API)
	// Get the static folder
	apiData.StaticFolder = getStaticFolder()
	// Connect to database
	apiData.Database, err = database.NewDatabase(getDatabaseDSN())
	if err != nil {
		log.WithError(err).Fatalln("cannot connect to database")
	}
	// Connect to redis
	sessionRedis, err := session.NewStorageFromRedis(getRedisOptions())
	if err != nil {
		log.WithError(err).Fatalln("cannot connect to redis for oauth2")
	}
	apiData.Session = session.NewSession(sessionRedis)
	// Make gin
	r := gin.Default()
	r.Use(api.CORS())
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// Authorization endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/login", apiData.AuthLogin)
		auth.POST("/signup", apiData.AuthSignup)
		auth.GET("/refresh", apiData.AuthRefresh)
		auth.POST("/logout", apiData.AuthLogout)
	}
	// User management
	users := r.Group("/user")
	users.Use(apiData.AuthorizeUserMiddleware())
	{
		users.POST("/password", apiData.UserChangePassword)
		users.POST("/photo", apiData.UserChangeProfilePic)
		users.DELETE("/photo", apiData.UserDeleteProfilePic)
	}
	// Static folder
	r.Static("/static", apiData.StaticFolder)
	// Listen
	srv := &http.Server{
		Addr:    getListenAddress(),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatalln("cannot serve http server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	apiData.Database.Close()
	apiData.Session.Close()
	log.Println("Server exiting")
}
