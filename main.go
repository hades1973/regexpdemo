package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "123 +123 -123 -321.5 789.001 +12345.789\n"
	// 匹配1个整数
	// r := regexp.MustCompile(`[+-]?[1-9]{1}[0-9]*`)
	// 匹配至少一个整数                             ----结尾是空格或是行尾
	// r := regexp.MustCompile(`([+-]?[1-9]{1}[0-9]*)([ \n])+`)
	r := regexp.MustCompile(`([+-]?[1-9]{1}[0-9]*)([ \n])`)
	if finds := r.FindAllString(s, -1); finds != nil {
		fmt.Printf("find %v in '%v'\n", finds, s)
		fmt.Println("--------------------\n")
	}
	// FindAllStringSubmatch与FindAllString的不同
	if finds := r.FindAllStringSubmatch(s, -1); finds != nil {
		fmt.Printf("find %v in '%v'\n", finds, s)
		fmt.Println("--------------------\n")
	}

	// 给每个匹配到的整数加上引号
	s2 := r.ReplaceAllStringFunc(s, func(matched string) string {
		return fmt.Sprintf("'%v'", matched)
	})
	fmt.Printf("Replaced result: %v\n", s2)

	// 给每个匹配到的整数加上引号
	s3 := r.ReplaceAllString(s, "'$1' ")
	fmt.Printf("Replaced result: %v\n", s3)

	// 匹配一个分数（不包括整数）
	r2 := regexp.MustCompile(`([+-]?[1-9]{1}[0-9]*)\.([0-9]*)[ \n]`)
	// 给每个匹配到的分数加上引号
	s4 := r2.ReplaceAllString(s, "'$1.$2' ")
	fmt.Printf("Replaced frac result: %v\n", s4)

	// 匹配一个分数或整数
	r3 := regexp.MustCompile(`([+-]?[1-9]{1}[0-9]*)\.?([0-9]*)[ \n]`)
	// 给每个匹配到的分数加上引号
	matchs := r3.FindAllStringSubmatchIndex(s, -1)
	fmt.Println(matchs)
	for _, match := range matchs {
		//		fmt.Println(s[match[0]:match[1]])

		if strings.Contains(s[match[0]:match[1]], ".") { // float number
			bs := r3.ExpandString(nil, "'$1.$2'", s, match)
			fmt.Printf("%s\n", string(bs))
		} else { // interger number
			bs := r3.ExpandString(nil, "'$1'", s, match)
			fmt.Printf("%s\n", string(bs))
		}

	}
	s3 = r3.ReplaceAllString(s, "'$1.$2' ")
	fmt.Printf("Replaced frac result: %v\n", s3)
}
