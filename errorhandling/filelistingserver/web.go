package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/kentliuqiao/learngo/errorhandling/filelistingserver/filelisting"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

type userError interface {
	error
	Message() string
}

// 精妙！
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// panic
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			// user error
			if usrErr, ok := err.(userError); ok {
				http.Error(w, usrErr.Message(), http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w,
				http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
