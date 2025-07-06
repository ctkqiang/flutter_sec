package apk

import (
	"fmt"
	"os/exec"
)

// ✅ 首字母大写，export 出去才能被 cmd/scan.go 调用
func Decompile(apkPath string, outDir string) error {
	fmt.Println("📦 使用 apktool 解包中...")

	cmd := exec.Command("apktool", "d", "-f", apkPath, "-o", outDir)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output)) // 可选：打印 apktool 输出

	return err
}
