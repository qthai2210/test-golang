package main

//Driver postgre để kết nối db
//go get github.com/lib/pg
//Thư viện để sử dụng sql đơn giản, thay vì sql của go hỗ trợ
//go get github.com/jmoiron/sqlx
//quản lý version của sql, do framework chưa hỗ trợ
//GitHub - rubenv/sql-migrate: SQL schema migrations tool for Go.

import (
	"backend-github/db"
	"backend-github/handler"
	"backend-github/helper"
	"backend-github/repository/repo_impl"
	"backend-github/router"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func init() {
	fmt.Println(">>>>", os.Getenv("APP_NAME"))

}

func main() {
	sql := &db.Sql{
		Host:         "34.80.118.81", //  host.docker.internal// localhost
		Port:         5432,
		Username:     "postgres",
		Password:     "22102003T",
		DatabaseName: "golang",
	}
	sql.Connect()
	// sau khi main function thuc thi xong thi se dong ket noi
	defer sql.Close()
	e := echo.New()
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetUpRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
