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

// 将文件内容写入到httpResponse中，不好直接创建一个静态目录浏览器直接下载比较好
func (c *FileController) Get() {
	// URL为了避免与其他字符冲突，通常会做特殊处理【特殊字符取%+16进制表示】
	// 将url转回来utf-8
	filepath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) // 去掉开始的“/”
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	// 如果以/开头会以为是绝对路径 以./和attachment/... 表示相对路径
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
