/*
@author: ledger
@since: 2024/2/23
*/
package utils

import (
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func TestExcelHelper_GenerateExcel(t *testing.T) {
	people := []any{
		Person{"Alice", 30, "USA"},
		Person{"Bob", 25, "Canada"},
		Person{"Charlie", 35, "UK"},
	}
	helper := NewExcelHelper("persons1.xlsx", "Sheet1", []string{"Name", "Age", "Country"}, people)
	err := helper.GenerateExcel()
	if err != nil {
		t.Error(err)
	}
}
