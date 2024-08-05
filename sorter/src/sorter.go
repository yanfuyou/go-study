package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// *开头表示指针
var infile *string = flag.String("i", "infile", "file contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file contains values for sorting")
var algo *string = flag.String("a", "qsort", "sort algo")

func main() {
	// 解析命令行参数
	flag.Parse()
	if infile != nil {
		fmt.Printf("infile: %s,outfile :%s,algo: %s\n", *infile, *outfile, *algo)
	}
	vals, err := readVals(*infile)
	if err == nil {
		fmt.Println("vals:", vals)
		writeVals(vals, *outfile)
	} else {
		fmt.Println(err)
	}
}

/**
 * 小写函数名不会暴露给外部调用
 */
func readVals(infile string) (values []int, err error) {
	file, err := os.Open(infile)

	if err != nil {
		fmt.Println("field to open the inputFile", infile)
		return
	}
	// 关闭文件流
	defer file.Close()

	// 文件， 缓冲区
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		fmt.Println("line:", string(line), isPrefix)
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
				break
			}
		}
		if isPrefix {
			fmt.Println("line too long")
			return
		}
		// 转为字符串
		str := string(line)
		if str == "" {
			break
		}
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		// 向数组追加元素
		values = append(values, value)
	}
	return
}

func writeVals(vals []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("file to create outfile", outfile)
		return err
	}
	defer file.Close()

	for _, val := range vals {
		str := strconv.Itoa(val)
		file.WriteString(str + "\n")
	}
	return nil
}
