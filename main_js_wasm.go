package main

import (
	"bytes"
	"fmt"
	"strconv"
	"syscall/js"

	gimage "github.com/meefer/kruskal-vis/image"
	"github.com/meefer/kruskal-vis/kruskal"
)

var done = make(chan struct{})

func main() {
	js.Global().Set("Kruskal", js.FuncOf(kruskalFunc))
	<-done
}

func kruskalFunc(value js.Value, args []js.Value) (ret interface{}) {
	Nstr := args[0].String()
	N, e := strconv.Atoi(Nstr)
	if e != nil || N <= 0 {
		fmt.Println("\"N\" must be of type \"Number\" and greater than zero")
		return
	}

	g := generateGraph(N)
	recorder := gimage.NewRecorder(g)
	kruskal.Kruskal(recorder, g)

	buf := new(bytes.Buffer)
	recorder.Gif(buf)

	gifbytes := buf.Bytes()
	gifarr := js.Global().Get("Uint8Array").New(len(gifbytes))
	js.CopyBytesToJS(gifarr, gifbytes)

	callback := args[len(args)-1]
	callback.Invoke(gifarr)
	return
}
