package zeimi_raft

import "crypto/sha1"

type NodeID struct {
	IP string
	Port string
	Hash string
}

type NodeGroup struct {
	self *NodeID
	Map map[string]*NodeID
}


func NewNodeID (IP string,Port string)*NodeID{
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(IP+Port))
	hash := Sha1Inst.Sum([]byte(""))
	n := NodeID{
		IP: IP,
		Port: Port,
		Hash: string(hash),
	}
	return &n
}

func (n *NodeID)equals(id NodeID)bool {
	return n.IP == id.IP && n.Port == id.Port
}

func (n *NodeID)GetHash()string{
	return n.Hash
}

func(n *NodeID)ConstructNodeGroup(Nodes ...*NodeID)*NodeGroup{
	NodeGrp := &NodeGroup{
		self: n,
		Map: make(map[string]*NodeID),
	}
	for _, iter := range Nodes{
		NodeGrp.Map[iter.Hash] = iter
	}
	return NodeGrp
}

func(n *NodeGroup)GetMember(hash string)*NodeID{
	return  n.Map[hash]
}