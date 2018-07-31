package controller

import (
	"net/http"
	"encoding/json"
)

type Json1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

type Json2 struct {
	Id int `json:"id"`
	Json1  `json:"data"`
}

func JsonController(w http.ResponseWriter, r *http.Request) {

	todos := Json1{
		Id:   2,
		Name: "haha",
		Sex:  "男",
	}

	pp := &Json2{}
	pp.Id = 1
	pp.Json1 = todos

	json.NewEncoder(w).Encode(pp)
}
