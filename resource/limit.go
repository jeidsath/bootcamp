package main

import (
	"fmt"
	"time"
)

type Connection struct {
}

type Database struct {
	availableConnections int
}

// START2 OMIT

func (db *Database) GetConnection() *Connection {
	if db.availableConnections > 0 {
		db.availableConnections--
		fmt.Printf("Connection used. %d connections left.\n", db.availableConnections)
	} else {
		panic("Out of connections!")
	}
	return new(Connection)
}

// END2 OMIT

func (db *Database) ReturnConnection(conn *Connection) {
	db.availableConnections++
}

// START OMIT

func main() {
	db := Database{10}           // 10 connections available
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

// END OMIT
