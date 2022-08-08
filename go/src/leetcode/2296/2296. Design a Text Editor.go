/**
@author ZhengHao Lou
Date    2022/06/06
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/design-a-text-editor/
*/
type TextEditor struct {
	texts  []byte
	cursor int
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
	n := len(text)
	tmp := make([]byte, n)
	this.texts = append(this.texts, tmp...)
	copy(this.texts[this.cursor+n:], this.texts[this.cursor:])
	copy(this.texts[this.cursor:this.cursor+n], []byte(text))
	this.cursor += n
}

func (this *TextEditor) DeleteText(k int) int {
	n := min(this.cursor, k)
	copy(this.texts[this.cursor-n:], this.texts[this.cursor:])
	this.texts = this.texts[:len(this.texts)-n]
	this.cursor -= n
	return n
}

func (this *TextEditor) CursorLeft(k int) string {
	for this.cursor > 0 && k != 0 {
		this.cursor--
		k--
	}
	return string(this.texts[max(0, this.cursor-10):this.cursor])
}

func (this *TextEditor) CursorRight(k int) string {
	for this.cursor < len(this.texts) && k != 0 {
		this.cursor++
		k--
	}
	return string(this.texts[max(0, this.cursor-10):this.cursor])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	textEditor := Constructor()
	textEditor.AddText("leetcode")
	fmt.Println(textEditor.DeleteText(4))
	textEditor.AddText("practice")
	fmt.Println(textEditor.CursorRight(3))
	fmt.Println(textEditor.CursorLeft(8))
	fmt.Println(textEditor.DeleteText(10))
	fmt.Println(textEditor.CursorLeft(2))
	fmt.Println(textEditor.CursorRight(6))
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
