/*
Copyright 2018 The AmrToMp3 Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    //"math/rand"
    "net/http"
)

type LoginForm struct {
    User     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}

// Binding from JSON
type LoginJson struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}


func main() {
    router := gin.Default()
    //登录入口
    passport := router.Group("/passport")
    {
        passport.POST("/login", func(c *gin.Context) {
            var json LoginJson
            if err := c.ShouldBindJSON(&json); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            if json.User == "user" && json.Password == "password" {
                c.JSON(http.StatusOK, gin.H{
                    "status": 200,
                    "token": "sdfsdf",
                    "message": "you are logged in",
                })
                return
            } else {
                    c.JSON(401, gin.H{"status": "unauthorized"})
            }
        })
    }

    //设置cookie
    router.GET("/cookie", func(c *gin.Context) {

        cookie, err := c.Cookie("gin_cookie")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

        fmt.Printf("Cookie value: %s \n", cookie)
    })

    //定义默认路由
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,
            "error":  "404, page not exists!",
        })
    })
    router.Run(":8888")
}
