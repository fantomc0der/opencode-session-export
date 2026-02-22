# General Agent Instructions

## Language-Specific Instructions

> **MANDATORY**: This section defines conditional instruction files that MUST be loaded based on the files you are working with. Failure to load relevant instructions will result in inconsistent code quality.

**BEFORE writing, editing, or reviewing any code file:**
1. Check the table below for matching glob patterns
2. If your target file matches a pattern, **READ the linked instruction file FIRST**
3. Apply all guidelines from that instruction file throughout your work

| Glob Pattern | Language | Description | Instruction File |
|--------------|----------|-------------|------------------|
| `**/*.go`, `**/go.mod`, `**/go.sum` | Go | Instructions for writing Go code following idiomatic Go practices and community standards | [go.custom.md](.opencode/instructions/go.custom.md) |

## Available Skills

> **CONDITIONAL**: Skills provide specialized knowledge and step-by-step guidance for specific domains. Load them on demand when a task matches a skill's expertise.

**BEFORE delegating or starting specialized work:**
1. Check the table below for skills matching the task domain
2. If a skill matches, **LOAD it** via the `skill` tool or pass it as `load_skills` when delegating
3. Skills are not auto-loaded — they must be explicitly activated per task

| Skill Name | Description | Skill File |
|------------|-------------|------------|
| `ast-grep` | Guide for writing ast-grep rules to perform structural code search and analysis. Use when users need to search codebases using Abstract Syntax Tree (AST) patterns, find specific code structures, or perform complex code queries that go beyond simple text search. This skill should be used when users ask to search for code patterns, find specific language constructs, or locate code with particular structural characteristics. | [SKILL.md](.opencode/skills/ast-grep/SKILL.md) |
| `gh-cli` | GitHub CLI (gh) reference for repositories, pull requests, Actions, organizations, and search from the command line. | [SKILL.md](.opencode/skills/gh-cli/SKILL.md) |
