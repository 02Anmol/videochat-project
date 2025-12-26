package handlers

import(
	"videochat/pkg/chat"
	w "videochat/pkg/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RoomChat(c *fiber.Ctx) error{
	returnn c.Render("chat", fiber.Map{}, "layouts/main")
}

func RoomChatWebSocket(c *websocket.Conn){
	uuid := c.Params("uuid")
	if uuid == ""{
		return
	}
	w.RoomsLock.Lock()
	room := w.Rooms[uuid]
	w.RoomsLock.Unlock()
	if room == nil{
		return
	}
	if room.Hub == nil{
		return
	}
	chat.PeerChatConn(c, room.Hub)
}

func StreamChatWebSocket(c *websocket.Conn){
	ssuid := c.Params("ssuid")
	if ssuid == ""{
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[ssuid]; ok{
		if stream.Hub == nil{
			hub := chat.NewHub()
			stream.Hub = hub
			go hub.Run()
		}
		chat.PeerChatConn(c, stream.Hub)
		return
	}
	w.RoomsLock.Unlock()

}