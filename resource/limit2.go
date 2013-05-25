package main

import (
	"fmt"
	"time"
)

type Connection struct {
}

func (*Connection) activate() {}

func (*Connection) deactivate() {}

type Database struct {
	availableConnections chan *Connection
}

// START OMIT

func (db *Database) Initialize(numConnections int) {
	db.availableConnections = make(chan *Connection, numConnections)
	for ii := 0; ii < cap(db.availableConnections); ii++ {
		db.availableConnections <- new(Connection)
	}
}

func (db *Database) GetConnection() *Connection {
	conn := <-db.availableConnections
	fmt.Printf("Connection used. %d connections left.\n",
		len(db.availableConnections))
	conn.activate()
	return conn
}

func (db *Database) ReturnConnection(conn *Connection) {
	conn.deactivate()
	db.availableConnections <- conn
}

// END OMIT

// START2 OMIT

func main() {
	db := Database{}
	db.Initialize(10)            // 10 connections available
	for ii := 0; ii < 11; ii++ { // But this loop goes up to 11!
		go func() {
			for {
				conn := db.GetConnection()
				time.Sleep(1 * time.Second)
				db.ReturnConnection(conn)
			}
		}()
	}
	select {} // Block forever
}

// END2 OMIT
