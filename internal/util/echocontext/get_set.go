package echocontext

import "github.com/labstack/echo/v4"

// GetQuery ...
func GetQuery(c echo.Context) interface{} {
	return c.Get(KeyQuery)
}

// SetQuery ...
func SetQuery(c echo.Context, value interface{}) {
	c.Set(KeyQuery, value)
}

// GetPayload ...
func GetPayload(c echo.Context) interface{} {
	return c.Get(KeyPayload)
}

// SetPayload ...
func SetPayload(c echo.Context, value interface{}) {
	c.Set(KeyPayload, value)
}

// SetParam ...
func SetParam(c echo.Context, key string, value interface{}) {
	c.Set(key, value)
}

// GetParam ...
func GetParam(c echo.Context, key string) interface{} {
	return c.Get(key)
}

// SetByKey ...
func SetByKey(c echo.Context, key string, value interface{}) {
	c.Set(key, value)
}

// GetByKey ...
func GetByKey(c echo.Context, key string) interface{} {
	return c.Get(key)
}

// GetStaff ...
func GetStaff(c echo.Context) interface{} {
	return c.Get(KeyStaff)
}

// SetStaff ...
func SetStaff(c echo.Context, value interface{}) {
	c.Set(KeyStaff, value)
}

// GetUser ...
func GetUser(c echo.Context) interface{} {
	return c.Get(KeyUser)
}

// SetUser ...
func SetUser(c echo.Context, value interface{}) {
	c.Set(KeyUser, value)
}

// SetCurrentUserID ...
func SetCurrentUserID(c echo.Context, id string) {
	c.Set(KeyCurrentUserID, id)
}

// GetCurrentUserID ...
func GetCurrentUserID(c echo.Context) string {
	idInterface := c.Get(KeyCurrentUserID)
	if idInterface == nil {
		return ""
	}
	return idInterface.(string)
}
