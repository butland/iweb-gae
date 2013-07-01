// +build appengine
package blog

import (
	"appengine"
	"encoding/json"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var apiIndexTPL = template.Must(template.ParseFiles(
	"themes/flat/home.html"))
var apiViewTPL = template.Must(template.ParseFiles(
	"themes/flat/view.html"))

func writeJSON(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	apiIndexTPL.Execute(w, nil)
}

func ApiArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.Args["size"] = 100
	ctx.GAEContext = appengine.NewContext(r)
	articleMetaData := &ArticleMetaData{}
	articleMetaDatas, err := articleMetaData.GetAll(ctx)
	if err != nil {
		serveError(w, err)
		return
	}
	articles := make([]Article, (len(articleMetaDatas)))
	for key, value := range articleMetaDatas {
		articles[key].MetaData = value
		articles[key].Text = template.HTML(blackfriday.MarkdownBasic([]byte(value.Summary)))
	}
	bytes, _ := json.Marshal(articles)
	writeJSON(w, bytes)

}

func ApiGetArticleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.GAEContext = appengine.NewContext(r)
	beginTime := time.Now()
	articleMetaData := &ArticleMetaData{}
	params := r.URL.Query()
	id := params.Get(":id")
	if id != "" {
		articleMetaData.Id = id
	} else {
		year := params.Get(":year")
		month := params.Get(":month")
		day := params.Get(":day")
		title := params.Get(":title")
		//month only in 1~12
		if m, err := strconv.Atoi(month); err != nil || m > 12 || m < 1 {
			http.NotFound(w, r)
			return
		}
		//day only in 1~31
		if d, err := strconv.Atoi(month); err != nil || d > 31 || d < 1 {
			http.NotFound(w, r)
			return
		}
		postTime, err := time.Parse("2006-01-02", year+"-"+month+"-"+day)
		if err != nil {
			serveError(w, err)
			return
		}
		articleMetaData.PostTime = postTime
		articleMetaData.UpdateTime = postTime.AddDate(0, 0, 1)
		articleMetaData.Title = title
	}
	// get article
	err := articleMetaData.GetOne(ctx)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// get comments
	comment := &Comment{ArticleId: articleMetaData.Id}
	comments, err := comment.GetAll(ctx)
	if err != nil {
		serveError(w, err)
		return
	}

	article := &Article{MetaData: *articleMetaData,
		Comments: comments,
		Text: template.HTML([]byte(blackfriday.
			MarkdownBasic(articleMetaData.Content)))}

	data := make(map[string]interface{})
	data["article"] = article
	endTime := time.Now()
	ctx.Args["costtime"] = endTime.Sub(beginTime)
	ctx.Args["title"] = articleMetaData.Title
	data["args"] = ctx.Args
	apiViewTPL.Execute(w, data)
}