package services

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
	"strings"
)

type File struct {
	ad string
	success int
	fail int
}
func ContrastFile()int {
	success :=0
	fail := 0

	//adData := &File{}
	var ad map[string]int
	ad = make(map[string]int)

	f, err := excelize.OpenFile("qq.xlsx")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err,1)
	//	return
	//}
	//fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err,2)
		return 0
	}
	//for {
	//	l, err := h.ReadString('\n') //读到换行
	//	fmt.Print(l)
	//	if err == io.EOF { //如果读到末尾就会进入
	//		return
	//	}
	//}
	//return
	for _, row := range rows {
		success = 0
		for key, colCell := range row {
			if key == 1 {
				//fmt.Print(colCell, "\t")
				//读取txt文件
				i,v := os.Open("lo1.txt")
				if v != nil {                                                                      //如果报错了进入if
					fmt.Println("打开文件失败", v)
					return 0
				}
				defer i.Close() //最后关闭文件
				h := bufio.NewReader(i)
				for {
					l, err := h.ReadString('\n') //读到换行
					//fmt.Print(l,colCell,strings.Contains(l,colCell))
					//fmt.Println(colCell+"相同")
					//return
					if strings.Contains(l,colCell) {
						//fmt.Println(colCell)
						success++
					}else{
						//fmt.Println(colCell+"失败")
						fail++
					}
					if err == io.EOF { //如果读到末尾就会进入
						//return
						break
					}
				}
				ad[colCell] = success
			}
		}
		//fmt.Println()
	}
	count := 0
	fmt.Print(success,fail)
	fmt.Println()
	for key,value := range ad{
		fmt.Println(key, "匹配成功", ad [ key ])
		count += value
	}
	fmt.Println(ad)
	return count
}
