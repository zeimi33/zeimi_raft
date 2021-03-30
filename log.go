package zeimi_raft

type raftLog struct {
	//已经提交的日志
	committed uint64

	applied uint64
	logger  Logger
}
