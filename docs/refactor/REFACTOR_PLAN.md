# Crush → Pluse 重构计划

## 概述

将 fork 仓库从 `charmbracelet/crush` 重命名为 `kehr/pluse`。

**基本替换规则：**
- 产品名称: `crush` → `pluse`
- 仓库地址: `github.com/charmbracelet/crush` → `github.com/kehr/pluse`
- 环境变量前缀: `CRUSH_` → `PLUSE_`
- 配置文件名: `.crush*` → `.pluse*`
- 数据目录: `.crush/` → `.pluse/`

---

## 已确认配置

| 配置项 | 决定 |
|--------|------|
| **向后兼容** | 不需要 - 完全替换为 .pluse* 文件 |
| **NPM 包名** | `@kehr/pluse` |
| **Schema URL** | `https://raw.githubusercontent.com/kehr/pluse/main/schema.json` |
| **发布渠道** | 保留全部但禁用 |
| **CLA 工作流** | 保留并更新为 kehr/pluse |
| **Homepage** | `https://github.com/kehr/pluse` |
| **Commit Author** | kehr |
| **Commit Email** | kehr.dev@gmail.com |
| **Apple 公证** | 保留但禁用 |

---

## 重构阶段

---

### 阶段 1: Go 模块和导入路径

**目标**: 修改 Go 模块路径，使项目可以编译

**修改文件**:

| 文件 | 修改内容 |
|------|----------|
| `go.mod` | `module github.com/charmbracelet/crush` → `module github.com/kehr/pluse` |
| `~120 个 Go 文件` | 导入路径 `github.com/charmbracelet/crush/internal/...` → `github.com/kehr/pluse/internal/...` |

**关键文件列表**:
- `main.go`
- `internal/cmd/*.go`
- `internal/app/*.go`
- `internal/agent/*.go`
- `internal/agent/tools/*.go`
- `internal/config/*.go`
- `internal/tui/**/*.go`
- `internal/lsp/*.go`
- `internal/session/*.go`
- `internal/db/*.go`
- 其他所有 `internal/` 子包

**验证计划**:
```bash
# 1. 编译检查
go build .

# 2. 检查是否有遗漏的导入
grep -r "charmbracelet/crush" --include="*.go" .
# 期望结果: 无输出
```

---

### 阶段 2: 核心常量和配置

**目标**: 修改应用名称、目录、环境变量等核心配置

**修改文件**:

| 文件 | 行号 | 修改内容 |
|------|------|----------|
| `internal/config/config.go` | 25 | `appName = "crush"` → `appName = "pluse"` |
| `internal/config/config.go` | 26 | `defaultDataDirectory = ".crush"` → `defaultDataDirectory = ".pluse"` |
| `internal/config/config.go` | 30-47 | 上下文路径: `crush.md` → `pluse.md`, `CRUSH.md` → `PLUSE.md` 等 |
| `internal/cmd/root.go` | 51 | `Use: "crush"` → `Use: "pluse"` |
| `internal/cmd/root.go` | 102 | GitHub issues URL |
| `internal/cmd/root.go` | 222 | `CRUSH_DISABLE_METRICS` → `PLUSE_DISABLE_METRICS` |

**环境变量替换** (所有文件):

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

**其他修改**:

| 文件 | 修改内容 |
|------|----------|
| `internal/config/load.go:105-122` | 环境变量前缀处理 `CRUSH_` → `PLUSE_` |
| `internal/update/update.go:15` | GitHub API URL |
| `internal/agent/tools/mcp/init.go:276` | MCP 工具名称 |
| `internal/tui/highlight/highlight.go:31` | 样式名称 |
| `internal/tui/components/core/core.go:204` | 样式名称 |
| `internal/log/log.go:56` | panic 日志文件名 `crush-panic-*` → `pluse-panic-*` |
| `internal/agent/agentic_fetch_tool.go:98` | 临时目录前缀 |
| `internal/agent/common_test.go:105` | 测试目录前缀 |

**验证计划**:
```bash
# 1. 构建并运行帮助命令
go build -o pluse .
./pluse --help
# 期望: 显示 "pluse" 而不是 "crush"

# 2. 检查环境变量名称
grep -r "CRUSH_" --include="*.go" .
# 期望结果: 无输出（除了可能的注释）
```

