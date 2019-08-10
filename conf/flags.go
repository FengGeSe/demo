package conf

import "os"

var (
	HttpAddr string
	GrpcAddr string
)

func GetEnv(name string, def string) string {
	env := os.Getenv(name)
	if env == "" {
		return def
	}
	return env
}
