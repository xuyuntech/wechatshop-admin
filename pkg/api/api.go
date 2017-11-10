package api

import (
	"github.com/qor/publish2"
	"github.com/qor/admin"
	"github.com/xuyuntech/wechatshop-admin/pkg/db"
	"net/http"
	"github.com/qor/middlewares"
	"github.com/Sirupsen/logrus"
	"github.com/xuyuntech/wechatshop-admin/pkg/models"
	"github.com/qor/media/media_library"
	"github.com/qor/media"
	"github.com/qor/qor"
	"bytes"
	"html/template"
	"fmt"
)

type api struct {
	listen string
}

type ApiConfig struct {
	Listen string
}

func NewApi(config *ApiConfig) (*api, error) {
	return &api{
		listen: config.Listen,
	}, nil
}

func (a *api) Run() error {
	mux := http.NewServeMux()

	aAdmin := admin.New(&admin.AdminConfig{
		SiteName: "虚云",
		Auth: nil, DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
	})
	aAdmin.AddMenu(&admin.Menu{Name: "控制台", Link: "/admin"})

	ProductImagesResource := aAdmin.AddResource(
		&models.ProductImage{},
		&admin.Config{Name: "药品图片", Menu: []string{"基础数据管理"}, Priority: -1})
	ProductImagesResource.Filter(&admin.Filter{
		Name:       "SelectedType",
		Label:      "图片类型",
		Operations: []string{"contains"},
		Config:     &admin.SelectOneConfig{Collection: [][]string{{"video", "Video"}, {"image", "Image"}, {"file", "File"}, {"video_link", "Video Link"}}},
	})
	ProductImagesResource.IndexAttrs("File", "Title")

	product := aAdmin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"药品管理"}})
	product.Meta(&admin.Meta{
		Name: "MainImage",
		Label: "头图",
		Config: &media_library.MediaBoxConfig{
			RemoteDataResource: ProductImagesResource,
			Max:                1,
			Sizes: map[string]*media.Size{
				"main": {Width: 560, Height: 700},
			},
		},
	})
	product.Meta(
		&admin.Meta{
			Name: "MainImageURL",
			Label: "头图链接",
			Valuer: func(record interface{}, context *qor.Context) interface{} {
				if p, ok := record.(*models.Product); ok {
					result := bytes.NewBufferString("")
					tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
					tmpl.Execute(result, map[string]string{"image": p.MainImageURL()})
					return template.HTML(result.String())
				}
				return ""
			},
		})
	product.UseTheme("grid")

	product.SearchAttrs("Name")
	product.IndexAttrs("MainImageURL", "Name", "Price", "VersionName", "PublishLiveNow")
	product.EditAttrs(
		&admin.Section{
			Title: "基础信息",
			Rows: [][]string{
				{"Name"},
				{"Price"},
				{"MainImage"},
			}},
		"PublishReady",
	)
	product.NewAttrs(product.EditAttrs())
	product.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if product, ok := record.(*models.Product); ok {
				return fmt.Sprintf("/products/%v", product.Code)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})

	mux.Handle("/", Router())
	aAdmin.MountTo("/admin", mux)

	logrus.Debugf("Listening on: %v\n", a.listen)
	return http.ListenAndServe(a.listen, middlewares.Apply(mux))
}