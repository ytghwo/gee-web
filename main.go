package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ytghwo/gee-web/gee"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Ytghwo\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"ytghwo"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
