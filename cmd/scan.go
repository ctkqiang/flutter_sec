package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ctkqiang/flutter_sec/internal/apk"
	"github.com/spf13/cobra"
)

// 定义 scan 命令：用于扫描 Flutter 编译出来的 APK 安全性
var scanCmd = &cobra.Command{
	Use:   "scan [apk文件路径]",
	Short: "扫描 Flutter APK 的常见安全问题",
	Long: `这个命令会对 APK 文件进行静态分析，分析包括：
- AndroidManifest.xml 的配置（debuggable, allowBackup 等）
- 明文 HTTP 配置（usesCleartextTraffic）
- 后续将添加签名检查、敏感信息搜索等模块`,
	Args: cobra.ExactArgs(1), // 必须输入 1 个参数：APK 文件路径
	Run: func(cmd *cobra.Command, args []string) {
		apkPath := args[0]

		// 1. 检查 APK 文件是否存在
		if _, err := os.Stat(apkPath); os.IsNotExist(err) {
			fmt.Printf("❌ 文件不存在：%s\n", apkPath)
			os.Exit(1)
		}

		fmt.Println("📦 开始解包 APK：", apkPath)
		outputDir := filepath.Join(os.TempDir(), "fluttersec_decompiled") // 临时输出目录

		// 2. 使用 apktool 解包 APK 文件
		if err := apk.Decompile(apkPath, outputDir); err != nil {
			fmt.Printf("❌ 解包失败：%v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ 解包完成，文件输出在：", outputDir)

		// 3. 分析 AndroidManifest.xml 文件中的安全配置
		manifestPath := filepath.Join(outputDir, "AndroidManifest.xml")
		if err := apk.AnalyzeManifest(manifestPath); err != nil {
			fmt.Printf("❌ Manifest 分析失败：%v\n", err)
		}

		// 🧩 TODO：后续模块可以在这里继续添加
		// - 分析签名证书是否为 debug key
		// - 搜索敏感信息字符串（token、apiKey 等）
		// - 检查网络安全策略
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// 👉 后续可以加 flag，例如输出报告路径
	// scanCmd.Flags().StringP("output", "o", "", "将扫描结果导出到文件")
}
