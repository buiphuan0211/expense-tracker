package routemiddleware

import (
	"github.com/labstack/echo/v4"
)

// Locale set default locale for user, based on header
func Locale(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Default support Vietnamese only
		c.Set("lang", "vi")
		return next(c)
	}
}

// // List allow languages
// var languages = []string{"vi", "en"}
//
// // Locale set default locale for user, based on header
// func Locale(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		lang := c.Request().Header.Get("Accept-Language")
//
// 		// Set default language
// 		if !funk.Contains(languages, lang) {
// 			lang = "vi"
// 		}
//
// 		// Set locale
// 		c.Set("lang", lang)
// 		return next(c)
// 	}
// }
