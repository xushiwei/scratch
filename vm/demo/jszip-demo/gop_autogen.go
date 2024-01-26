package main

import (
	"github.com/goplus/yap"
	"os"
)

const _ = true

type jszip struct {
	yap.App
}
//line vm/demo/jszip-demo/jszip_yap.gox:3
func (this *jszip) MainEntry() {
//line vm/demo/jszip-demo/jszip_yap.gox:3:1
	this.Static__0("/dist", os.DirFS("../../dist"))
//line vm/demo/jszip-demo/jszip_yap.gox:4:1
	this.Static__0("/")
//line vm/demo/jszip-demo/jszip_yap.gox:6:1
	this.Run(":8888")
}
func main() {
//line vm/demo/jszip-demo/jszip_yap.gox:6:1
	yap.Gopt_App_Main(new(jszip))
}
