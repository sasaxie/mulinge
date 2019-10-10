package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/prometheus/common/log"
)

type CellValue struct {
	Sheet string
	Cell  string
	Value string
}

type CellMerge struct {
	Sheet     string
	StartCell string
	EndCell   string
}

type CellStyle struct {
	Sheet     string
	StartCell string
	EndCell   string
	Style     *Style
}

type Style struct {
	Alignment *Alignment `json:"alignment,omitempty"`
	Font      *Font      `json:"font,omitempty"`
	Fill      *Fill      `json:"fill,omitempty"`
	Border    []*Border  `json:"border,omitempty"`
}

type Alignment struct {
	Horizontal string `json:"horizontal,omitempty"` // 水平对齐方式
	Vertical   string `json:"vertical,omitempty"`   // 垂直对齐方式
	WrapText   bool   `json:"wrap_text,omitempty"`  // 自动换行设置
}

type Font struct {
	Bold bool `json:"bold,omitempty"` // 粗体
}

type Fill struct {
	Type    string   `json:"type,omitempty"`    // 填充类型
	Color   []string `json:"color,omitempty"`   // 填充颜色
	Pattern int      `json:"pattern,omitempty"` // 填充模式
}

type Border struct {
	Type  string `json:"type,omitempty"`  // 类型
	Color string `json:"color,omitempty"` // 颜色
	Style int    `json:"style,omitempty"` // 风格
}

func (s *Style) Formatter() string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}

