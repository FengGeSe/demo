package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	conf "demo/conf"
	grpc "demo/server/grpc"
	http "demo/server/http"

	_ "demo/router/grpc"
	_ "demo/router/http"
)

func init() {
	// flags
	flag.StringVar(&conf.HTTPAddr, "http-addr", GetEnv("HTTPAddr", "0.0.0.0:8080"), "http服务地址")
	flag.StringVar(&conf.GRPCAddr, "grpc-addr", GetEnv("GRPCAddr", "0.0.0.0:5000"), "grpc服务地址")

	// log
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2005-01-02 15:04:05",
	})

	log.WithFields(log.Fields{
		"http-addr": conf.HTTPAddr,
		"grpc-addr": conf.GRPCAddr,
	}).Info("run flags:")
}

func main() {
	flag.Parse()

	errc := make(chan error)

	// http server
	{
		log.WithField("http-addr", conf.HTTPAddr).Info("http server is running...")
		go http.Run(conf.HTTPAddr, errc)
	}

	// grpc server
	{
		log.WithField("grpc-addr", conf.GRPCAddr).Info("grpc server is running...")
		go grpc.Run(conf.GRPCAddr, errc)
	}

	log.WithField("error", <-errc).Info("Exit")
}

func GetEnv(name string, def string) string {
	env := os.Getenv(name)
	if env == "" {
		return def
	}
	return env
}
