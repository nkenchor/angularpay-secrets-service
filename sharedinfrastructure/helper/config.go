package helper

import (
	"flag"
	"os"
)

var address string
var port string
var mode string
var dbhost string
var dbname string

func LoadConfig() (string, string, string, string, string) {

	flag.StringVar(&address, "address", "http://localhost", "local host")
	flag.StringVar(&port, "port", "20220", "application ports")
	flag.StringVar(&mode, "mode", "dev", "application mode, either dev or prod")
	flag.StringVar(&dbhost, "dbhost", "mongodb://localhost:27017", "database host")
	flag.StringVar(&dbname, "dbname", "redisTest", "database name")

	flag.Parse()

	for i, val := range flag.Args() {
		os.Args[i] = val
	}

	return address, port, mode, dbhost, dbname
}
