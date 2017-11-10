package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"

	"github.com/xuyuntech/wechatshop-admin/pkg/db"
)

var rootMux *http.ServeMux
var WildcardRouter *wildcard_router.WildcardRouter

func Router() *http.ServeMux {
	if rootMux == nil {
		router := chi.NewRouter()

		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				var (
					tx         = db.DB
					qorContext = &qor.Context{Request: req, Writer: w}
				)

				if locale := utils.GetLocale(qorContext); locale != "" {
					tx = tx.Set("l10n:locale", locale)
				}

				ctx := context.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
				next.ServeHTTP(w, req.WithContext(ctx))
			})
		})

		router.Get("/", HomeIndex)
		router.Get("/products", ProductIndex)
		router.Get("/products/{code}", ProductShow)

		rootMux = http.NewServeMux()


		WildcardRouter = wildcard_router.New()
		WildcardRouter.MountTo("/", rootMux)
		WildcardRouter.AddHandler(router)
	}
	return rootMux
}
