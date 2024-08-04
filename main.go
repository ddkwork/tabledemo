package main

import (
	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
	"tabledemo"
)

func main() {
	app.Run("tabledemo", func(w *unison.Window) {
		//w.Content().AddChild(widget.NewButton("111", nil))
		w.Content().AddChild(tabledemo.New().Layout())
	})
}
