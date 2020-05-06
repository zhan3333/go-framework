package conf

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type connections struct {
	Mysql
}

type redis struct {
	RedisDefault redisDefault
}

type redisDefault struct {
	Host     string
	Password string
	Port     int
	Database int
}

type database struct {
	Default     string
	Connections connections
	Redis       redis
}
