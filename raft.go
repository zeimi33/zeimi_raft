package zeimi_raft

import (
	"errors"
	"math"
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

type Config struct {
	//自己的id
	ID uint64
	//组员
	peers []uint64
	//选举时间，如果一个follow在这段时间没有收到心跳，他将发起选举
	ElectionTick int
	//心跳
	HeartBeatTick int
	//raft落盘的地方
	Storage Storage
	//最后一个应用的日志index
	Applied uint64
	//日志记录
	Logger Logger
	//是否该检测自己的领导地位，如果没超过一半，leader下台
	CheckQuorum bool
	//消息的最大值
	MaxSizePerMsg uint64
	//最大的可以应用的，已经提交的日志数量
	MaxCommittedSizePerReady uint64
	//最大的没有提交的条目数量
	MaxUncommittedEntriesSize uint64
	//最大的追加条目数量，防止发送缓存区溢出
	MaxInflightMsgs int
}

func (c *Config) valid() error {
	if c.ID == 0 {
		return errors.New("can't use 0 as id")
	}
	if c.HeartBeatTick <= 0 {
		return errors.New("heart beat should >= 0")
	}
	if c.ElectionTick <= c.HeartBeatTick {
		return errors.New("ElectionTick >= HeartBeatTick ")
	}
	if c.Storage == nil {
		return errors.New("storage can't be nil")
	}
	if c.MaxUncommittedEntriesSize == 0 {
		c.MaxUncommittedEntriesSize = math.MaxUint64
	}
	if c.MaxCommittedSizePerReady == 0 {
		c.MaxCommittedSizePerReady = c.MaxSizePerMsg
	}
	if c.MaxInflightMsgs <= 0 {
		return errors.New("max inflight messages must be greater than 0")
	}
	if c.Logger == nil {
		c.Logger = NewRaftLogger()
	}
	return nil
}

func newRaft(c *Config) *raft {
	if err := c.valid(); err != nil {
		panic(err)
	}
	newLogWithSize(c.Storage, c.Logger, c.MaxCommittedSizePerReady)
	_, cs, err := c.Storage.InitialState()
	if err != nil {
		panic(err)
	}
	cs.Voters = c.peers
	r := &raft{}
	return r
}
