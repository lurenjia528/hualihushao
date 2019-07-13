// 参考 https://yami.io/output-emojis/
package main

import (  
    "fmt"
    "html"
    "strconv"
)

func main() {  
    // 相关范围可以参考這里：http://apps.timwhitlock.info/emoji/tables/unicode
    emoji := [][]int{
        // 表情表示的范围。
        {128513, 128591},
        // 装饰符号的范围。
        {9986, 10160},
        // 交通工具,地图的范围。
        {128640, 128704},
    }

    for _, value := range emoji {
        for x := value[0]; x < value[1]; x++ {
            // 将范围转成 Unicode 字串后反脫逸字元，这样就会出现表情了。
            str := html.UnescapeString("&#" + strconv.Itoa(x) + ";")
            // 输出表情。
            fmt.Println(str)
        }
    }
}
