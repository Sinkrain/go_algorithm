package deduplicate


import (
	"fmt"
	"strings"
)

// CenterWidth The width of Output
const (
	CenterWidth int = 90
	LinkeStr string = "="
)


// CenterMsg Format the msg to the center of Output
func CenterMsg(msg string){

	if len(msg) < CenterWidth{
		w := (CenterWidth - len(msg)) / 2
		midStr := strings.Join(make([]string, w), LinkeStr)
		fmt.Printf("\n<< %s %s %s >>\n", midStr, msg, midStr)
		
	} else {
		fmt.Println(msg)
	}
}
