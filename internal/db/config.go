package db

type Config struct {
	Host        string // host addr
	Port        int    // port
	User        string // user
	Password    string // user password
	CompaniesDb string // companies database name
}

func validateCfg(cfg Config) error {
	// TODO
	return nil
}
