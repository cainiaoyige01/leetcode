package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/12 20:55
 * @Description:
 */
func main() {
	//读取文件
	f, err := excelize.OpenFile("excelFile.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭链接
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//获取工作表中指定的单元格的值
	cell, err := f.GetCellValue("Sheet1", "C5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	//rows, err := f.GetRows("Sheet1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	//cols, err := f.GetCols("Sheet1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, col := range cols {
	//	//for _, rowCell := range col {
	//	//	fmt.Print(rowCell, "\t")
	//	//}
	//	//fmt.Println()
	//	fmt.Println(col)
	//}
}
