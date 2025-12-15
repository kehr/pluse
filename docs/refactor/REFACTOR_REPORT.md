# Crush → Pluse 品牌重命名完成报告

**执行日期**: 2025-12-15
**执行状态**: ✅ 完成

---

## 概述

本次重构将 fork 自 `charmbracelet/crush` 的仓库成功重命名为 `kehr/pluse`，包括：

| 变更项 | 原值 | 新值 |
|--------|------|------|
| 产品名称 | crush | pluse |
| Go 模块路径 | github.com/charmbracelet/crush | github.com/kehr/pluse |
| 环境变量前缀 | CRUSH_ | PLUSE_ |
| 配置文件名 | .crush*, crush.json | .pluse*, pluse.json |
| 数据目录 | .crush/ | .pluse/ |
| 忽略文件 | .crushignore | .pluseignore |
| 仓库所有者 | charmbracelet | kehr |
| 维护者邮箱 | - | kehr.dev@gmail.com |

---

## 执行阶段详情

### 阶段 1: Go 模块和导入路径 ✅

**修改文件数**: ~120+ 文件

| 文件 | 修改内容 |
|------|---------|
| `go.mod` | `module github.com/charmbracelet/crush` → `module github.com/kehr/pluse` |
| 所有 `*.go` 文件 | 导入路径批量替换 |

**批量替换命令**:
```bash
find . -name "*.go" -exec sed -i '' 's|github.com/charmbracelet/crush|github.com/kehr/pluse|g' {} \;
```

---

### 阶段 2: 核心常量和环境变量 ✅

**修改文件**:

| 文件 | 修改内容 |
|------|---------|
| `internal/config/config.go` | `appName = "pluse"`, `defaultDataDirectory = ".pluse"` |
| `internal/cmd/root.go` | CLI Use 名称、示例、错误消息 URL |
| `internal/config/load.go` | 环境变量处理逻辑 |
| `main.go` | `PLUSE_PROFILE` 环境变量 |
| `internal/shell/coreutils.go` | `PLUSE_CORE_UTILS` 环境变量 |
| `internal/agent/agent.go` | `PLUSE_DISABLE_ANTHROPIC_CACHE` 环境变量 |
| `internal/agent/common_test.go` | 测试环境变量 |
| `internal/log/log.go` | panic 日志文件名前缀 |
| `internal/agent/agentic_fetch_tool.go` | 临时文件前缀 |

**环境变量映射**:

| 原变量 | 新变量 |
|--------|--------|
| `CRUSH_PROFILE` | `PLUSE_PROFILE` |
| `CRUSH_DISABLE_PROVIDER_AUTO_UPDATE` | `PLUSE_DISABLE_PROVIDER_AUTO_UPDATE` |
| `CRUSH_DISABLE_METRICS` | `PLUSE_DISABLE_METRICS` |
| `CRUSH_DISABLE_ANTHROPIC_CACHE` | `PLUSE_DISABLE_ANTHROPIC_CACHE` |
| `CRUSH_CORE_UTILS` | `PLUSE_CORE_UTILS` |
| `CRUSH_ANTHROPIC_API_KEY` | `PLUSE_ANTHROPIC_API_KEY` |
| `CRUSH_OPENAI_API_KEY` | `PLUSE_OPENAI_API_KEY` |
| `CRUSH_OPENROUTER_API_KEY` | `PLUSE_OPENROUTER_API_KEY` |
| `CRUSH_ZAI_API_KEY` | `PLUSE_ZAI_API_KEY` |

---

### 阶段 3: 文件重命名和忽略模式 ✅

**文件重命名**:

| 原文件名 | 新文件名 |
|---------|---------|
| `crush.json` | `pluse.json` |
| `CRUSH.md` | `PLUSE.md` |

**忽略模式更新**:

| 文件 | 修改内容 |
|------|---------|
| `.gitignore` | `**/.crush/**` → `**/.pluse/**`, `/crush` → `/pluse` |
| `internal/fsext/ls.go` | `.crushignore` → `.pluseignore` |
| `internal/agent/tools/grep.go` | `.crushignore` → `.pluseignore` |
| `internal/agent/tools/grep.md` | `.crushignore` → `.pluseignore` |

---

### 阶段 4: 构建和发布配置 ✅

#### Taskfile.yaml

| 行 | 修改内容 |
|----|---------|
| 38 | LDFLAGS 路径更新 |
| 42 | 二进制名称 `crush` → `pluse` |
| 69 | 环境变量 `CRUSH_PROFILE` → `PLUSE_PROFILE` |
| 76 | LDFLAGS 路径更新 |

#### .goreleaser.yml

| 配置项 | 原值 | 新值 |
|--------|------|------|
| `project_name` | crush | pluse |
| `homepage` | charmbracelet/crush URL | https://github.com/kehr/pluse |
| `maintainers` | Charm | kehr <kehr.dev@gmail.com> |
| 所有仓库 `owner` | charmbracelet | kehr |
| NPM 包名 | @charmland/crush | @kehr/pluse |
| 所有 commit_author | charmbracelet 相关 | kehr <kehr.dev@gmail.com> |
| Shell 补全文件 | crush.* | pluse.* |
| Archive 名称模板 | crush_* | pluse_* |
| LDFLAGS | charmbracelet/crush 路径 | kehr/pluse 路径 |

