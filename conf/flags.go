package conf

import "os"

var (
	HTTPAddr string
	GRPCAddr string
)

func GetEnv(name string, def string) string {
	env := os.Getenv(name)
	if env == "" {
		return def
	}
	return env
}
