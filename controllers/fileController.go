package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	// URL为了避免与其他字符冲突，通常会做特殊处理【16进制表示】
	// 将url转回来utf-8
	filepath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filepath)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	// 将文件f 读取出来送到 responseWriter里去
	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
