# Paste Go 📋

![Version](https://img.shields.io/visual-studio-marketplace/v/cointem.paste-go) ![Installs](https://img.shields.io/visual-studio-marketplace/i/cointem.paste-go) ![License](https://img.shields.io/github/license/cointem/paste-go)

**Paste Go** is a smart clipboard assistant for VS Code. It automatically detects the content in your clipboard (JSON, XML, SQL) and converts it into the corresponding data structure code for your current file's language.

**Paste Go** 是一个 VS Code 智能粘贴助手。它能自动检测剪贴板内容（JSON, XML, SQL），并将其转换为当前文件语言对应的结构体或类定义代码。

---

## ✨ Features / 功能特性

- ⚡ **Lightning Fast / 极速响应**:
  - Local parsing for standard formats. No internet required for basic usage.
  - 本地解析标准格式，基础功能无需联网，毫秒级响应。

- 🧠 **AI Fallback / AI 智能兜底**:
  - When local parsing fails or logic is complex, it automatically calls AI to generate the code.
  - 当本地解析失败或逻辑复杂时，自动调用 AI 生成代码。
  - Supports **DeepSeek**, **OpenAI**, **Gemini**, **Moonshot (Kimi)** and more.
  - 支持 DeepSeek, OpenAI, Gemini, Moonshot (Kimi) 等多种模型。

- 🔌 **Unified Architecture / 统一架构**:
  - **Inputs / 输入**: JSON, XML, SQL (CREATE TABLE).
  - **Outputs / 输出**: Go (Struct), TypeScript (Interface), Python (Pydantic), Java (Lombok), Rust (Serde).

---

## 🚀 Usage / 使用方法

1. **Copy** some JSON/SQL/XML text to your clipboard.
   这里复制一段 JSON/SQL/XML 文本。
2. Open a file (e.g., `user.go` or `types.ts`).
   打开一个代码文件（如 `user.go`）。
3. Press `Ctrl + Alt + V` (Mac: `Cmd + Alt + V`) or run command `Paste Go: Smart Paste (Struct)`.
   按下快捷键 `Ctrl + Alt + V` 或执行命令 `Paste Go: Smart Paste`。
4. 🎉 The code struct is automatically inserted!
   代码结构体即刻生成！

---

## ⚙️ Configuration / 配置 AI

To enable AI superpowers using your own API Key (e.g. DeepSeek):
如需启用 AI 增强功能（例如使用 DeepSeek），请在设置中配置：

### Method 1: GUI Settings (推荐)
1. Open Settings (`Ctrl + ,`) -> Search `Paste Go`.
   打开设置 -> 搜索 `Paste Go`。
2. **AI Provider**: Select `deepseek` (or `openai`, `gemini`).
   选择对应的服务商。
3. **API Key**: Enter your key (e.g., `sk-xxxx`).
   填入你的 API Key。
4. **Base URL**: (Crucial for DeepSeek/Moonshot) Enter the API endpoint.
   DeepSeek/Kimi 等模型必填，例如 `https://api.deepseek.com`。

### Method 2: `settings.json`

```json
{
    // DeepSeek Example
    "pasteGo.aiProvider": "deepseek",
    "pasteGo.aiApiKey": "sk-your-deepseek-key",
    "pasteGo.aiBaseUrl": "https://api.deepseek.com",
    "pasteGo.aiModel": "deepseek-chat",

    // Google Gemini Example
    // "pasteGo.aiProvider": "gemini",
    // "pasteGo.aiApiKey": "your-gemini-key", // No BaseURL needed usually
    // "pasteGo.aiModel": "gemini-1.5-flash"
}
```

---

## 🛠️ Requirements / 依赖

- **None!** The extension comes with a bundled lightweight Go binary (~6MB). You don't need to install Go or Node.js.
- **无依赖！** 插件自带精简版 Go 二进制核心 (~6MB)，无需安装 Go 或 Node.js 环境即可使用。

---

## 🤝 Contributing / 贡献

We welcome PRs! This project is built with **Go** (Core Logic) and **TypeScript** (VS Code Extension).
欢迎提交 PR！本项目由 **Go** (核心逻辑) 和 **TypeScript** (插件前端) 构建。

- **Core**: `core/` (Golang) - Parsers and Generators.
- **Extension**: `extension/` (Typescript) - UI and Process management.

---

**Enjoy Coding!** 🚀
