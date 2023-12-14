package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize"
	"os"
	"strconv"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/12/13
  @desc: 解析excel转化成json文件
**/

// excelToJson 正常读取excel文件，大文件会等个20几秒
func excelToJson() {
	f, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Printf("open file false:%v\n", err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("close file false:%v\n", err)
			return
		}
	}()

	//cell, err := f.GetCellValue("IT保险利益", "A1")
	//if err != nil {
	//	fmt.Printf("getCellValue failed:%v\n", err)
	//	return
	//}
	//fmt.Println("get cell: ", cell)

	rows, err := f.GetRows("IT保险利益")
	if err != nil {
		fmt.Printf("getRows failed:%v\n", err)
		return
	}
	key := ""
	var jsonMap = make(map[string]float64)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, colCell := range row {
			//fmt.Print(colCell, "\t")
			switch j {
			case 0:
				if colCell == "0" {
					key += "M"
				} else {
					key += "F"
				}
			case 1:
				key += "_" + colCell
			case 2:
				key += "_Y" + colCell
			//	行3、4、5在文件中隐藏了
			case 6: // 起领年龄
				key += "_" + colCell
			case 8: // 可选责任
				if colCell == "0" {
					key += "_" + colCell
				}
			case 9:
				key += "_" + colCell
			case 10: //身故保险金
				str := strings.Replace(colCell, ",", "", -1)
				jsonMap[key], _ = strconv.ParseFloat(str, 64)
				// string to float

			default:
				continue
			}
		}
		// 重置key
		key = ""
		//fmt.Println()
		if i == 15 {
			break
		}
	}

	fmt.Printf("get jsonMap:\n %v\n", jsonMap)
	jsonFile, _ := json.Marshal(jsonMap)

	//err = ioutil.WriteFile("./jsonFile.json", jsonFile, 0644) //过时的写法
	err = os.WriteFile("./jsonFile.json", jsonFile, 0644)
	if err != nil {
		fmt.Printf("write file false: %v\n", err)
		return
	}
}

func excelToJson2() {
	f, err := excelize.OpenFile("E:\\CompanyFile\\新品知识\\信泰七金年金342909\\费率和现价.xlsx")
	if err != nil {
		fmt.Printf("open file false:%v\n", err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("close file false:%v\n", err)
			return
		}
	}()

	rows, err := f.GetRows("IT保险利益")
	if err != nil {
		fmt.Printf("getRows failed:%v\n", err)
		return
	}
	key := ""
	var jsonMap = make(map[string]float64)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, colCell := range row {
			//fmt.Print(colCell, "\t")
			switch j {
			case 0:
				if colCell == "0" {
					key += "M"
				} else {
					key += "F"
				}
			case 1:
				key += "_" + colCell
			case 2:
				key += "_Y" + colCell
			//	行3、4、5在文件中隐藏了
			case 6: // 起领年龄
				key += "_" + colCell
			case 8: // 可选责任
				if colCell == "0" {
					key += "_" + colCell
				}
			case 9:
				key += "_" + colCell
			case 10: //身故保险金
				str := strings.Replace(colCell, ",", "", -1)
				jsonMap[key], _ = strconv.ParseFloat(str, 64)
				// string to float

			default:
				continue
			}
		}
		// 重置key
		key = ""
		//fmt.Println()
		if i == 15 {
			break
		}
	}

	fmt.Printf("get jsonMap:\n %v\n", jsonMap)
	jsonFile, _ := json.Marshal(jsonMap)

	//err = ioutil.WriteFile("./jsonFile.json", jsonFile, 0644) //过时的写法
	err = os.WriteFile("./jsonFile.json", jsonFile, 0644)
	if err != nil {
		fmt.Printf("write file false: %v\n", err)
		return
	}
}
func main() {
	//excelToJson()

}
