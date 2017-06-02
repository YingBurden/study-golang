package cfg2json

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Object = "Object"
	Int    = "Int"
	String = "String"
	Array  = "Array"
	Init   = "Init"
)

type jsonstruct struct {
	prv       *jsonstruct
	next      *jsonstruct
	child     *jsonstruct
	father    *jsonstruct
	key       string
	value     string
	valueType string
}

func NewJson(filepath string) *jsonstruct {
	f, _ := os.Open(filepath)
	r := bufio.NewReader(f)

	node := new(jsonstruct)
	initdata := node

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		s := strings.TrimSpace(string(b))
		if len(s) == 0 {
			continue
		}
		if strings.Index(s, "#") == 0 {
			continue
		}
		if s == "}" {

			firstnode := first(node)

			if firstnode.father.key == "" {
				node = first(firstnode.father).father

			} else {
				node = firstnode.father

			}

		} else {
			spaceNum := strings.Count(s, " ")
			if spaceNum > 1 {
				strings.Replace(s, " ", "", spaceNum-1)
			}
			slice := strings.Split(s, " ")
			var status bool
			node, status = checkSlice(node, slice[0])
			if slice[1] == "{" {
				newnode := new(jsonstruct)
				newnode.valueType = Object
				newnode.prv = node

				//若同级别node并没有相同key
				if !status {
					newnode.key = slice[0]
				}

				childnode := new(jsonstruct)
				childnode.father = newnode
				childnode.valueType = Init
				childnode.key = Init
				newnode.child = childnode
				node.next = newnode
				node = childnode

			} else {
				nextnode := new(jsonstruct)
				if node.valueType == Init {
					nextnode = node
				} else {
					node.next = nextnode
					nextnode.prv = node
				}
				nextnode.valueType = String
				nextnode.key = slice[0]
				nextnode.value = slice[1]
				node = nextnode

			}
		}
		node = final(node)

	}
	return initdata

}

func checkSlice(node *jsonstruct, key string) (*jsonstruct, bool) {
	initNode := node
	status := false
	for node.prv != nil {
		if node.key == key {
			if node.valueType == Object {
				newNode := new(jsonstruct)
				newNode.next = node.next
				newNode.prv = node.prv
				node.prv.next = newNode
				newNode.key = node.key
				newNode.valueType = Array
				newNode.father = node.father
				newNode.child = node

				node.father = newNode
				node.key = ""
				node.prv = nil

				status = true
				return node, status

			}

			if node.valueType == Array {

				prvNode := final(node.child)

				status = true
				return prvNode, status
			}

			if node.valueType == String {
				// 暂时不考虑这个环节
			}
		}
		node = node.prv
	}
	return initNode, status
}

func final(node *jsonstruct) *jsonstruct {
	for node.next != nil {
		node = node.next
	}
	return node
}

func first(node *jsonstruct) *jsonstruct {
	for node.prv != nil {
		//		fmt.Println(*node)
		node = node.prv
	}
	return node
}

func start(node *jsonstruct) *jsonstruct {
	for {
		if node.prv != nil {
			node = node.prv
		} else if node.father != nil {
			node = node.father
		} else {
			break
		}
	}
	return node
}
