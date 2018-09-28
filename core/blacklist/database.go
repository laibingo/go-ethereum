package blacklist

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"sync"
)

var bldb *Blacklist

func init() {
	bldb = &Blacklist{
		CurHeight: 0,
		M:         make(map[common.Address]int64),
	}
}

type Database interface {
	OpenBlacklist(root common.Hash, height int64) *Blacklist
	//CommitBlacklist(blacklist *Blacklist, hash common.Hash)
}

func NewDatabase(db ethdb.Database) Database {
	return &myDB{
		db: db,
		m:  make(map[common.Hash]*Blacklist),
	}
}

type myDB struct {
	db ethdb.Database
	m  map[common.Hash]*Blacklist
	mu sync.Mutex
}

func (m *myDB) OpenBlacklist(root common.Hash, height int64) *Blacklist {
	bldb.CurHeight = height
	return bldb
	//TODO
	/*
	m.mu.Lock()
	defer m.mu.Unlock()
	root = common.HexToHash("0x00");
	if v, ok := m.m[root]; ok {
		return v
	} else {
		k := key(root)
		fmt.Println("lk:" + root.Hex())
		fmt.Println(k)
		v, err := m.db.Get(k)
		b := &Blacklist{}
		if err == nil && len(v) > 0 {
			json.Unmarshal(v, b)
			m.m[root] = b
		} else {
			b = &Blacklist{
				M: make(map[common.Address]int64),
			}
		}
		b.CurHeight = height
		m.m[root] = b
		return b;
	}
	*/
}

/*
func (m *myDB) CommitBlacklist(blacklist *Blacklist, hash common.Hash) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[hash] = blacklist
	//TODO
	hash = common.HexToHash("0x00");
	bytes, err := json.Marshal(*blacklist)
	fmt.Println("pk:" + hash.Hex())
	if err == nil {
		m.db.Put(key(hash), bytes)
	}
}
*/

func key(hash common.Hash) []byte {
	return []byte(hash.Hex() + ":blacklist")
}
