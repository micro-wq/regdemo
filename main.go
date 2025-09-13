package main

import (
	"fmt"
	"regexp"
	"strings"
)

func extractProjectLast(source string) string {
	//str := "更新一个主机([5092])所属项目 BI,无 => GW"
	if strings.Contains(source, "取消关联项目") {
		compileRegex := regexp.MustCompile("项目\\((.*?)\\)$")
		matchArr := compileRegex.FindStringSubmatch(source)
		return matchArr[1]
	} else {
		compileRegex := regexp.MustCompile(" => (.*?)$")
		matchArr := compileRegex.FindStringSubmatch(source)
		return matchArr[1]
	}
}

func main() {
	str := "主机([1664 1663 1634 1633 1632 1631 1630 1629 1628 1627 937 197])取消关联项目(BI,BI,BI,BI,BI,BI,BI,BI,BI,BI,BI,BI)"
	fmt.Println(extractProjectLast(str))
}
