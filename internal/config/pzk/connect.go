package pzk

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var zkClient *zk.Conn

func Connect(uri string) error {

	c, _, err := zk.Connect([]string{uri}, time.Second*10)
	if err != nil {
		fmt.Println("Error when connect to Zookeeper", uri, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO ZOOKEEPER: " + uri))

	zkClient = c
	return nil
}
