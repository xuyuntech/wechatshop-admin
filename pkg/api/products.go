package api

import (
	"net/http"
	"github.com/xuyuntech/wechatshop-admin/pkg/models"
	"github.com/xuyuntech/wechatshop-admin/pkg/utils"
	"github.com/xuyuntech/wechatshop-admin"
)

func ProductIndex(w http.ResponseWriter, req *http.Request) {
	var (
		products        []models.Product
		tx              = utils.GetDB(req)
	)
	tx.Find(&products)
	xuyuntech.Resp(w, xuyuntech.STATUS_OK, products)
}

func ProductShow(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ProductShow>>>>>>"))
}

