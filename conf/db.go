package conf

type DBCfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Prefix   string `json:"prefix"`
	Charset  string `json:"charset"`
}

func GetDBCfg() (db DBCfg) {
	db.Username, _ = Cfg.GetValue("db", "username")
	db.Password, _ = Cfg.GetValue("db", "password")
	db.Database, _ = Cfg.GetValue("db", "database")
	db.Host, _ = Cfg.GetValue("db", "host")
	db.Port, _ = Cfg.GetValue("db", "port")
	db.Prefix, _ = Cfg.GetValue("db", "prefix")
	db.Charset, _ = Cfg.GetValue("db", "charset")
	return
}
