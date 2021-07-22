package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	rows int
	columns int
	twoD [][]int
}
func (m Matrix) getNumberOfRows() int{
	return m.rows
}
func (m Matrix) getNumberOfColumns() int{
    return m.columns
}
func (m *Matrix) setelements(i,j,num int) [][]int {
	m.twoD[i][j]=num
	return m.twoD
}
func (m *Matrix) addmatrices(mat1 [][]int) [][]int{

	for i :=0;i<m.rows;i++{
		for j :=0;j<m.columns;j++{
			m.twoD[i][j] = mat1[i][j] + m.twoD[i][j]
		}
	}
	return m.twoD
}

func (m Matrix) printJSON() {
	data,_:=json.Marshal(m)
	fmt.Println(string(data))
}


func main() {
   s := Matrix{rows:3,columns:3,twoD: [][]int{{1,2,3},{4,5,6},{7,8,9}}}
   r := s.getNumberOfColumns()
   c := s.getNumberOfRows()

   fmt.Println(r)
   fmt.Println(c)
   fmt.Println(s.setelements( 2,2, 1))
   fmt.Println(s.addmatrices([][]int{{1,2,3},{4,5,6},{7,8,9}}))

	s.printJSON()
}
