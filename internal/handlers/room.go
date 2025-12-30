package handlers

import(
	"fmt"
	"os"
	"time"
	w "videochat/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error{
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func Room(c *fiber.Ctx) error{
	uuid := c.Params("uuid")
	if uuid == ""{
		c.Status(400)
		return nil
	}
	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION"{
		ws = "wss"
	}
	uuid, suuid, _ := createOrGetRoom(uuid)
	return c.Render("peer", fiber.Map{
		"RoomWebSocketAddr":fmt.Sprintf("%s://%s/room/websocket/%s", ws, c.Hostname(), uuid),
		"RoomLink":fmt.Sprintf("%s://%s/room/%s", c.Protocol(), c.Hostname(), suuid),
		"ChatWebSocketAddr":fmt.Sprintf("%s://%s/chat/websocket/%s", ws, c.Hostname(), uuid),
		"StreamLink":fmt.Sprintf("%s://%s/stream/%s", c.Protocol(), c.Hostname(), suuid),
		"ViewerWebSocketAddr":fmt.Sprintf("%s://%s/room/viewer/websocket/%s", ws, c.Hostname(), uuid),
		"Type": "room",
	},	"layouts/main")
}

func RoomWebSocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == ""{
		return
	}
	_,_, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}

func createOrGetRoom(uuid string)(string, string, *w.Room){
	
}

func RoomViewerWebSocket(c *websocket.Conn){
	uuid := c.Pareams("uuid")
	if uuid==""{
		return
	}

	w.RoomsLock.Lock()
	if peer, ok := w.Rooms[uuid]; ok{
		roomViewerConn(c, peer.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func roomViewerConn(c *websocket.Conn, p *w.Peers){

}

type websocketMessage struct{
	Event string `json:"event"`
	Data  string    `json:"data"`
}