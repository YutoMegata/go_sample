package main


import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/jinzhu/gorm"
)

type Product struct {
  ID          int    `gorm:"primary_key;not null"`
  ProductName string `gorm:"type:varchar(200);not null"`
  Memo        string `gorm:"type:varchar(400)"`
  Status      string `gorm:"type:char(2);not null"`
}

func getGormConnect() *gorm.DB {
  DBMS := "mysql"
  USER := "root"
  PASS := "password"
  PROTOCOL := "tcp(localhost:3306)"
  DBNAME := "app"
  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
  db, err := gorm.Open(DBMS, CONNECT)

  if err != nil {
      panic(err.Error())
  }

  // DBエンジンを「InnoDB」に設定
  db.Set("gorm:table_options", "ENGINE=InnoDB")

  // 詳細なログを表示
  db.LogMode(true)

  // 登録するテーブル名を単数形にする（デフォルトは複数形）
  db.SingularTable(true)

  // マイグレーション（テーブルが無い時は自動生成）
  db.AutoMigrate(&Product{})

  fmt.Println("db connected: ", &db)
  return db
}

func main() {
  title := "top-megata"

  router := gin.Default()
  router.Static("styles", "./styles")
  router.LoadHTMLGlob("templates/*")
  router.GET("/", func(g *gin.Context) {
    g.HTML(http.StatusOK, "index.html", gin.H{
        "title": title,
    })
})
  router.Run(":8080")
}