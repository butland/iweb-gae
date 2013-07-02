package blog

import (
	"appengine"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func SaveCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.GAEContext = appengine.NewContext(r)
	comment := &Comment{}
	comment.ArticleId = r.FormValue("articleId")
	comment.ParentId = r.FormValue("parentId")
	comment.Author = r.FormValue("name")
	comment.Email = r.FormValue("email")
	comment.Website = r.FormValue("website")
	comment.Flag = 2 //public
	comment.Content = r.FormValue("content")
	now := time.Now()
	comment.PostTime = now
	err := comment.Save(ctx)
	message := Message{true, "Comment is saved!", comment}
	if err != nil {
		message.Ok = false
		message.Text = err.Error()
		//serveError(w, err)
	}
	bytes, _ := json.Marshal(message)
	writeJSON(w, bytes)
	//urlStr := r.FormValue("urlStr")
	//http.Redirect(w, r, urlStr, http.StatusFound)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.GAEContext = appengine.NewContext(r)
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	comment := new(Comment)
	fmt.Println(id)
	comment.Id = id
	err := comment.Delete(ctx)
	if err != nil {
		serveError(w, err)
		return
	}
	http.Redirect(w, r, "/admin/comment", http.StatusFound)
}

func ListCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.GAEContext = appengine.NewContext(r)
	//beginTime := time.Now()
	if r.FormValue("size") != "" {
		size, err := strconv.Atoi(r.FormValue("size"))
		if err != nil {
			serveError(w, err)
			return
		}
		ctx.Args["size"] = size
	}
	if r.FormValue("pageSize") != "" {
		pageSize, err := strconv.Atoi(r.FormValue("pageSize"))
		if err != nil {
			serveError(w, err)
			return
		}
		ctx.Args["pageSize"] = pageSize
		if pageSize > 100 {
			http.NotFound(w, r)
			return
		}
	}

	comments, err := GetAllComments(ctx)
	if err != nil {
		serveError(w, err)
		return
	}
	bytes, _ := json.Marshal(comments)
	writeJSON(w, bytes)
	/*
		data := make(map[string]interface{})
		data["data"] = comments
		prePageSize := ctx.Args["pageSize"].(int) - 1
		ctx.Args["prePageSize"] = prePageSize
		if prePageSize > 0 {
			ctx.Args["hasPre"] = true
		}
		nextPageSize := ctx.Args["pageSize"].(int) + 1
		ctx.Args["nextPageSize"] = nextPageSize

		ctx.Args["url"] = r.URL.Path
		endTime := time.Now()
		ctx.Args["costtime"] = endTime.Sub(beginTime)
		data["args"] = ctx.Args

		commentTPL.ExecuteTemplate(w, "main", data)
	*/
}

func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := GetContext()
	ctx.GAEContext = appengine.NewContext(r)
	// get comments
	comment := &Comment{ArticleId: r.FormValue("id")}
	comments, err := comment.GetAll(ctx)

	if err != nil {
		serveError(w, err)
		return
	}
	bytes, _ := json.Marshal(comments)
	writeJSON(w, bytes)
}
