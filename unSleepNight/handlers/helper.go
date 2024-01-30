package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"unsleepNight/models"
)

// 通过cookie判断用户是否已经登陆
func session(c *gin.Context) (sess models.Session, err error) {
	cookie, err := c.Cookie("cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return
}

// 解析HTML模版
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应html
func generateHTML(c *gin.Context, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("view/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		return
	}
}

// Version 返回版本号
func Version() string {
	return "0.1"
}
