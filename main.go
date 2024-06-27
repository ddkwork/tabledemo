package main

import (
	"tabledemo"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("tabledemo", func(w *unison.Window) {
		w.Content().AddChild(tabledemo.New().Layout())
	})
}