func main() {
	Example1()
	f := excelize.NewFile()

	cellValues := make([]*CellValue, 0)

	cellValues = append(cellValues, &CellValue{
		Sheet: "Sheet1",
		Cell:  "A1",
		Value: "序号",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "B1",
		Value: "巡检日期\r\n(XXXX/XX/XX)",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "C1",
		Value: "缺陷分类\r\n(处)",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "A3",
		Value: "1",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "B3",
		Value: "2019/07/25\r\n2019/09/26",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "C2",
		Value: "一般",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "D2",
		Value: "严重",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "E2",
		Value: "危急",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "C3",
		Value: "15",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "D3",
		Value: "0",
	}, &CellValue{
		Sheet: "Sheet1",
		Cell:  "E3",
		Value: "0",
	})

	for _, cellValue := range cellValues {
		err := f.SetCellValue(cellValue.Sheet, cellValue.Cell, cellValue.Value)
		if err != nil {
			log.Error(err)
		}
	}

	cellMerges := make([]*CellMerge, 0)
	cellMerges = append(cellMerges, &CellMerge{
		Sheet:     "Sheet1",
		StartCell: "A1",
		EndCell:   "A2",
	}, &CellMerge{
		Sheet:     "Sheet1",
		StartCell: "B1",
		EndCell:   "B2",
	}, &CellMerge{
		Sheet:     "Sheet1",
		StartCell: "C1",
		EndCell:   "E1",
	})

	for _, cellMerge := range cellMerges {
		err := f.MergeCell(cellMerge.Sheet, cellMerge.StartCell, cellMerge.EndCell)
		if err != nil {
			log.Error(err)
		}
	}

	cellStyles := make([]*CellStyle, 0)
	cellStyles = append(cellStyles, &CellStyle{
		Sheet:     "Sheet1",
		StartCell: "A1",
		EndCell:   "A2",
		Style: &Style{
			Alignment: &Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
			Font: &Font{Bold: true},
			Border: []*Border{
				&Border{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "top",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				},
			},
		},
	}, &CellStyle{
		Sheet:     "Sheet1",
		StartCell: "B1",
		EndCell:   "B2",
		Style: &Style{
			Alignment: &Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
			Fill: &Fill{
				Type:    "pattern",
				Color:   []string{"#CCFFFF"},
				Pattern: 1,
			},
			Font: &Font{Bold: true},
			Border: []*Border{
				&Border{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "top",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				},
			},
		},
	}, &CellStyle{
		Sheet:     "Sheet1",
		StartCell: "C1",
		EndCell:   "E2",
		Style: &Style{
			Alignment: &Alignment{
				Horizontal: "center",
				Vertical:   "center",
				WrapText:   true,
			},
			Fill: &Fill{
				Type:    "pattern",
				Color:   []string{"#CC99FF"},
				Pattern: 1,
			},
			Font: &Font{Bold: true},
			Border: []*Border{
				&Border{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "top",
					Color: "#000000",
					Style: 1,
				},
				&Border{
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				},
			},
		},
	}, &CellStyle{
		Sheet:     "Sheet1",
		StartCell: "A3",
		EndCell:   "E3",
		Style: &Style{
			Alignment: &Alignment{
				Horizontal: "center",
				WrapText:   true,
			},
		},
	})

	for _, cellStyle := range cellStyles {
		style, err := f.NewStyle(cellStyle.Style.Formatter())
		if err != nil {
			fmt.Println(err)
		}
		err = f.SetCellStyle(cellStyle.Sheet, cellStyle.StartCell, cellStyle.EndCell, style)
		if err != nil {
			fmt.Println(err)
		}
	}

	height := getFitRowHeight(cellValues)
	width := getFitColWidth("Sheet1", cellValues)

	// 设置行高
	for row, height := range height {
		err := f.SetRowHeight("Sheet1", row, height)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 设置列宽
	for col, width := range width {
		err := f.SetColWidth("Sheet1", col, col, width)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 根据指定路径保存文件
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func Example1() {
	f := excelize.NewFile()
	err := f.SetCellValue("Sheet1", "D7", "hi")
	if err != nil {
		fmt.Println(err)
	}

	style, err := f.NewStyle(`{"alignment":{"horizontal":"center","Vertical":"center"},"font":{"bold":true},"border":[{"type":"left","color":"FF0000","style":1}],"fill":{"type":"pattern","color":["#CCFFFF"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetCellStyle("Sheet1", "D7", "D7", style)
	if err != nil {
		fmt.Println(err)
	}

	// 根据指定路径保存文件
	err = f.SaveAs("./Example1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

// 根据单元格内容自动计算行高
// 算法说明：按行进行判断
func getFitRowHeight(cellValues []*CellValue) map[int]float64 {
	fitMap := make(map[int]float64)
	var rate int = 15

	for _, value := range cellValues {
		reg := regexp.MustCompile(`[\d]+`)
		numStrs := reg.FindAllString(value.Cell, -1)
		// 理论上第1个为数字
		row := 0
		if len(numStrs) > 0 {
			var err error
			row, err = strconv.Atoi(numStrs[0])
			if err != nil {
				break
			}
		}
		// 计算换行次数，每次乘以倍率rate
		count := strings.Count(value.Value, "\r\n") + 1
		height := float64(count * rate)
		// 最大限制为409
		if height > 409 {
			height = 409
		}
		if vv, ok := fitMap[row]; ok {
			if vv < height {
				fitMap[row] = height
			}
		} else {
			fitMap[row] = height
		}
	}

	return fitMap
}

func getFitColWidth(sheet string, cellValues []*CellValue) map[string]float64 {
	var rate float64 = 1.2

	maxFix := make(map[string]float64)
	for _, value := range cellValues {
		// 先暂时粗略的跳过跨列的
		if strings.EqualFold(value.Cell, "C1") || strings.EqualFold(value.Cell, "D1") || strings.EqualFold(value.Cell, "E1") {
			continue
		}
		reg := regexp.MustCompile(`[[:upper:]]+`)
		lettersStrs := reg.FindAllString(value.Cell, -1)
		// 理论上第1个为字母
		col := ""
		if len(lettersStrs) > 0 {
			col = lettersStrs[0]
		}
		// 计算字符串长度，每次乘以倍率rate
		split := strings.Split(value.Value, "\r\n")
		maxLength := 0
		for _, s := range split {
			length := searchCount(s)
			if maxLength < length {
				maxLength = length
			}
		}

		width := float64(maxLength) * rate
		// 最大限制为255
		if width > 255 {
			width = 255
		}

		if vv, ok := maxFix[col]; ok {
			if vv < width {
				maxFix[col] = width
			}
		} else {
			maxFix[col] = width
		}
	}

	return maxFix
}

// Excel所有出现的值进行匹配给出对应宽度值
func searchCount(src string) int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	letters = letters + strings.ToUpper(letters)
	nums := "0123456789"
	chars := "()/#"

	numCount := 0
	letterCount := 0
	othersCount := 0
	charsCount := 0

	for _, i := range src {
		switch {
		case strings.ContainsRune(letters, i) == true:
			letterCount += 1
		case strings.ContainsRune(nums, i) == true:
			numCount += 1
		case strings.ContainsRune(chars, i) == true:
			charsCount += 1
		default:
			othersCount += 1
		}
	}

	return numCount*1 + letterCount*1 + charsCount*1 + othersCount*2
}
