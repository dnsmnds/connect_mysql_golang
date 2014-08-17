// Connection with MySQL - Go

// Copyright 2014 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>
// Use of this source code is governed by a BSD-style

package main

import (
    "fmt"
    "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

const (
  DATABASE = "test"
  USERNAME = "root"
  PASSWORD = "12345"
  HOST = "localhost:3306"
)

func main(){
  con, err := sql.Open("mysql",USERNAME+":"+PASSWORD+"@tcp("+HOST+")/"+DATABASE)

  if err != nil {
    fmt.Println("Error connect :( --> %s",err)
  }

  var sqlInsert = "INSERT INTO User (username, password) VALUES (?, ?)"

  _, err = con.Exec(sqlInsert,"golang", "010101")

  if err != nil {
    fmt.Println("Error insert :( --> %s",err)
  }
}
