package main

import (
	"golang.org/x/net/webdav"
	"net/http"
	"log"
)

func main() {
	fs := webdav.NewMemFS()
	ls := webdav.NewMemLS()
	h := &webdav.Handler{
		FileSystem: fs,
		LockSystem: ls,
		Prefix: "/",
		Logger: func(req *http.Request, err error) {
			if err != nil {
				log.Printf("server: error %s", err)
			} else {
				log.Printf("server: %v", req)
			}
		},
	}

	srv := http.Server{
		Handler: h,
		Addr: ":4321",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Server exited: %s", err)
	}

}
