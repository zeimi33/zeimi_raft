package zeimi_raft

import "github.com/zeimi33/zeimi_raft/raftpb"

type raftLog struct {
	//已经提交的日志
	committed       uint64
	storage         Storage
	applied         uint64
	logger          Logger
	maxNextEntsSize uint64
	unstable        unstable
}

func newLogWithSize(storage Storage, logger Logger, maxNextEntSize uint64) *raftLog {
	l := &raftLog{
		storage:         storage,
		logger:          logger,
		maxNextEntsSize: maxNextEntSize,
	}
	firstIndex, err := storage.FirstIndex()
	if err != nil {
		panic(err)
	}
	lastIndex, err := storage.LastIndex()
	if err != nil {
		panic(err)
	}
	l.logger = logger
	//指向最后一个节点
	l.unstable.offset = lastIndex + 1
	l.committed = firstIndex - 1
	l.applied = firstIndex - 1
	return l
}

type unstable struct {
	snapshot raftpb.Snapshot
	entries  []raftpb.Entry
	offset   uint64
	Logger   Logger
}