---

### 阶段 3: 文件名和忽略模式

**目标**: 重命名配置文件，修改忽略文件模式

**文件重命名**:

| 原文件 | 新文件 |
|--------|--------|
| `crush.json` | `pluse.json` |
| `CRUSH.md` | `PLUSE.md` |

**修改文件**:

| 文件 | 修改内容 |
|------|----------|
| `.gitignore:45` | `**/.crush/**` → `**/.pluse/**` |
| `.gitignore:47` | `/crush` → `/pluse` |
| `internal/fsext/ls.go` | `.crushignore` → `.pluseignore` (多处) |
| `internal/agent/tools/grep.go` | `.crushignore` → `.pluseignore` |
| `internal/agent/tools/grep_test.go` | `.crushignore` → `.pluseignore` |
| `internal/fsext/fileutil_test.go` | `.crushignore` → `.pluseignore` |
| `internal/fsext/ignore_test.go` | `.crushignore` → `.pluseignore` |

**验证计划**:
```bash
# 1. 检查忽略文件引用
grep -r "crushignore" --include="*.go" .
# 期望结果: 无输出

# 2. 检查配置文件加载
./pluse --debug 2>&1 | grep -i "config"
# 观察是否正确查找 pluse.json
```

---

### 阶段 4: 构建和发布配置

**目标**: 更新构建脚本和发布配置

**修改文件**:

| 文件 | 修改内容 |
|------|----------|
| `Taskfile.yaml:38,76` | LDFLAGS 路径 |
| `Taskfile.yaml:42` | 二进制名称 `crush` → `pluse` |
| `Taskfile.yaml:69` | `CRUSH_PROFILE` → `PLUSE_PROFILE` |

**.goreleaser.yml 修改**:

| 行号 | 原值 | 新值 |
|------|------|------|
| 5 | `project_name: crush` | `project_name: pluse` |
| 9 | `charmbracelet/meta/main/notarize.yaml` | 注释掉或设置 disable |
| 14 | `https://charm.sh/crush` | `https://github.com/kehr/pluse` |
| 18-19 | Charm maintainers | `kehr <kehr.dev@gmail.com>` |
| 36-39 | `crush.bash/zsh/fish` | `pluse.bash/zsh/fish` |
| 73 | LDFLAGS 路径 | 更新为 kehr/pluse |
| 78-84 | `crush_` | `pluse_` |
| 102 | AUR URL | 添加 `disable: true` |
| 104-105 | Charm commit author | `kehr <kehr.dev@gmail.com>` |
| 144-149 | AUR bin 配置 | 添加 `disable: true` |
| 176 | Fury account | 添加 `disable: "true"` |
| 180-182 | Homebrew tap | `kehr/homebrew-tap` + `skip_upload: true` |
| 185-187 | Charm commit author | `kehr <kehr.dev@gmail.com>` |
| 196-198 | Scoop bucket | `kehr/scoop-bucket` + `skip_upload: true` |
| 201-203 | Charm commit author | `kehr <kehr.dev@gmail.com>` |
| 206 | NPM name | `@kehr/pluse` |
| 207-208 | GitHub URLs | `kehr/pluse` |
| 210 | NPM | 添加 `disable: true` |
| 256-258 | Nix/NUR | `kehr/nur` + `skip_upload: true` |
| 260-262 | Charm commit author | `kehr <kehr.dev@gmail.com>` |
| 266 | Man page | `pluse.1.gz` |
| 269-274 | Winget | `kehr` + `skip_upload: true` |
| 276-278 | Charm commit author | `kehr <kehr.dev@gmail.com>` |
| 279-283 | Winget repo | `kehr/winget-pkgs` + `skip_upload: true` |
| 341 | Release footer | 移除或注释 |

**验证计划**:
```bash
# 1. 构建测试
task build

# 2. 检查二进制名称
ls -la pluse
# 期望: 存在 pluse 二进制文件

# 3. 验证 goreleaser 配置语法
goreleaser check
# 期望: 无错误
```

---

### 阶段 5: 文档和 Schema 更新

**目标**: 更新所有文档和配置 schema

**修改文件**:

