package chatroom

import (
	"net"
)

// 心跳集合
var heartbeats map[net.Conn]*Heartbeat
