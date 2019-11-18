# Dfa
Dfa  a  faster  search  keyword  from  a long string   


Dfa 是一个快速检索关键字的pkg.

Dfa is a project which to fast to search keyword from long strings .

it only find the first one keyword and then return .

is is fast than regexp .

Keyword Set is a map , every char in the keyword string is a key of this map .

it only fall through one times then get the keyword . .


dfs is一个 快速检索的 项目. 

快速的从一个字符串中,  检索出  特定的关键字.  

关键字 需要先 build  到 一个map 结构中. 

buildkeymap --> map[strig]interface{} 

Add    合并两个  keyemap , 如果分别包含相同的 分支, 则会忽略掉.  
          Add(set, keymap) . 
          为了可以实现并行化, 增加的参数. 
          
buildkeymap   可以 从关键字  --->  keymap结构.  

search      search(set,string) 从string 中检索  Set 中的关键字.
                              如果存在 .return (ｔｒｕｅ，　ｓｔｒｉｎｇ）　　
                              返回对应的ｋｅｙｗｏｒｄ　．　

ｅｘｐ：　　　ｓｅｔ　：＝　ｂｕｉｌｄｋｅｙｍａｐ（ｓｅｔ，　＂ａｂｃ＂）　　
            ｓｅａｒｃｈ　（ｓｅｔ　，　＂abcdef") ---> (true, "abc")  
            search (set , "def" ) --> (false, "") 
            searcha ( set  ," xxxxxxabc" ) -- >  (true , "abc") . 
            

Buildkeymap  :  var  set  Keyword 
                Buildkeymap (set , string)  --> map[string]interface{} 
                
 exp :    buildkeymap(set,"abcdef") -->
                    map["a":map["b":map["c":map["d":map["e":map["f":map["\x00":"abceef"]]]]]]]]
 
 

Add  :    Add(set , keymap) --> 

add  keymap( map[string]interface{} )  to  Set (map[string]interface{} ) 

Search  :  Search(set Keyword ,str string )  ( bool, string)   :

from  the string  to search a keyword in set . 


