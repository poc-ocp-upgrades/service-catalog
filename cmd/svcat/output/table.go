package output

import (
	"io"
	"github.com/olekukonko/tablewriter"
)

const DefaultPageWidth = 80

type ListTable struct {
	table		*tablewriter.Table
	columnWidths	[]int
	variableColumn	int
	pageWidth	int
	headers		[]string
	rows		[][]string
}

func (lt *ListTable) SetBorder(b bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lt.table.SetBorder(b)
}
func (lt *ListTable) SetVariableColumn(c int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lt.variableColumn = c
}
func (lt *ListTable) SetColMinWidth(c, w int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lt.table.SetColMinWidth(c, w)
}
func (lt *ListTable) SetPageWidth(w int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lt.pageWidth = w
}
func (lt *ListTable) SetHeader(keys []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tmp := (len(keys) - len(lt.columnWidths)); tmp > 0 {
		lt.columnWidths = append(lt.columnWidths, make([]int, tmp)...)
	}
	for i, header := range keys {
		if tmp := len(header); tmp > lt.columnWidths[i] {
			lt.columnWidths[i] = tmp
		}
	}
	lt.headers = keys
}
func (lt *ListTable) Append(row []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tmp := (len(row) - len(lt.columnWidths)); tmp > 0 {
		lt.columnWidths = append(lt.columnWidths, make([]int, tmp)...)
	}
	for i, cell := range row {
		if tmp := len(cell); tmp > lt.columnWidths[i] {
			lt.columnWidths[i] = tmp
		}
	}
	lt.rows = append(lt.rows, row)
}
func (lt *ListTable) Render() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lt.variableColumn > 0 && lt.variableColumn <= len(lt.columnWidths)+1 {
		total := 2
		for i, w := range lt.columnWidths {
			if i+1 != lt.variableColumn {
				total = total + 3 + w
			}
		}
		remaining := lt.pageWidth - total - 3
		colWidth := lt.columnWidths[lt.variableColumn-1]
		if remaining >= colWidth {
			lt.SetColMinWidth(lt.variableColumn-1, colWidth)
		} else {
			lt.SetColMinWidth(lt.variableColumn-1, remaining)
		}
	}
	lt.table.SetHeader(lt.headers)
	for _, row := range lt.rows {
		lt.table.Append(row)
	}
	lt.table.Render()
}
func NewListTable(w io.Writer) *ListTable {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := tablewriter.NewWriter(w)
	t.SetBorder(false)
	t.SetColumnSeparator(" ")
	return &ListTable{table: t, pageWidth: DefaultPageWidth}
}
func NewDetailsTable(w io.Writer) *tablewriter.Table {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := tablewriter.NewWriter(w)
	t.SetAlignment(tablewriter.ALIGN_LEFT)
	t.SetBorder(false)
	t.SetColumnSeparator(" ")
	t.SetAutoWrapText(false)
	return t
}
