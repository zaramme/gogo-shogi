package tt

import (
	"github.com/zaramme/gogo-shogi/common"
	mv "github.com/zaramme/gogo-shogi/move"
	sc "github.com/zaramme/gogo-shogi/score"
	"unsafe"
)

// sizeof(TTEntry)をunsafe.SizeOf(TTEntry{})で代替。正しいかどうか不明
// ちなみにunsafe.SizeOf(TTEntry{}) = 16(byte)
const (
	ClusterSize = common.CacheLineSize / unsafe.Sizeof(TTEntry{})
)

type TTCluster [ClusterSize]TTEntry

const (
	TTClusterSize = unsafe.Sizeof(TTCluster{})
)

type TranspositionTable struct {
	Size       int // 置換表のバイト数。2のべき乗である必要がある。
	entries    *TTCluster
	generation uint8 // iterative deepening していくとき、過去の探索で調べたものかを判定する。
}

func (tt *TranspositionTable) setSize(mbSize int) {
	// 確保する要素数を取得する。
	newSize := mbSize << 20 / int(TTClusterSize)
	if newSize < 1024 {
		newSize = 1024 // 最小値は 1024 としておく。
	}
	msbIndex := uint(63 - common.FirstOneFromMSB(uint64(newSize)))
	newSize = 1 << msbIndex

	if newSize == tt.Size {
		// 現在と同じサイズなら何も変更する必要がない。
		return
	}

	tt.Size = newSize
	tt.entries = new(TTCluster)

	// if (!entries_) {
	// 	std::cerr << "Failed to allocate transposition table: " << mbSize << "MB";
	// 	exit(EXIT_FAILURE);
	// }
	// clear();
}

func (tt *TranspositionTable) clear() {
	tt.entries = new(TTCluster) // ←怪しいような。
	// ↓かといって、こんなことするのもあれだし。
	// for i := 0; i < ClusterSize; i++ {
	// 	ClusterSize[i] = TTEntry{}
	// }

	//memset(entries_, 0, size() * sizeof(TTCluster));
}

func (tt *TranspositionTable) FirstEntry(posKey common.Key) *TTEntry {
	return &tt.entries[int(posKey)&(len(tt.entries)-1)]
}

func (tt *TranspositionTable) store(
	posKey common.Key, score sc.Score, bound sc.Bound,
	depth Depth, mv mv.Move, evalScore sc.Score) {
	tte := tt.FirstEntry(posKey)
	replace := tte

	posKeyHigh32 := u32int(posKey >> 32)

	if depth < Depth0 {
		depth = Depth0
	}

	for i := 0; i < ClusterSize; i++ {
		// // 置換表が空か、keyが同じな古い情報が入っているとき
		// if (!tte->key() || tte->key() == posKeyHigh32) {
		// 	// move が無いなら、とりあえず古い情報でも良いので、他の指し手を保存する。
		// 	if (move.isNone()) {
		// 		move = tte->move();
		// 	}

		// 	tte->save(depth, score, move, posKeyHigh32,
		// 			  bound, this->generation(), evalScore);
		// 	return;
		// }

		// int c = (replace->generation() == this->generation() ? 2 : 0);
		// c    += (tte->generation() == this->generation() || tte->type() == BoundExact ? -2 : 0);
		// c    += (tte->depth() < replace->depth() ? 1 : 0);

		// if (0 < c) {
		// 	replace = tte;
		// }
	}
}

func (tt *TranspositionTable) probe(posKey common.Key) *TTEntry {
	posKeyHigh32 = posKey >> 32
	tte := tt.FirstEntry(posKey)

	for i := 0; i < ClusterSize; i++ {
		// if (tte->key() == posKeyHigh32) {
		// 	return tte;
		// }
	}

	return nil
}

// 元はinline関数なので場合によってはベタ書きする
func (tt *TranspositionTable) NewSearch() {
	tt.generation++
}

// 元はinline関数なので場合によってはベタ書きする
func (tt TranspositionTable) Refresh(tte *TTEntry) {
	tte.SetGeneration(tt.generation)
}
