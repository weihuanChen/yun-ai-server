package banner

import (
	"embed"
	"fmt"
	"log"
)

//go:embed banner.txt
var bannerFS embed.FS

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	italicStart = "\033[3m"
	italicEnd   = "\033[23m"
)

func PrintBanner() {
	// 读取嵌入的 banner 文件内容「BlurVision ASCII」
	content, err := bannerFS.ReadFile("banner.txt")
	if err != nil {
		log.Printf("无法读取 banner 文件: %v", err)
		return
	}

	banner := colorCyan + italicStart + string(content) + italicEnd + colorReset

	fmt.Println(banner)
}
