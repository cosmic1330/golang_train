package main

import "fmt"

/* type1 */
type TextMes struct {
	Text string
	Type string
}

func (t *TextMes) setText(s string) {
	t.Text = s
}

func (t *TextMes) setType(input string) {
	t.Type = input
}

/* type2 */
type ImageMes struct {
	Img  string
	Type string
}

func (i *ImageMes) setImage(s string) {
	i.Img = s
}

func (i *ImageMes) setType(input string) {
	i.Type = input
}

/* interface */
type Messenger interface {
	setType(input string)
}

func send(m Messenger) {
	m.setType("sendSameType")
	fmt.Println("m", m)
}

// 類型選擇 可調用其他未定義的方法
func sendOthers(m Messenger) {
	switch v := m.(type) {
	case *TextMes:
		v.setText("sendOthersText")
	case *ImageMes:
		v.setImage("sendOthersImage")
	default:
		fmt.Println("default", v)
	}
	fmt.Println("m", m)
}

func main() {
	// call 已定義的方法
	tm := &TextMes{}
	send(tm)

	im := &ImageMes{}
	send(im)

	// call 未定義的方法
	sendOthers(tm)
	sendOthers(im)

	// 空接口 可以接受任何類型
	// 類型斷言 可以將空接口轉換為指定類型
	value := 1
	i := interface{}(value)
	iv, ok := i.(int)
	if ok {
		fmt.Println("i value", iv, ok)
	} else {
		fmt.Println("類型斷言失敗")
	}

	// 当一个接口变量没有被赋值时，它的值就是 nil
	var i2 interface{}
	fmt.Println("i2", i2)
	i2 = 1 // 這樣才能使用
	fmt.Println("i2", i2)
}
