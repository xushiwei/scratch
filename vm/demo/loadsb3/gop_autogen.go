package main

import "github.com/goplus/yap"

const _ = true

type loadsb3 struct {
	yap.App
}
//line vm/demo/loadsb3/loadsb3_yap.gox:1
func (this *loadsb3) MainEntry() {
//line vm/demo/loadsb3/loadsb3_yap.gox:1:1
	this.Static__0("/")
//line vm/demo/loadsb3/loadsb3_yap.gox:3:1
	this.Run(":8888")
}
func main() {
//line vm/demo/loadsb3/loadsb3_yap.gox:3:1
	yap.Gopt_App_Main(new(loadsb3))
}
