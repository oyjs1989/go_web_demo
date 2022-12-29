package consensus

import (
	"log"
	"time"
)

const TERM_TIME = 10

type Node struct {
	log       string // 日志
	term_no   int    // 当前任期
	term_time int    // term任期的时间
}

// 记时
func (n *Node) run() {
	for true {
		log.Printf("node %d run", n.term_no)
		if n.term_time > 0 { // 如果term时间大于0
			n.term_time-- // 时间减1
		} else { // 否则
			n.term_no++                    // 任期加1
			n.term_time = TERM_TIME * 1000 // 任期时间为10
		}
		time.Sleep(1 * time.Millisecond) // 睡眠1秒
	}
}

func (n *Node) receivePrepare(term int) bool {
	// 收到提案
	// 如果提案的任期大于当前任期
	// 则更新当前任期
	// 并且重置term时间
	if term > n.term_no {
		n.term_no = term
		n.term_time = TERM_TIME * 1000
		return true
	}
	return false
}

func (n *Node) sendPrepare() int {
	// 发送提案
	return n.term_no
}

func (n *Node) sendAccept() {
	// 发送结果
	// 给所有的node节点发送提案
	// 如果有超过半数的节点接受了提案
	// 则发送决议

	// 如果没有超过半数的节点接受了提案
	// 则重新发送提案

}

func receiveAccept() {
	// 收到提案
	// 如果提案的任期大于当前任期
	// 则更新当前任期
	// 并且重置term时间
}

func receive() {

}