**禁用的 charmbracelet 特有配置**:
- `includes` (notarize.yaml)
- `release.footer` (footer.md)
- `furies` (保留但禁用)

---

### 阶段 5: 文档和 Schema ✅

| 文件 | 修改内容 |
|------|---------|
| `README.md` | 标题、badge 链接、安装命令、配置示例、日志路径、GitHub 链接 |
| `CLAUDE.md` | 项目描述、配置文件路径、环境变量 |
| `PLUSE.md` | 标题和内容 (从 CRUSH.md 重命名) |
| `schema.json` | Schema ID URL、描述、示例路径 |
| `pluse.json` | Schema 引用 URL (从 crush.json 重命名) |
| `cspell.json` | 添加 "pluse" 到词典 |

---

### 阶段 6: GitHub 工作流 ✅

| 文件 | 修改内容 |
|------|---------|
| `.github/workflows/cla.yml` | 仓库判断条件、CLA 文档 URL、allowlist |
| `.github/workflows/schema-update.yml` | commit_user_name/email/author |
| `.gitattributes` | `crush-schema.json` → `pluse-schema.json` |

**保留未修改的工作流** (使用 charmbracelet/meta 共享工作流):
- `build.yml` - 引用 `charmbracelet/meta/.github/workflows/build.yml@main`
- `nightly.yml` - 引用 `charmbracelet/meta/.github/workflows/nightly.yml@main`
- `release.yml` - 引用 `charmbracelet/meta/.github/workflows/goreleaser.yml@main`

> **注意**: 这些是可复用的 GitHub Actions 工作流，可被任何仓库使用。如需完全独立，可创建自己的 meta 仓库或内联这些工作流。

---

### 阶段 7: 其他修改 ✅

| 文件 | 修改内容 |
|------|---------|
| `internal/agent/tools/bash.tpl` | Git commit/PR 消息中的 "Crush" → "Pluse", 邮箱更新 |
| `internal/agent/tools/mcp/init.go` | MCP 工具名称 |
| `internal/tui/highlight/highlight.go` | 样式名称 |
| `internal/tui/components/core/core.go` | 样式名称 |
| `internal/update/update.go` | GitHub API URL |

---

## 未修改的文件

### VCR 测试数据 (testdata)

`internal/agent/testdata/**/*.yaml` 文件包含录制的 HTTP 响应，其中包含 "crush" 字符串。这些是历史记录的 API 响应，**不应修改**，因为：

1. 它们是实际 API 调用的精确记录
2. 修改会导致测试失败
3. 新测试录制将自动使用新名称

如需重新录制所有测试数据：
```bash
task test:record
```

### 文档历史

`docs/refactor/REFACTOR_PLAN.md` 保留原始计划文档，包含旧名称作为历史参考。

---

## 本地清理建议

以下本地文件/目录可能需要手动清理：

```bash
# 旧的编译二进制
rm -f /Users/wangkaixuan.33/Studio/pluse/crush

# 旧的数据目录 (如果存在)
rm -rf /Users/wangkaixuan.33/Studio/pluse/.crush/
```

---

## 验证结果

### 源代码搜索

```bash
# crush 关键字搜索 (排除 testdata 和 docs/refactor)
grep -r "crush" --include="*.go" --include="*.md" --include="*.json" \
  --include="*.yaml" --include="*.yml" --include="*.tpl" | \
  grep -v testdata | grep -v "docs/refactor"
# 结果: 无匹配

# charmbracelet/crush 搜索
grep -r "charmbracelet/crush" --include="*.go" --include="*.md" \
  --include="*.json" --include="*.yaml" --include="*.yml" --include="*.tpl" | \
  grep -v "docs/refactor"
# 结果: 无匹配
```

---

## 后续步骤

1. **编译验证**:
   ```bash
   go build .
   # 或
   task build
   ```

2. **运行测试**:
   ```bash
   task test
   ```

3. **运行 lint**:
   ```bash
   task lint
   ```

4. **验证 CLI**:
   ```bash
   ./pluse --help
   ./pluse --version
   ```

5. **配置 GitHub Secrets** (用于发布):
   - `PERSONAL_ACCESS_TOKEN`
   - `HOMEBREW_TAP_GITHUB_TOKEN`
   - `AUR_KEY`
   - `NPM_TOKEN`
   - 其他 secrets 根据需要配置

6. **创建相关仓库** (可选):
   - `kehr/homebrew-tap` - Homebrew tap
   - `kehr/scoop-bucket` - Scoop bucket
   - `kehr/nur` - Nix user repository
   - `kehr/winget-pkgs` - Windows Package Manager

---

## 统计摘要

| 指标 | 数量 |
|------|------|
| 修改的 Go 文件 | ~120 |
| 修改的配置文件 | 12 |
| 修改的文档文件 | 5 |
| 修改的工作流文件 | 3 |
| 重命名的文件 | 2 |
| 替换的环境变量 | 9 |
| 更新的 URL | 15+ |

---

**重构完成！**
