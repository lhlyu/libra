package module

import "sort"

type Module interface {
	SetUp()
	seq() int // 模块执行顺序
}

type module struct {
	mi  Module
	seq int
}

type modules []*module

var mods []*module

// 自定义排序
func (ms modules) Len() int {
	return len(ms)
}

func (ms modules) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// 升序
func (ms modules) Less(i, j int) bool {
	return ms[i].seq < ms[j].seq
}

func Register(mi ...Module) {
	for _, m := range mi {
		mod := new(module)
		mod.mi = m
		mod.seq = m.seq()
		mods = append(mods, mod)
	}
	sort.Sort(modules(mods))
}

func Init() {
	for _, m := range mods {
		m.mi.SetUp()
	}
}
