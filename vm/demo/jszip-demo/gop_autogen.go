package main

import (
	"github.com/goplus/yap"
	"os"
)

const _ = true

type loadsb3 struct {
	yap.App
}
//line vm/demo/loadsb3/loadsb3_yap.gox:3
func (this *loadsb3) MainEntry() {
//line vm/demo/loadsb3/loadsb3_yap.gox:3:1
	this.Static__0("/dist", os.DirFS("../../dist"))
//line vm/demo/loadsb3/loadsb3_yap.gox:4:1
	this.Static__0("/")
//line vm/demo/loadsb3/loadsb3_yap.gox:6:1
	this.Run(":8888")
}
func main() {
//line vm/demo/loadsb3/loadsb3_yap.gox:6:1
	yap.Gopt_App_Main(new(loadsb3))
}
