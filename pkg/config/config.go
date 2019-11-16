package config

type Mode int

const (
	ModeUnknown Mode = iota
	ModeRelease
	ModeTesting
)

type Config struct {
	Host string
	Port string

	Mode Mode

	DBPath string
}

func Testing() Config {
	conf := Config{
		Host:   "localhost",
		Port:   "8080",
		Mode:   ModeTesting,
		DBPath: "./db.sqlite3",
	}
	return conf
}
