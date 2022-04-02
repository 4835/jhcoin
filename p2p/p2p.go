package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/6233/jhcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}