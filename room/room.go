package room

var (
	DefaultRoom *ChatRoom
)

func init() {
	DefaultRoom = NewChatRoom()
}

type ChatRoom struct {
	clients   map[*Client]bool
	broadcast chan []byte
}

// NewChatRoom
// 新建房间
func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
	}
}

func (cr *ChatRoom) Run() {
	for {
		select {
		case message := <-cr.broadcast:
			for client := range cr.clients {
				client.ch <- message
			}
		}
	}
}

func GetOnlineUser() int {
	count := 0
	if DefaultRoom.clients == nil {
		return 0
	} else {
		for _, value := range DefaultRoom.clients {
			if value {
				count++
			}
		}
	}
	return count
}
