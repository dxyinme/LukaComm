package protocol

import "time"

const (
	MaxMsgLength = 40000
	KCPPort = 38306

	SocketLength = 4096 // 4KB , length of socket

	ServerMessageChannelLength = 1000 // the length of server channel

	ConfirmCode = "OK_"

	UsualKCPLiveTime = time.Second
)

