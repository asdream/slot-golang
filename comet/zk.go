package main

import (
	"github.com/cuixin/cloud/zk"
	"github.com/golang/glog"
	zookeeper "github.com/samuel/go-zookeeper/zk"
	"time"
)

var zkConn *zookeeper.Conn

func InitZK(zkAddrs []string) {
	conn, err := zk.Connect(zkAddrs, time.Second)
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infof("Connect zk[%v] OK!", zkAddrs)
	nodes, event, nerr := zk.GetNodesW(conn, "/MsgBusServers")
	if nerr != nil {
		glog.Fatal(nerr)
	}
	for _, n := range nodes {
		addr, err := zk.GetNodeData(conn, "/MsgBusServers/"+n)
		if err != nil {
			glog.Fatal(err)
		}
		GMsgBusManager.Online(addr)
	}
	for e := range event {
		glog.Info(e)
	}
	zkConn = conn
}

func CloseZK() {
	if zkConn != nil {
		zkConn.Close()
	}
}
