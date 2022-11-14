package todoCtrl

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-backend-go/src/models/todoModel"
	"todo-backend-go/src/service/todoSvc"
)

const RET_OK = 1

func GetTodos(c *gin.Context) {
	userId := 1 // TODO
	todos, err := todoModel.GetTodos(userId)
	if err != nil {
		panic(err) // TODO Error 함수 생성
	}
	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": todos,
	})
	return
}

func AddTodo(c *gin.Context) {
	if q, has := c.GetPostForm("q"); has {
		fmt.Println(q)
	}
	err := todoSvc.AddTodo("todo1", 1)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})

}

// 1. getpostform을 이용해 데이터 받기
func UpdateTodo(c *gin.Context) {
	if q, has := c.GetPostForm("q"); has {

		//var data []todoModel.Todo      //Json의 데이터가 복수일때는 배열로 변환 //data[i].Id 으로 사용
		var data todoModel.Todo          // 받은 JSON을 todo 구조체로 변환하기 위한 변수(한개 데이터만 수신)
		json.Unmarshal([]byte(q), &data) // q를 바이트 슬라이스로 변환하여 넣고, data(todo구조체)의 포인터를 넣어줌
		fmt.Println(data.Id)             // kim 25: 맵에 키를 지정하여 값을 가져옴
		//err := todoSvc.UpdateTodo(222, 222, true)
		err := todoSvc.UpdateTodo(data)

		if err != nil {
			panic(err)
		}
		fmt.Println(q)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})
}

func UpdateTodoCsv(c *gin.Context) {
	if q, has := c.GetPostForm("q"); has {

		// csv으로 받은 데이터를 쉼표 기준으로 분리하여 전달
		data := strings.Split(q, ",")

		fmt.Println(data) // kim 25: 맵에 키를 지정하여 값을 가져옴
		//err := todoSvc.UpdateTodo(222, 222, true)
		err := todoSvc.UpdateTodoCsv(data)

		if err != nil {
			panic(err)
		}
		fmt.Println(q)
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})
}
