package database

// ConfDB is structure config database
type ConfDB struct {
	Driver          string
	Host            string
	Schema          string
	User            string
	Pass            string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
}
