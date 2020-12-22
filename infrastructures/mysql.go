package infrastructures

import (
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func NewMySQLSession(dbname string) (*gorm.DB, error) {
  username := os.Getenv("DATABASE_USERNAME")
  password := os.Getenv("DATABASE_PASSWORD")
  host := os.Getenv("DATABASE_HOST")
  port := os.Getenv("DATABASE_PORT")
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
