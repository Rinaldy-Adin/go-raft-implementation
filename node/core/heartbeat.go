package core

import (
	"context"

	pb "tubes.sister/raft/gRPC/node/core"
)

func (rn *RaftNode) aptries(followerIdx int) error {
	ctx, cancel := context.WithTimeout(context.Background(), HEARTBEAT_SEND_INTERVAL)
	defer cancel()

	prevLogIndex := int32(rn.Volatile.ClusterList[followerIdx].NextIndex - 1)
	prevLogTerm := int32(0)
	if prevLogIndex >= 0 {
		prevLogTerm = int32(rn.Persistence.Log[prevLogIndex].Term)
	}

	args := &pb.AppendEntriesArgs{
		Term: int32(rn.Persistence.CurrentTerm),
		LeaderAddress: &pb.AppendEntriesArgs_LeaderAddress{
			Ip:   rn.Address.IP,
			Port: int32(rn.Address.Port),
		},
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  int32(prevLogTerm),
		Entries:      []*pb.AppendEntriesArgs_LogEntry{},
		LeaderCommit: int32(rn.Volatile.CommitIndex),
	}

	rn.Volatile.ClusterList[followerIdx].Client.AppendEntries(ctx, args)

	// if err != nil {
	// 	log.Printf("Failed to call AppendEntries: %v", err)
	// }

	return nil
}

func (rn *RaftNode) heartbeat() {
	for idx, clusterData := range rn.Volatile.ClusterList {
		if !clusterData.Address.Equals(&rn.Address) {
			go rn.aptries(idx)
		}
	}
}
