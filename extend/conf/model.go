package conf

// server configuration
type server struct {
	GrpcPort int `mapstructure:"grpcPort"`
	HTTPPort int `mapstructure:"httpPort"`
}

// database configuration
type database struct {
	DBType   string `mapstructure:"dbtype"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}
