package main

import (
	rbtree "github.com/trpedersen/algorithms/rbtree"
	// "fmt"
	"log"
	"os"
	"runtime/pprof"
	// "strconv"
	// "testing"
	"flag"
	"time"
)

const (
	trials int = 1000000
)

type StringKey struct {
	key   string
	value string
}

func (this *StringKey) Compare(other rbtree.Key) int {
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

func (this *IntKey) Compare(other rbtree.Key) int {
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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	TestDeleteIntKeys()
}

// func TestPutStringKey(t *testing.T) {

// 	defer logElapsedTime(time.Now(), "TestPutStringKey")

// 	tree := rbtree.NewRBTree()

// 	tree.Put(&StringKey{key: "a", value: "a_value"})

// 	key, err := tree.Get(&StringKey{key: "a"}).(*StringKey)
// 	if !err {
// 		t.Error(err)
// 	}
// 	if key.value != "a_value" {
// 		t.Error(key.value)
// 	}
// 	fmt.Println(key)

// }

// func TestPutStringKeys(t *testing.T) {

// 	f := func() {
// 		defer logElapsedTime(time.Now(), "TestPutStringKeys")
// 		tree := rbtree.NewRBTree()

// 		for i := 0; i < trials; i++ {

// 			s := strconv.Itoa(i)
// 			tree.Put(&StringKey{key: s, value: s})

// 			key, err := tree.Get(&StringKey{key: s}).(*StringKey)
// 			if !err {
// 				t.Error(err)
// 			}
// 			if key.value != s {
// 				t.Error(key)
// 			}
// 			//fmt.Println(key)
// 		}
// 	}

// 	for j := 0; j < 1; j++ {
// 		f()
// 	}

// }

// func TestPutIntKey(t *testing.T) {

// 	defer logElapsedTime(time.Now(), "TestPutIntKey")

// 	tree := rbtree.NewRBTree()

// 	tree.Put(&IntKey{key: 1, value: 1})

// 	key, err := tree.Get(&IntKey{key: 1}).(*IntKey)
// 	if !err {
// 		t.Error(err)
// 	}
// 	if key.value != 1 {
// 		t.Error(key.value)
// 	}
// 	log.Println(key)

// }

// func TestPutIntKeys(t *testing.T) {

// 	f := func() {
// 		defer logElapsedTime(time.Now(), "TestPutIntKeys")
// 		tree := rbtree.NewRBTree()

// 		for i := 0; i < trials; i++ {

// 			tree.Put(&IntKey{key: i, value: i})

// 			key, err := tree.Get(&IntKey{key: i}).(*IntKey)
// 			if !err {
// 				t.Error(err)
// 			}
// 			if key.value != i {
// 				t.Error(key.value)
// 			}
// 			//fmt.Println(key)
// 		}
// 	}

// 	for j := 0; j < 1; j++ {
// 		f()
// 	}

// 	//for i

// 	//log.Printf("tree.Size(): %d\n", tree.Size())
// }

func TestDeleteIntKeys() {

	f := func() {
		defer logElapsedTime(time.Now(), "TestDeleteIntKeys")
		tree := rbtree.NewRBTree()

		for i := 0; i < trials; i++ {

			tree.Put(&IntKey{key: i, value: i})

			key, err := tree.Get(&IntKey{key: i}).(*IntKey)
			if !err {
				log.Println(err)
			}
			if key.value != i {
				log.Println(key.value)
			}
		}
		for i := 0; i < trials; i++ {

			tree.Delete(&IntKey{key: i})
			key := tree.Get(&IntKey{key: i})
			if key != nil {
				//key, err := .(*IntKey)
				//if !err {
				log.Println(key)
			}
			// if key != nil {
			// 	t.Error(key.value)
			// }
		}
		if tree.Size() != 0 {
			log.Println("tree size")
		}
	}

	for j := 0; j < 1; j++ {
		f()
	}

	//for i

	//log.Printf("tree.Size(): %d\n", tree.Size())
}
