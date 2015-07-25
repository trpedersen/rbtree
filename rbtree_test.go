package rbtree

import (
	"log"
	"math"
	"strconv"
	"testing"
	"time"
)

const (
	maxKeys int = 28870 //1000000
	loops   int = 1
)

var stringKey *StringKey = &StringKey{key: "stringKey", value: "stringValue"}
var intKey *IntKey = &IntKey{key: 2887, value: 2887}

type StringKey struct {
	key   string
	value string
}

func (this *StringKey) Compare(other Key) int {
	otherK := other.(*StringKey)
	if this.key == otherK.key {
		return 0
	} else if this.key < otherK.key {
		return -1
	} else {
		return 1
	}
}

type IntKey struct {
	key   int
	value int
}

func (this *IntKey) Compare(other Key) int {
	otherK := other.(*IntKey)
	if this.key == otherK.key {
		return 0
	} else if this.key < otherK.key {
		return -1
	} else {
		return 1
	}
}

func logElapsedTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func TestPutGetDeleteStringKey(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestPutGetDeleteStringKey")

	tree := NewRBTree()

	tree.Put(stringKey)

	key, err := tree.Get(&StringKey{key: stringKey.key}).(*StringKey)
	if !err {
		t.Error("tree.Get: invalid type")
	}
	if key != stringKey {
		t.Errorf("tree.Get: expected: %t, got: %t", stringKey, key)
	}

	tree.Delete(key)
	if tree.Size() != 0 {
		t.Errorf("tree.Delete, size wrong, expected: 0, got: %d", tree.Size())
	}
}

func TestPutGetDeleteStringKeys(t *testing.T) {

	f := func() {
		defer logElapsedTime(time.Now(), "TestPutGetDeleteStringKeys")
		tree := NewRBTree()

		for i := 0; i < maxKeys; i++ {
			s := strconv.Itoa(i)
			tree.Put(&StringKey{key: s, value: s})
		}

		for i := 0; i < maxKeys; i++ {
			s := strconv.Itoa(i)
			key, err := tree.Get(&StringKey{key: s}).(*StringKey)
			if !err {
				t.Error("invalid type")
			}
			if key.value != s {
				t.Error("tree.Get: expected: %s, got: %s", s, key.value)
			}
		}

		for i := 0; i < maxKeys; i++ {
			s := strconv.Itoa(i)
			tree.Delete(&StringKey{key: s})
			key := tree.Get(&StringKey{key: s})
			if key != nil {
				t.Errorf("found key when shouldn't have: %t", key)
			}

		}

		if tree.Size() != 0 {
			t.Error("tree size not zero: %d", tree.Size())
		}
	}

	for j := 0; j < loops; j++ {
		f()
	}
}

func TestPutIntKey(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestPutIntKey")

	tree := NewRBTree()

	tree.Put(intKey)

	key, err := tree.Get(&IntKey{key: intKey.key}).(*IntKey)
	if !err {
		t.Error("tree.Get: invalid type")
	}
	if key != intKey {
		t.Errorf("tree.Get: expected: %t, got: %t", intKey, key)
	}

	tree.Delete(key)
	if tree.Size() != 0 {
		t.Errorf("tree.Delete, size wrong, expected: 0, got: %d", tree.Size())
	}
}

func TestPutGetDeleteIntKeys(t *testing.T) {

	f := func() {
		defer logElapsedTime(time.Now(), "TestPutGetDeleteIntKeys")
		tree := NewRBTree()

		for i := 0; i < maxKeys; i++ {
			tree.Put(&IntKey{key: i, value: i})
		}

		for i := 0; i < maxKeys; i++ {
			key, err := tree.Get(&IntKey{key: i}).(*IntKey)
			if !err {
				t.Error("invalid type")
			}
			if key.value != i {
				t.Errorf("tree.Get: expected: %d, got: %d", i, key.value)
			}
		}

		for i := 0; i < maxKeys; i++ {

			tree.Delete(&IntKey{key: i})
			key := tree.Get(&IntKey{key: i})
			if key != nil {
				t.Errorf("found key when shouldn't have: %t", key)
			}

		}

		if tree.Size() != 0 {
			t.Error("tree size not zero: %d", tree.Size())
		}
	}

	for j := 0; j < loops; j++ {
		f()
	}
}

func TestContainsIntKeys(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestContainsIntKeys")
	tree := NewRBTree()

	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}
	for i := 0; i < maxKeys; i++ {
		if !tree.Contains(&IntKey{key: i}) {
			t.Errorf("contains: missing %d", i)
		}
	}
}

func TestContainsStringKeys(t *testing.T) {

	defer logElapsedTime(time.Now(), "TestContainsStringKeys")
	tree := NewRBTree()

	for i := 0; i < maxKeys; i++ {
		s := strconv.Itoa(i)
		tree.Put(&StringKey{key: s, value: s})
	}
	for i := 0; i < maxKeys; i++ {
		s := strconv.Itoa(i)
		if !tree.Contains(&StringKey{key: s}) {
			t.Errorf("contains: missing %s", s)
		}
	}
}

