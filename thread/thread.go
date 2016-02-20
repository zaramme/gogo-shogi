package thread

import (
	c "github.com/zaramme/gogo-shogi/color"
	mv "github.com/zaramme/gogo-shogi/move"
	po "github.com/zaramme/gogo-shogi/position"
	sc "github.com/zaramme/gogo-shogi/score"
	tt "github.com/zaramme/gogo-shogi/tt"
	"sync"
)

const (
	MaxThreads             = 64
	MasSplitPointPerThread = 8
)

type NodeType int

const (
	Root            NodeType
	PV              NodeType
	NonPV           NodeType
	SplitPointRoot  NodeType
	SplitPointPV    NodeType
	SplitPointNonPV NodeType
)

// 時間や探索深さの制限を格納する為の構造体
type LimitsType struct {
	Time      [c.ColorNum]int
	Increment [c.ColorNum]int
	MovesToFo int
	Depth     sc.Ply
	Nodes     uint32
	MoveTime  int
	Infinite  bool
	Ponder    bool
}

type SplitPoint struct {
	pos *po.Position
	//	ss *SearchStack
	masterThread Thread
	Depth        tt.Depth
	beta         sc.Score
	NodeType     NodeType
	Move         mv.Move
	cutNode      bool

	//	MovePicker *MovePicker
	// ParentSplitPoint *SplitPoint
	mutex sync.Mutex

	SlavesMask uint64 // volatile
	Nodes      int64
	Alpha      sc.Score
	BestScore  sc.Score
	BestMove   mv.Move
	MoveCount  int
	CutOff     bool
}

// スレッド共通で持たせたい処理をインターフェース化
type LoopWorker interface {
	IdleLoop()
}

type Thread struct {
	SplitPoints    [MasSplitPointPerThread]SplitPoint
	ActivePosition *po.Position
	Idx            int
	MaxPly         int
	SleepLock      sync.Mutex
	// sleepCond std::condition_variable -> goChannelで処理する
	// handle std::thread -> 不要
	SplitPointsSize int  // volatile
	Searching       bool // valatile
	// Searchar *Searchar
}

type MainThread struct {
	Thread
	Thinking bool
}

type TimerThread struct {
	Thread
	Msec int
}

type ThreadPool struct {
	Depth                   tt.Depth
	sleepWhileIdele         bool
	maxThreadsPerSplitPoint int
	mutex                   sync.Mutex
	// std::condition_variable sleepCond_; -> goChannelで処理する
	/////////////////
	//private
	timerThread       *TimerThread
	minimumSplitDepth tt.Depth
}
