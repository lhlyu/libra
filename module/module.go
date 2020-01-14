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

var mods []*module

func Register(mi ...Module) {
	for _, m := range mi {
		mod := new(module)
		mod.mi = m
		mod.seq = m.seq()
		mods = append(mods, mod)
	}
	// 排序
	sort.SliceStable(mods, func(i, j int) bool {
        return mods[i].seq < mods[j].seq
    })
}

func Init() {
	for _, m := range mods {
		m.mi.SetUp()
	}
}
