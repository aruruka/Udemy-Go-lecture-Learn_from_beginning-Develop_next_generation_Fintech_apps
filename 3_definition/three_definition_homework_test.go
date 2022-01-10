package three_definition

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Q1. 以下の1.11をint型に変換して出力してください。

f := 1.11

Q2. コードを書かずに以下の出力結果を答えてください。

s := []int{1, 2, 5, 6, 2, 3, 1}
fmt.Println(s[2:4])

R: [5 6]

Q3. 以下のコードを実行した時に
fmt.Printf("%T %v", m, m)

以下のような出力結果となるmを作成してください。

map[string]int map[Mike:20 Nancy:24 Messi:30]
*/

func TestThreeDefinitionHomework(t *testing.T) {
	f := 1.11
	convertedF := int(f)
	fmt.Printf("convertedF's type is: %v\n", reflect.TypeOf(convertedF))
	fmt.Printf("convertedF's value is: %v\n", convertedF)

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}

/* Q1. 解答例

package main

import (
    "fmt"
)

func main() {
    f := 1.11
    i := int(f)
    fmt.Printf("%T %v\n", i, i)
}
Q2. 解答　[5 6]

package main

import (
    "fmt"
)

func main() {
    s := []int{1, 2, 5, 6, 2, 3, 1}
    fmt.Println(s[2:4])
}
Q3. 解答例

package main

import (
    "fmt"
)

func main() {
    m := map[string]int{
        "Mike": 20,
        "Nancy": 24,
        "Messi": 30,
    }
    fmt.Printf("%T %v", m, m)
} */
