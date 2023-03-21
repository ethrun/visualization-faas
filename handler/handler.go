package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Excute(c *gin.Context) {
	resp := make(map[string]interface{})
	resp["message"] = "suc"
	database := c.Query("database")
	excute := c.Query("excute")
	resp["command"] = excute
	db, err := sql.Open("mysql", database)
	if err != nil {
		fmt.Println("Error executing connect:", err)
		resp["message"] = err
		c.JSON(http.StatusNotFound, resp)
		return
	}
	rows, err := db.Exec(excute)

	columns, err := rows.RowsAffected()
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		resp["message"] = err
	}

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		resp["message"] = err
	} else {
		resp["resulte"] = columns
	}
	c.JSON(http.StatusOK, resp)
	defer db.Close()
}

func Query(c *gin.Context) {
	resp := make(map[string]interface{})
	resp["message"] = "suc"
	database := c.Query("database")
	excute := c.Query("excute")
	resp["command"] = excute
	db, err := sql.Open("mysql", database)
	if err != nil {
		fmt.Println("Error executing connect:", err)
		resp["message"] = err
		c.JSON(http.StatusNotFound, resp)
		return
	}
	rows, err := db.Query(excute)

	columns, err := rows.Columns()
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		resp["message"] = err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		var v interface{}
		values[i] = &v
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			resp["message"] = err
		}
	}

	resp["values"] = values

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		resp["message"] = err
	} else {
		resp["resulte"] = columns
	}
	c.JSON(http.StatusOK, resp)
	defer db.Close()
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
