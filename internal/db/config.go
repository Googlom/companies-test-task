package db

type Config struct {
	Host     string // host addr
	Port     int    // port
	User     string // user
	Password string // user password
	DbName   string // companies database name

	MigrationsPath string // directory where migration scripts are stored
}

func validateCfg(cfg Config) error {
	// TODO
	return nil
}
