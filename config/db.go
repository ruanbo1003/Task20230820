package config

var (
	MySqlDsn = getEnv(
		"MYSQL_DSN", "user1:abcqwe321@tcp(localhost:3306)/heidi?charset=utf8&parseTime=True&loc=Local",
	)
)
