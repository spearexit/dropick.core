package main

import (
	"github.com/spearexit/dropick.core/v2/internal/apis"
)

func main() {
	app := apis.New()
	app.Start()
}
