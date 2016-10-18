package main

import (
	"fmt"
	"net/http"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/goware/urlx"
	"github.com/smartystreets/go-aws-auth"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("s3-signed-redirect")
	viper.AddConfigPath("/etc/s3-signed-redirect/")
	viper.AddConfigPath("$HOME/.s3-signed-redirect")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetEnvPrefix("s3sr")
	viper.AutomaticEnv()
	viper.BindEnv("aws_access_key_id")
	viper.BindEnv("aws_secret_access_key")

	viper.SetDefault("aws_access_key_id", viper.Get("aws_access_key_id"))
	viper.SetDefault("aws_secret_access_key", viper.Get("aws_secret_access_key"))

	viper.SetDefault("debug", false)
	viper.SetDefault("address", "127.0.0.1")
	viper.SetDefault("port", 8080)
	viper.SetDefault("timeout", "3600s")

	log.SetLevel(log.DebugLevel)
	if !viper.GetBool("debug") {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}
	address := viper.Get("address")
	port := viper.GetInt("port")

	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {

		log.Debug("Path: ", c.Request.URL.String())
		log.Debug("S3 Endpoint: ", viper.GetString("s3_endpoint"))

		url, _ := urlx.Parse(viper.GetString("s3_endpoint"))
		url.Path = path.Join(viper.GetString("s3_bucket"), c.Request.URL.String())
		log.Debug("New URL: ", url.String())

		expireTime := time.Now().Add(viper.GetDuration("timeout"))
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			log.WithError(err)
			return
		}

		awsauth.SignS3Url(req, expireTime, awsauth.Credentials{
			AccessKeyID:     viper.GetString("aws_access_key_id"),
			SecretAccessKey: viper.GetString("aws_secret_access_key"),
		})

		log.Debug("Redirect: ", req.URL.String())
		c.Redirect(http.StatusFound, req.URL.String())
	})
	r.Run(fmt.Sprintf("%s:%d", address, port))
}
