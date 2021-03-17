package common

import "fmt"

// 前景 背景 颜色
// ------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色

// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

// 控制台打印颜色
type Color string

const (
	Black  Color = "\033[0;30m%s\033[0m"
	Red    Color = "\033[0;31m%s\033[0m"
	Green  Color = "\033[0;32m%s\033[0m"
	Yellow Color = "\033[0;33m%s\033[0m"
	Blue   Color = "\033[0;34m%s\033[0m"
	White  Color = "\033[0;37m%s\033[0m"
)

type ColorMessage struct {
}

func (ColorMessage) PrintMsg(color Color, msg string) string {
	return fmt.Sprintf(string(color), msg)
}

var (
	C ColorMessage
)
