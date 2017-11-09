package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/sorting"
	"github.com/qor/publish2"
	"github.com/qor/media/media_library"
	"strings"
	"github.com/qor/validations"
	"encoding/json"
)

type Product struct {
	gorm.Model
	l10n.Locale
	sorting.SortingDESC

	Name string
	Code                  string       `l10n:"sync"`
	MainImage             media_library.MediaBox

	Price                 float32          `l10n:"sync"`


	publish2.Version
	publish2.Schedule
	publish2.Visible
}


func (product Product) DefaultPath() string {
	defaultPath := "/"
	return defaultPath
}

func (product Product) MainImageURL(styles ...string) string {
	style := "main"
	if len(styles) > 0 {
		style = styles[0]
	}

	if len(product.MainImage.Files) > 0 {
		return product.MainImage.URL(style)
	}

	return "/images/default_product.png"
}


type ProductImage struct {
	gorm.Model
	Title        string
	SelectedType string
	File         media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
}


func (productImage ProductImage) Validate(db *gorm.DB) {
	if strings.TrimSpace(productImage.Title) == "" {
		db.AddError(validations.NewError(productImage, "Title", "Title can not be empty"))
	}
}

func (productImage *ProductImage) SetSelectedType(typ string) {
	productImage.SelectedType = typ
}

func (productImage *ProductImage) GetSelectedType() string {
	return productImage.SelectedType
}

func (productImage *ProductImage) ScanMediaOptions(mediaOption media_library.MediaOption) error {
	if bytes, err := json.Marshal(mediaOption); err == nil {
		return productImage.File.Scan(bytes)
	} else {
		return err
	}
}

func (productImage *ProductImage) GetMediaOption() (mediaOption media_library.MediaOption) {
	mediaOption.Video = productImage.File.Video
	mediaOption.FileName = productImage.File.FileName
	mediaOption.URL = productImage.File.URL()
	mediaOption.OriginalURL = productImage.File.URL("original")
	mediaOption.CropOptions = productImage.File.CropOptions
	mediaOption.Sizes = productImage.File.GetSizes()
	mediaOption.Description = productImage.File.Description
	return
}
