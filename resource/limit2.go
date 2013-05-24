package main

import (
	"fmt"
	"time"
)

type Connection struct {
}

type Database struct {
	availableConnections chan bool
}

// START OMIT

func (db *Database) Initialize() {
	for ii := 0; ii < cap(db.availableConnections); ii++ {
		db.availableConnections <- true
	}
}

func (db *Database) GetConnection() *Connection {
	<-db.availableConnections
	fmt.Printf("Connection used. %d connections left.\n", len(db.availableConnections))
	return new(Connection)
}

func (db *Database) ReturnConnection(conn *Connection) {
	db.availableConnections <- true
}

// END OMIT

// START2 OMIT

func main() {
	db := Database{make(chan bool, 10)} // 10 connections available
	db.Initialize()
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
