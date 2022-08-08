/**
@author ZhengHao Lou
Date    2022/06/29
*/
package main

import (
	"encoding/base64"
	"fmt"
)

/**
https://leetcode.cn/problems/encode-and-decode-tinyurl/
*/
type Codec struct {
	i     int
	cache map[string]string
}

func Constructor() Codec {
	return Codec{
		cache: map[string]string{},
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	e := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", this.i)))
	this.i++
	this.cache[e] = longUrl
	return e
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.cache[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
