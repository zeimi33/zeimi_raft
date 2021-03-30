package zeimi_raft

import (
	"math/rand"
	"sync"
	"time"
)

type raft struct {
	id               uint64
	Term             uint64
	Vote             uint64
	leaderID         uint64
	heartBeatTimeOut int
	electionTimeOut  int
	state            int
	//为了避免多数节点同一时间选举，产生随机值，使选举更快
	randomizedElectionTimeOut int
	//在不同的状态执行不同的函数（虚函数）
	step stepFunc
	//记录log
	logger Logger
	//raft log
	log *raftLog
}

func (r *raft) SetTerm(term uint64) {
	r.Term = term
	r.Vote = 0
	r.leaderID = 0
}

type Message string
type stepFunc func(r *raft, m Message) error

func (r *raft) becomeFollower(term uint64, lead uint64) {
	r.step = stepFollower
	r.SetTerm(term)
	r.leaderID = lead
	r.state = StateFollower
	r.logger.Info(r.id, "become follower,the leader is", r.leaderID)

}

func (r *raft) becomeCandidate() {
	if r.state == StateLeader {
		panic("leader can't become candidate")
	}
	r.step = stepFCandidate
	r.SetTerm(r.Term + 1)
	r.Vote = r.id
	r.state = StateCandidate
	r.logger.Infof("%x become candidate at term %d", r.id, r.Term)
}

func (r *raft) becomeLeader() {
	if r.state == StateFollower {
		panic("follower can't become leader")
	}
	r.step = stepLeader
	r.SetTerm(r.Term)
	r.leaderID = r.id
	r.state = StateLeader
}

func stepFollower(r *raft, m Message) error {
	return nil
}

func stepFCandidate(r *raft, m Message) error {
	return nil
}

func stepLeader(r *raft, m Message) error {
	return nil
}

type LockRander struct {
	Lock sync.Mutex
	rand *rand.Rand
}

var GlobalRander = &LockRander{
	rand: rand.New(rand.NewSource(time.Now().Unix())),
}

func (r *LockRander) Init(i int) int {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	return r.rand.Intn(i)
}
