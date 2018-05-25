package main

type Table struct {
	array []*tableItem
}

type tableItem struct {
	key   interface{}
	value interface{}
	table *Table
}

//创建指定大小的哈希表

func New(size int) *Table {
	return &Table{
		array: make([]*tableItem, size),
	}
}

//设置键值对数据

func (t *Table) Set(key int, value interface{}) {
	hash := key % len(t.array)
	if t.array[hash] != nil {
		if t.array[hash].key != key {
			if t.array[hash].table == nil {
				t.array[hash].table = New(len(t.array))

			}
			t.array[hash].table.Set(key, value)
		} else {
			t.array[hash].value = value
		}

	} else {
		t.array[hash] = &tableItem{
			key:   key,
			value: value,
		}
	}

}

//根据键名查询键值

func (t *Table) Get(key int) interface{} {
	hash := key % len(t.array)
	if t.array[hash] != nil {
		if t.array[hash].key != key {
			return t.array[hash].table.Get(key)
		} else {
			return t.array[hash].value
		}
	}
}
