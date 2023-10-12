package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	// セッションがDBに存在しない場合（未ログイン）の場合は、"/" にアクセスする
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		// セッションがDBに存在する場合（ログイン済）の場合は、"/todos" にアクセスする
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	// セッションがDBに存在しない場合（未ログイン）の場合
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		// セッションがDBに存在する場合（ログイン済）の場合は、"/todos" にアクセスする
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
