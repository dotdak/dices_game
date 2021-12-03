package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/dotdak/dices_game/pkg"
)

const maxLeaderBoard = 10

func main() {
	port := ":8080"
	leaderboard := &pkg.LeaderBoard{
		Data: make([]*pkg.User, 0, maxLeaderBoard+1),
		Size: maxLeaderBoard,
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		page, err := ioutil.ReadFile("index.html")
		if err != nil {
			(&pkg.Response{
				Error: err,
				Code:  http.StatusInternalServerError,
			}).Write(rw)
			return
		}
		rw.Write(page)
	})

	http.HandleFunc("/leaderboard", func(rw http.ResponseWriter, r *http.Request) {
		(&pkg.Response{
			Data: leaderboard.Data,
			Code: http.StatusOK,
		}).Write(rw)
	})

	http.HandleFunc("/roll", func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" || name == "undefined" {
			(&pkg.Response{
				Error: fmt.Errorf("please give a name"),
				Code:  http.StatusBadRequest,
			}).Write(rw)
			return
		}
		score := rand.Intn(1000000)
		leaderboard.Push(&pkg.User{
			Name:  name,
			Score: score,
		})

		(&pkg.Response{
			Data:  score,
			Error: nil,
			Code:  http.StatusOK,
		}).Write(rw)
	})

	log.Println("Start server at ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
