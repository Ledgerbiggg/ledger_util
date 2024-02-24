/*
@author: ledger
@since: 2024/2/23
*/

package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"strconv"
)

type ExcelHelper struct {
	FileName  string
	SheetName string
	Header    []string
	Data      []any
}

func (e *ExcelHelper) GenerateExcel() error {
	file := excelize.NewFile()
	for i := 0; i < len(e.Header); i++ {
		k := 'A' + i
		file.SetCellValue(e.SheetName, string(rune(k))+"1", e.Header[i])
	}
	for i := 0; i < len(e.Data); i++ {
		row := i + 2 // 行索引从2开始，因为第一行是表头
		a := e.Data[i]
		t := reflect.ValueOf(a)
		// 遍历结构体的字段
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			k := 'A' + j
			file.SetCellValue(e.SheetName, string(rune(k))+strconv.Itoa(row), field.Interface())
		}
	}
	err := file.SaveAs(e.FileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func NewExcelHelper(fileName string, sheetName string, header []string, data []any) *ExcelHelper {
	helper := ExcelHelper{FileName: fileName, SheetName: sheetName, Header: header, Data: data}
	return &helper
}
