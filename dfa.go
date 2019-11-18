package Dfa

import (
	"fmt"
	"strings"
)

// the keymap  is a map
// every one  in  keywords  buile as a key in this map
//exp :   "XBCDEF" --> map{"X":map{"B":map{"C":map{"D":map{"E":map{"F":map{"\x00":"XBCDEF" }}}}}}}
// all  to lowercase  if the key is ascii .
type Keyword map[string]interface{}

// give  a string  to find  whether  it contains  a  keyword  build  into keymap .
//  only  get the first Keyword  in keymap then return.
//exp :
//  Buildkeymap(set, "abc")
//  Search (set ,"abcsdabcfdsfasdfa")  --> (true "abc") get the first "abc"

func Search (set Keyword,str string) (find bool,  key string) {
	slc :=make([]string,0)
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	for _,v:=range str {
		slc=append(slc,string(v))
	}
	keymap := make(Keyword)
	keymap = set
	keylen :=len(slc)-1
	for ix,k :=range slc {
		//fmt.Println(ix,"k:",k)
		val,ok := keymap[k]
		//fmt.Print("val:",val,"\n")
		if ok {
			keymap = val.(Keyword)
			if v,ok :=keymap["\x00"]; ok {
				// 如果最后的叶子节点 是 \x00 那么到此就结束了.
				// 字符串没有结束, 但是关键字结束了. 直接返回.
				return true , v.(string)
			}

			if ix == keylen {
				if v,ok :=keymap["\x00"] ; ok {
					// 如果字符串结束了, 关键字也结束了.刚好找到. 返回.
					return true , v.(string)
				}else {  // 字符串结束了,关键字没没结束, 没有找到关键字, 返回false
					return  false ,""
				}
			}else {
				continue
			}
		}else {
			keymap = set
			continue
		}
	}
	return false,""  //循环结束, 没有找到关键字return false
}


// build a map
//  every one  in  KEYWORD   will be a key in the map .
func Buildkeymap (set Keyword, str string) {
	keymap := make(Keyword)
	keymap = set
	keys :=make([]string, 0)
	str =strings.TrimSpace(str)
	str = strings.TrimSpace(str)
	for  _,key := range str {
		keys = append(keys,string(key))
	}
	keylength :=len(keys)
	for  ix, key :=range keys {
		val,ok := keymap[key]
		if ok {
			keymap = val.(Keyword)
			if ix == keylength -1  {
				if _,ok := keymap["\x00"] ; ok {
					return
				}else {
					keymap["\x00"] = str
					return
				}
			}
			continue
		}else {
			keymap[key]=make(Keyword)
			keymap = keymap[key].(Keyword)
			if ix == keylength -1 {
				keymap["\x00"]=str
			}

		}

	}

}

//  keyset  is  one  Keyword {map}   it is a sub set of  set .
//  then add  keyset  to the Set
// keyset is build  by Buildkeymap
func Add(set Keyword, keyset Keyword) {
	var keymap Keyword = set

	for {

		for k, v := range keyset {
			val, ok := keymap[k]
			if ok {
				keymap = val.(Keyword)
				keyset=v.(Keyword)
				if _,ok := keyset["\x00"]; ok { //keymap  is allready in keymap exit
					return
				}
				break
			}else{
				keymap[k] = v  // if  k not in Keymap then  add key:v into keymap
				return
			}
		}
		return
	}
	return

}



//  for test  print   the keymap in  recursive .
func Pk(km Keyword) {
	for k, v := range km {
		if k == "\x00" {
			fmt.Println("\\x00:", v.(string))
		} else {
			fmt.Println(k, v)
			Pk(v.(Keyword))
		}
	}
}




