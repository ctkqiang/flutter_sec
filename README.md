# 🐍 FlutterSec - Flutter 应用安全扫描工具

FlutterSec 是一个使用 Go 编写的命令行工具，用于扫描 Flutter 打包后的 APK 文件中的安全风险。它通过调用 apktool 解包 APK，并静态分析如 AndroidManifest.xml 的安全配置。

---

## ✨ 功能亮点

- ✅ 解包 Flutter APK（基于 apktool）
- ✅ 分析 AndroidManifest.xml：
  - 是否设置 `android:debuggable="true"`
  - 是否允许 `allowBackup="true"`
  - 是否允许明文通信 `usesCleartextTraffic="true"`
- 🛠️ 后续可扩展模块：
  - 敏感信息字符串扫描（如 apiKey、token）
  - 签名证书分析（是否是 debug key）
  - Flutter obfuscation 检查
  - 输出 Markdown/HTML 报告

---

## 🚀 快速开始

### ✅ 安装依赖

确保你本地已安装：

```bash
brew install apktool go
```

或使用其他方式安装 apktool：https://ibotpeaches.github.io/Apktool/install/

### ✅ 克隆项目

```bash
git clone https://github.com/ctkqiang/flutter_sec.git

或

git clone https://gitcode.com/ctkqiang_sr/flutter_sec.git

cd flutter_sec
```

### ✅ 构建工具

```bash
go build -o fluttersec
```

### ✅ 扫描 APK

```bash
./fluttersec scan ./path/to/your.apk
```

### ✅ 示例输出

```bash
📦 开始解包 APK： ./example.apk
📦 使用 apktool 解包中...
✅ 解包成功，路径: /tmp/fluttersec_decompiled

🔎 开始分析 AndroidManifest.xml ...
⚠️  Manifest 第 12 行包含 `debuggable=true`
⚠️  Manifest 第 22 行包含 `usesCleartextTraffic=true`
✅ Manifest 分析完成
```

---

## 📁 项目结构说明

```
fluttersec/
├── main.go                    // 程序入口
├── go.mod                     // Go 依赖管理
├── cmd/                       // 命令注册
│   ├── root.go
│   └── scan.go
├── internal/apk/             // APK 分析逻辑
│   ├── decompile.go
│   └── manifest.go
```

---

## 🧠 开发计划（Roadmap）

- [x] 支持基础的 APK 解包 + Manifest 分析
- [ ] 支持敏感字符串扫描
- [ ] 支持签名证书分析（keytool）
- [ ] 支持导出 Markdown 报告
- [ ] 支持 Flutter 源码路径扫描
- [ ] 支持 iOS `.app` 包分析

---

## ❤️ 贡献指南

欢迎提交 PR 或 issue，让工具更完整 ✨

---

## 📄 License

MIT License

---

## 🌟 开源项目赞助计划

### 用捐赠助力发展

感谢您使用本项目！您的支持是开源持续发展的核心动力。  
每一份捐赠都将直接用于：  
✅ 服务器与基础设施维护（魔法城堡的维修费哟~）  
✅ 新功能开发与版本迭代（魔法技能树要升级哒~）  
✅ 文档优化与社区建设（魔法图书馆要扩建呀~）

点滴支持皆能汇聚成海，让我们共同打造更强大的开源工具！  
（小仙子们在向你比心哟~）

---

### 🌐 全球捐赠通道

#### 国内用户

<div align="center" style="margin: 40px 0">

<div align="center">
<table>
<tr>
<td align="center" width="300">
<img src="https://github.com/ctkqiang/ctkqiang/blob/main/assets/IMG_9863.jpg?raw=true" width="200" />
<br />
<strong>🔵 支付宝</strong>（小企鹅在收金币哟~）
</td>
<td align="center" width="300">
<img src="https://github.com/ctkqiang/ctkqiang/blob/main/assets/IMG_9859.JPG?raw=true" width="200" />
<br />
<strong>🟢 微信支付</strong>（小绿龙在收金币哟~）
</td>
</tr>
</table>
</div>
</div>

#### 国际用户

<div align="center" style="margin: 40px 0">
  <a href="https://qr.alipay.com/fkx19369scgxdrkv8mxso92" target="_blank">
    <img src="https://img.shields.io/badge/Alipay-全球支付-00A1E9?style=flat-square&logo=alipay&logoColor=white&labelColor=008CD7">
  </a>
  
  <a href="https://ko-fi.com/F1F5VCZJU" target="_blank">
    <img src="https://img.shields.io/badge/Ko--fi-买杯咖啡-FF5E5B?style=flat-square&logo=ko-fi&logoColor=white">
  </a>
  
  <a href="https://www.paypal.com/paypalme/ctkqiang" target="_blank">
    <img src="https://img.shields.io/badge/PayPal-安全支付-00457C?style=flat-square&logo=paypal&logoColor=white">
  </a>
  
  <a href="https://donate.stripe.com/00gg2nefu6TK1LqeUY" target="_blank">
    <img src="https://img.shields.io/badge/Stripe-企业级支付-626CD9?style=flat-square&logo=stripe&logoColor=white">
  </a>
</div>

---

### 📌 开发者社交图谱

#### 技术交流

<div align="center" style="margin: 20px 0">
  <a href="https://github.com/ctkqiang" target="_blank">
    <img src="https://img.shields.io/badge/GitHub-开源仓库-181717?style=for-the-badge&logo=github">
  </a>
  
  <a href="https://stackoverflow.com/users/10758321/%e9%92%9f%e6%99%ba%e5%bc%ba" target="_blank">
    <img src="https://img.shields.io/badge/Stack_Overflow-技术问答-F58025?style=for-the-badge&logo=stackoverflow">
  </a>
  
  <a href="https://www.linkedin.com/in/ctkqiang/" target="_blank">
    <img src="https://img.shields.io/badge/LinkedIn-职业网络-0A66C2?style=for-the-badge&logo=linkedin">
  </a>
</div>

#### 社交互动

<div align="center" style="margin: 20px 0">
  <a href="https://www.instagram.com/ctkqiang" target="_blank">
    <img src="https://img.shields.io/badge/Instagram-生活瞬间-E4405F?style=for-the-badge&logo=instagram">
  </a>
  
  <a href="https://twitch.tv/ctkqiang" target="_blank">
    <img src="https://img.shields.io/badge/Twitch-技术直播-9146FF?style=for-the-badge&logo=twitch">
  </a>
  
  <a href="https://github.com/ctkqiang/ctkqiang/blob/main/assets/IMG_9245.JPG?raw=true" target="_blank">
    <img src="https://img.shields.io/badge/微信公众号-钟智强-07C160?style=for-the-badge&logo=wechat">
  </a>
</div>

---

愿每一个 Flutter 项目都能免受安全威胁 💖
