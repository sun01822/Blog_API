package utils

import (
	"github.com/labstack/echo/v4"
)

type Page struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func (p Page) GetPageInformation(context echo.Context) (Page, error) {

	page := Page{}
	err := (&echo.DefaultBinder{}).BindQueryParams(context, &page)

	if err != nil {
		return page, err
	}

	return page, nil
}
