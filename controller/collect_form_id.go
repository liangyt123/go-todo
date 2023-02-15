package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/liangyt123/go-todo/middleware"
	"github.com/liangyt123/go-todo/models"
	"github.com/liangyt123/go-todo/utils"
)

func init() {
	s := g.Server()
	s.BindHandler("/collect", collectFormId)
}

type CollectionFormID struct {
	FormID string `json:"formId"`
}

// 收集用户的formId
func collectFormId(r *ghttp.Request) {
	collection := new(CollectionFormID)
	r.GetToStruct(collection)
	openID, _ := middleware.GetOpenID(r)
	err := models.CollectFormID(openID, collection.FormID)
	if err != nil {
		r.Response.WriteJson(utils.ErrorResponse(err.Error()))
		r.Exit()
	}
	r.Response.WriteJson(utils.SuccessResponse("OK"))
}
