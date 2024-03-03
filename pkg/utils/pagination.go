package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Page struct {
	Offset *int `query:"offset"`
	Limit  *int `query:"limit"`
}

func (p Page) GetPageInformation(context echo.Context) (*Page, error) {
	page := Page{}
	err := (&echo.DefaultBinder{}).BindQueryParams(context, &page)
	fmt.Println(page)
	if err != nil {
		return nil, err
	}
	if page.Offset == nil {
		offset := 0
		page.Offset = &offset
	}
	if page.Limit == nil {
		limit := 10
		page.Limit = &limit
	}
	return &page, nil
}
