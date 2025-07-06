package apk

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AnalyzeManifest(manifestPath string) error {
	file, err := os.Open(manifestPath)
	if err != nil {
		return fmt.Errorf("无法打开 AndroidManifest.xml: %v", err)
	}
	defer file.Close()

	fmt.Println("🔎 分析 AndroidManifest.xml ...")

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "android:debuggable=\"true\"") {
			fmt.Printf("⚠️  Manifest 第 %d 行包含 `debuggable=true`（易被反调试绕过）\n", lineNum)
		}

		if strings.Contains(line, "android:allowBackup=\"true\"") {
			fmt.Printf("⚠️  Manifest 第 %d 行包含 `allowBackup=true`（用户数据可能被备份）\n", lineNum)
		}

		if strings.Contains(line, "android:usesCleartextTraffic=\"true\"") {
			fmt.Printf("⚠️  Manifest 第 %d 行允许明文流量（http，易受 MITM 攻击）\n", lineNum)
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取 Manifest 时发生错误: %v", err)
	}

	fmt.Println("✅ Manifest 分析完成")
	return nil
}
