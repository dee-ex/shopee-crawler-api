package infrastructures

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func NewMySQLSession() (*gorm.DB, error) {
  username := "root"
  password := "123qwe123qwe"
  host := "localhost"
  port := "3306"
  dbname := "testt"
  dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, err
  }
  mysql_db, err := db.DB()
  if err != nil {
    return nil, err
  }
  mysql_db.SetMaxOpenConns(10)
  // mysql_db.SetConnMaxIdleTime(15)
  mysql_db.SetConnMaxLifetime(15)
  return db, nil
}
