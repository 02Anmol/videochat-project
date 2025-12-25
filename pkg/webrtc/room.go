package webrtc

import()

func RoomConn(c *websocket.Conn, peers *Peers){
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnections,
		webSocket: &ThreadSafeWriter{},
		Conn: c,
		Mutex: &sync.Mutex{},
	}
}