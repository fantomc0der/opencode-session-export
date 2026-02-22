---
description: 'Sync language specific instructions that apply to files matching specific glob patterns to AGENTS.md for OpenCode to conditionally load'
---

<role>

You are a configuration synchronization assistant responsible for keeping the OpenCode instruction table in `AGENTS.md` synchronized with the instruction files in `.opencode/instructions/`.
</role>

<target-file>

The only file that will be updated by this prompt:

| File | Section Path |
|------|--------------|
| `AGENTS.md` | `## Language-Specific Instructions` |

</target-file>

<instructions>

## Task Overview

Scan the `.opencode/instructions/` folder for instruction files and update the table in `AGENTS.md` under `## Language-Specific Instructions` to reflect the current set of language-specific instructions.

## Step 1: Discover Instruction Files

List all files matching `*.custom.md` in the `.opencode/instructions/` folder.

Other files in this folder (e.g., `global.md`) serve different purposes and are NOT candidates for the table. `global.md` is where instructions that target all files go — it is loaded automatically by OpenCode and is not considered an exclusion.

## Step 2: Filter Out Misconfigured Files

Read the YAML frontmatter of each `*.custom.md` file and check the `applyTo` property.

**Exclude `*.custom.md` files where:**
- `applyTo` is missing or empty
- `applyTo` matches all files (i.e., `'**'`, `'*'`, or `'**/*'`) — global instructions belong in `global.md`, not in `*.custom.md` files

These files are misconfigured and should NOT appear in the Language-Specific Guidelines table. They will be reported to the user in Step 6.

**Include `*.custom.md` files where:**
- `applyTo` targets specific file extensions (e.g., `'**/*.cs'`, `'**/*.ts'`)

## Step 3: Extract Metadata from Qualifying Files

For each qualifying instruction file, extract from the YAML frontmatter:

| Property | Table Column | Notes |
|----------|--------------|-------|
| `applyTo` | Glob Patterns | See comma-separated handling below |
| (derived) | Language | Infer language name from file extension in `applyTo` |
| `description` | Description | Use exact value from frontmatter |
| (derived) | Instruction File | Link using filename with path to `.opencode/instructions/` |

### Handling Comma-Separated `applyTo` Values

The `applyTo` property may contain multiple glob patterns separated by commas (e.g., `'**/*.go,**/go.mod,**/go.sum'`).

When `applyTo` contains commas, **keep all patterns in a single Glob Patterns cell** for that instruction file. Do NOT split them into separate rows.

1. Split the value by `,` to get individual glob patterns
2. Trim whitespace from each pattern
3. Place all patterns in the Glob Patterns cell, each wrapped in backticks and separated by `, `
4. Derive the Language from the **first** pattern that has a recognizable file extension

**One instruction file = one table row.** Never duplicate rows for the same file.

**Example:** Given `applyTo: '**/*.go,**/go.mod,**/go.sum'`, produce one row with Glob Patterns: `` `**/*.go`, `**/go.mod`, `**/go.sum` `` and Language: Go

### Language Derivation

Map common file extensions to language names:

| Extension | Language |
|-----------|----------|
| `.cs` | C# |
| `.ts` | TypeScript |
| `.tsx` | TypeScript (React) |
| `.js` | JavaScript |
| `.jsx` | JavaScript (React) |
| `.py` | Python |
| `.go` | Go |
| `.rs` | Rust |
| `.java` | Java |
| `.kt` | Kotlin |
| `.rb` | Ruby |
| `.php` | PHP |
| `.swift` | Swift |
| `.tf` | Terraform |
| `.md` | Markdown |

For filenames without a traditional extension (e.g., `go.mod`, `go.sum`, `Makefile`), derive the language from the full filename or its prefix:

| Filename Pattern | Language |
|------------------|----------|
| `go.mod`, `go.sum` | Go |
| `Makefile` | Make |
| `Dockerfile` | Docker |
| `Jenkinsfile` | Groovy |

For extensions or filenames not listed, use the extension without the dot, capitalized (e.g., `.yaml` becomes "YAML"). If no extension exists and no known filename match, use the filename capitalized.

## Step 4: Read Current AGENTS.md State

Read `AGENTS.md` and locate the section headed `## Language-Specific Instructions`.

### If the section is empty

If the heading exists but has no content beneath it (no blockquote, no instructions list, no table), populate it with this exact boilerplate before proceeding to Step 5:

```markdown
> **MANDATORY**: This section defines conditional instruction files that MUST be loaded based on the files you are working with. Failure to load relevant instructions will result in inconsistent code quality.

**BEFORE writing, editing, or reviewing any code file:**
1. Check the table below for matching glob patterns
2. If your target file matches a pattern, **READ the linked instruction file FIRST**
3. Apply all guidelines from that instruction file throughout your work

| Glob Patterns | Language | Description | Instruction File |
|---------------|----------|-------------|------------------|
```

### If the section already has content

Locate the existing table (identified by the `| Glob Patterns | Language | Description | Instruction File |` header row). The table structure is:

```markdown
| Glob Patterns | Language | Description | Instruction File |
|---------------|----------|-------------|------------------|
| ... | ... | ... | ... |
```

## Step 5: Update the Table

**Replace** the table rows (not the header) with entries for each qualifying instruction file.

**Important behaviors:**
- This is a **full replacement** of table rows, not a partial update
- **Add** rows for new instruction files
- **Update** rows where frontmatter values have changed
- **Remove** orphaned rows that reference instruction files that no longer exist
- Preserve the table header row exactly as-is
- Sort rows alphabetically by the first Glob Patterns in each row

### Row Format

Each row must follow this exact format:

```markdown
| `<glob-pattern>` | <Language> | <description-value> | [<filename>](.opencode/instructions/<filename>) |
```

Where:
- `<glob-pattern>` is wrapped in backticks. For single patterns, use the exact `applyTo` value. For comma-separated patterns, list each pattern in backticks separated by `, ` (e.g., `` `**/*.go`, `**/go.mod`, `**/go.sum` ``)
- `<Language>` is derived from the file extension of the first recognizable pattern
- `<description-value>` matches frontmatter exactly
- `<filename>` is the instruction file name with extension

## Example

### Single Pattern

Given this instruction file at `.opencode/instructions/csharp.instructions.md`:

```yaml
---
description: 'Guidelines for building Dotnet applications with C# language'
applyTo: '**/*.cs'
---
```

The corresponding table row should be:

```markdown
| `**/*.cs` | C# | Guidelines for building Dotnet applications with C# language | [csharp.instructions.md](.opencode/instructions/csharp.instructions.md) |
```

### Comma-Separated Patterns

Given this instruction file at `.opencode/instructions/go.custom.md`:

```yaml
---
description: 'Instructions for writing Go code following idiomatic Go practices and community standards'
applyTo: '**/*.go,**/go.mod,**/go.sum'
---
```

The corresponding table row should be:

```markdown
| `**/*.go`, `**/go.mod`, `**/go.sum` | Go | Instructions for writing Go code following idiomatic Go practices and community standards | [go.custom.md](.opencode/instructions/go.custom.md) |
```

## Step 6: Report Changes

After updating, report:
- Number of instruction files found
- Any excluded `*.custom.md` files, listed by name with the reason (e.g., missing `applyTo`, or `applyTo` targets all files). Remind the user that global instructions belong in `global.md`, not in `*.custom.md` files.
- Number of rows added/updated/removed
- Final state of the table

If no changes were needed, confirm that `AGENTS.md` is already in sync.

</instructions>
