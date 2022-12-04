package cluster

import (
	"Go_redis/interface/resp"
	"Go_redis/resp/reply"
)

func flushDB(cluster *ClusterDatabase, conn resp.Connection, cmdArgs [][]byte) resp.Reply {
	replies := cluster.broadcast(conn, cmdArgs)
	var errReply reply.ErrorReply
	for _, r := range replies {
		if reply.IsErrReply(r) {
			errReply = r.(reply.ErrorReply)
			break
		}
	}
	if errReply != nil {
		return reply.MakeErrReply("error: " + errReply.Error())
	} else {
		return reply.MakeOkReply()
	}
}
