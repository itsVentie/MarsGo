package transposition

type EntryType int

const (
	Exact EntryType = iota
	LowerBound
	UpperBound
)

type Entry struct {
	Value int
	Depth int
	Type  EntryType
}

type Table struct {
	data map[uint64]Entry
}

func NewTable() *Table {
	return &Table{data: make(map[uint64]Entry)}
}

func (t *Table) Get(hash uint64) (Entry, bool) {
	e, ok := t.data[hash]
	return e, ok
}

func (t *Table) Store(hash uint64, entry Entry) {
	t.data[hash] = entry
}