| 文件 | 修改内容 |
|------|----------|
| `README.md` | 全面更新（标题、链接、安装命令、配置示例） |
| `CLAUDE.md` | 更新项目描述和配置路径引用 |
| `PLUSE.md` (原 CRUSH.md) | 更新标题 |
| `pluse.json` (原 crush.json) | 更新 schema 引用 |
| `schema.json:3` | Schema ID URL |
| `schema.json:25` | "Generated with Crush" → "Generated with Pluse" |
| `schema.json:364,428` | `CRUSH.md` → `PLUSE.md` |
| `schema.json:392,394` | `.crush` → `.pluse` |
| `cspell.json` | 添加 "pluse" 到词典 |

**README.md 主要修改**:
- 标题: `# Crush` → `# Pluse`
- 所有 badge URL: `charmbracelet/crush` → `kehr/pluse`
- Logo 图片: 移除或更换
- 安装命令中的包名
- 配置示例中的文件名
- 日志路径: `./.crush/logs/crush.log` → `./.pluse/logs/pluse.log`
- 所有 GitHub 链接

**验证计划**:
```bash
# 1. 检查文档中的旧引用
grep -r "charmbracelet/crush" README.md CLAUDE.md PLUSE.md
# 期望结果: 无输出

# 2. 检查 schema 中的旧引用
grep -r "crush" schema.json | grep -v "pluse"
# 期望结果: 无输出（除了可能的上游引用说明）

# 3. 验证 schema JSON 语法
python3 -m json.tool schema.json > /dev/null
# 期望: 无错误
```

---

### 阶段 6: GitHub 工作流

**目标**: 更新 GitHub Actions 工作流

**修改文件**:

| 文件 | 行号 | 修改内容 |
|------|------|----------|
| `.github/workflows/cla.yml:17` | `charmbracelet/crush` → `kehr/pluse` |
| `.github/workflows/cla.yml:30` | CLA 文档 URL |
| `.github/workflows/cla.yml:32` | allowlist: `charmcli,charmcrush` → `kehr` |
| `.github/workflows/cla.yml:34` | PR 评论中的 CLA URL |

**验证计划**:
```bash
# 1. 检查工作流文件语法
for f in .github/workflows/*.yml; do
  echo "Checking $f..."
  python3 -c "import yaml; yaml.safe_load(open('$f'))"
done
# 期望: 无错误

# 2. 检查旧引用
grep -r "charmbracelet" .github/workflows/
# 期望结果: 无输出
```

---

### 阶段 7: 最终验证

**目标**: 全面验证重构结果

**验证清单**:

```bash
# 1. 全局搜索旧名称
echo "=== 检查 crush 引用 ==="
grep -ri "crush" --include="*.go" --include="*.json" --include="*.yaml" --include="*.yml" --include="*.md" . | grep -v "pluse" | grep -v ".git" | head -50

# 2. 检查 charmbracelet 引用
echo "=== 检查 charmbracelet 引用 ==="
grep -ri "charmbracelet" --include="*.go" --include="*.json" --include="*.yaml" --include="*.yml" --include="*.md" . | grep -v ".git" | head -50

# 3. 编译测试
echo "=== 编译测试 ==="
go build -o pluse .

# 4. 运行测试
echo "=== 运行测试 ==="
task test

# 5. 运行 lint
echo "=== 运行 lint ==="
task lint

# 6. 功能测试
echo "=== 功能测试 ==="
./pluse --version
./pluse --help
```

**手动测试项目**:
- [ ] `./pluse --help` 显示正确的命令名
- [ ] `./pluse --version` 显示正确的版本信息
- [ ] 配置文件 `~/.config/pluse/pluse.json` 可以被正确读取
- [ ] 项目配置文件 `.pluse.json` 或 `pluse.json` 可以被正确读取
- [ ] 环境变量 `PLUSE_*` 可以正常工作
- [ ] `.pluseignore` 文件可以被正确识别

---

## 执行顺序

1. **阶段 1** → 验证通过 → 继续
2. **阶段 2** → 验证通过 → 继续
3. **阶段 3** → 验证通过 → 继续
4. **阶段 4** → 验证通过 → 继续
5. **阶段 5** → 验证通过 → 继续
6. **阶段 6** → 验证通过 → 继续
7. **阶段 7** → 全部验证通过 → 完成

**注意**: 每个阶段完成后必须通过验证才能进入下一阶段。如果验证失败，需要先修复问题。
