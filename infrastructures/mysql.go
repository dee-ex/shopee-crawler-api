package infrastructures

import (
  "os"
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"

  "github.com/joho/godotenv"
)


func OpenMySQLSession(dbname string) (*gorm.DB, error) {
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

func NewMySQLSession() (*gorm.DB, error) {
  return OpenMySQLSession(os.Getenv("DATABASE_MAINDATABASENAME"))
}

func NewMockMySQLSession() (*gorm.DB, error) {
  err := godotenv.Load("../../env/database.env")
  if err != nil {
    log.Fatal("Error loading database.env file")
  }
  if os.Getenv("DATABASE_CONFIGURED") == "NO" {
    log.Fatal("Database is not configured yet")
  }
  return OpenMySQLSession(os.Getenv("DATABASE_TESTDATABASENAME"))
}