func TestSize(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestSize")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}
	if tree.Size() != maxKeys {
		t.Errorf("tree.Size(), expected: %d, got: %d", maxKeys, tree.Size())
	}
}

func TestHeight(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestHeight")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}
	maxHeight := math.Ilogb(float64(maxKeys))
	if tree.Height() > maxHeight {
		t.Errorf("tree.Height(), expected: %d, got: %d", maxHeight, tree.Height())
	}
}

func TestIsEmpty(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestIsEmpty")
	tree := NewRBTree()
	if !tree.IsEmpty() {
		t.Errorf("tree.IsEmpty(), expected: %t, got: %t", true, tree.IsEmpty())
	}
}

func TestFloorCeiling(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestFloorCeiling")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	mid := maxKeys / 2

	keyFloor, _ := tree.Floor(&IntKey{key: mid}).(*IntKey)
	keyCeiling, _ := tree.Ceiling(&IntKey{key: mid}).(*IntKey)
	if !((keyFloor.key <= mid) && (mid <= keyCeiling.key)) {
		t.Errorf("!(floor:%d <= mid:%d <= ceiling:%d)", keyFloor.key, mid, keyCeiling.key)
	}
}

func TestRankSelect(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestRankSelect")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	mid := maxKeys / 2

	key, _ := tree.Floor(&IntKey{key: mid}).(*IntKey)

	rank := tree.Rank(key)
	if rank != key.key {
		t.Errorf("tree.Rank: expected: %d, got: %d", key.key, rank)
	}

	selected, _ := tree.Select(rank).(*IntKey)
	if selected != key {
		t.Errorf("tree.Select: expected: %t, got: %t", key, selected)
	}

}

func TestMinMax(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestMinMax")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	min := tree.Min().(*IntKey).key
	if min != 0 {
		t.Errorf("tree.Min: expected: %d, got: %d", 0, min)
	}

	max := tree.Max().(*IntKey).key
	if max != maxKeys-1 {
		t.Errorf("tree.Max: expected: %d, got: %d", maxKeys-1, max)
	}
}

func TestDeleteMinMax(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestDeleteMinMax")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	tree.DeleteMin()
	min := tree.Min().(*IntKey).key
	if min != 1 {
		t.Errorf("tree.Min: expected: %d, got: %d", 1, min)
	}

	tree.DeleteMax()
	max := tree.Max().(*IntKey).key
	if max != maxKeys-2 {
		t.Errorf("tree.Max: expected: %d, got: %d", maxKeys-2, max)
	}
}

func TestKeysSlice(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestKeysSlice")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	keys := tree.Keys()
	if len(keys) != maxKeys {
		t.Errorf("tree.Keys, size invalid, expected: %d, got: %d", maxKeys, len(keys))
	}
	for i, _key := range keys {
		key := _key.(*IntKey)
		if i != key.key {
			t.Errorf("tree.Keys: invalid item, expected: %d, got: %d", i, key.key)
		}
		//log.Println(key.value)
	}
}

func TestKeysCh(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestKeysCh")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i++ {
		tree.Put(&IntKey{key: i, value: i})
	}

	quit := make(chan struct{})

	keys := tree.KeysCh(quit)

	var count int = 0
	var sum int = 0
	for _key := range keys {
		key := _key.(*IntKey)
		count++
		sum += key.value
		if count == 101 {
			break
		}
		//log.Println(key)
	}
	close(quit)
	if count != 101 { //maxKeys {
		t.Errorf("tree.KeysCh, count wrong, expected: %d, got: %d", maxKeys, count)
	}
	if sum != 5050 {
		t.Errorf("tree.KeysCh, sum wrong, expected: %d, got: %d", 5050, sum)
	}
	//log.Println(sum)
}

func TestKeysInRangeCh(t *testing.T) {
	defer logElapsedTime(time.Now(), "TestKeysInRangeCh")
	tree := NewRBTree()
	for i := 0; i < maxKeys; i += 2 {
		tree.Put(&IntKey{key: i, value: i})
	}

	quit := make(chan struct{})

	keys := tree.KeysInRangeCh(quit, &IntKey{key: 2000}, &IntKey{key: 3000})

	var count int = 0
	var sum int = 0
	for _key := range keys {
		key := _key.(*IntKey)
		log.Println(key)
		tree.Put(&IntKey{key: key.key + 1, value: key.value * 10})
		count++
		sum += key.value
		if count == 101 {
			break
		}
	}
	close(quit)
	if count != 101 { //maxKeys {
		t.Errorf("tree.KeysInRangeCh, count wrong, expected: %d, got: %d", maxKeys, count)
	}
	//log.Println(sum)
}
