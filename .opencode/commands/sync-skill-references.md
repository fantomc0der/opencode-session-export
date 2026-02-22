---
description: 'Sync available skill references to AGENTS.md so agents can discover and load them conditionally'
---

<role>

You are a configuration synchronization assistant responsible for keeping the skill reference table in `AGENTS.md` synchronized with the skill files in `.opencode/skills/`.

</role>

<target-file>

The only file that will be updated by this prompt:

| File | Section Path |
|------|--------------|
| `AGENTS.md` | `## Available Skills` |

</target-file>

<instructions>

## Task Overview

Scan the `.opencode/skills/` folder for skill definition files and update the table in `AGENTS.md` under `## Available Skills` to reflect the current set of available skills.

## Step 1: Discover Skill Files

List all `SKILL.md` files in immediate subdirectories of `.opencode/skills/`. Each skill lives in its own folder:

```
.opencode/skills/
├── ast-grep/
│   └── SKILL.md
├── gh-cli/
│   └── SKILL.md
└── ...
```

Only `SKILL.md` files are candidates. Ignore all other files in skill folders (e.g., `references/`, examples, etc.).

## Step 2: Extract Metadata from Each Skill

Read the YAML frontmatter of each `SKILL.md` file and extract:

| Property | Table Column | Notes |
|----------|--------------|-------|
| `name` | Skill Name | Use exact value from frontmatter |
| `description` | Description | Use exact value from frontmatter |
| (derived) | Skill File | Link using relative path to the `SKILL.md` file |

### Validation

**Exclude `SKILL.md` files where:**
- `name` is missing or empty
- `description` is missing or empty

These files are misconfigured and should be reported to the user in Step 6.

## Step 3: Read Current AGENTS.md State

Read `AGENTS.md` and locate the section headed `## Available Skills`.

### If the section does not exist

Create it after the last existing section in `AGENTS.md`. Populate it with this exact boilerplate before proceeding to Step 4:

```markdown
## Available Skills

> **CONDITIONAL**: Skills provide specialized knowledge and step-by-step guidance for specific domains. Load them on demand when a task matches a skill's expertise.

**BEFORE delegating or starting specialized work:**
1. Check the table below for skills matching the task domain
2. If a skill matches, **LOAD it** via the `skill` tool or pass it as `load_skills` when delegating
3. Skills are not auto-loaded — they must be explicitly activated per task

| Skill Name | Description | Skill File |
|------------|-------------|------------|
```

### If the section already has content

Locate the existing table (identified by the `| Skill Name | Description | Skill File |` header row).

## Step 4: Update the Table

**Replace** the table rows (not the header) with entries for each qualifying skill file.

**Important behaviors:**
- This is a **full replacement** of table rows, not a partial update
- **Add** rows for new skill files
- **Update** rows where frontmatter values have changed
- **Remove** orphaned rows that reference skills that no longer exist
- Preserve the table header row exactly as-is
- Preserve the blockquote and numbered instructions above the table exactly as-is
- Sort rows alphabetically by Skill Name

### Row Format

Each row must follow this exact format:

```markdown
| `<skill-name>` | <description-value> | [SKILL.md](.opencode/skills/<folder-name>/SKILL.md) |
```

Where:
- `<skill-name>` is wrapped in backticks and matches the `name` frontmatter value exactly
- `<description-value>` matches the `description` frontmatter value exactly
- `<folder-name>` is the skill's directory name under `.opencode/skills/`

## Step 5: Example

Given this skill file at `.opencode/skills/ast-grep/SKILL.md`:

```yaml
---
name: ast-grep
description: Guide for writing ast-grep rules to perform structural code search and analysis.
---
```

And this skill file at `.opencode/skills/gh-cli/SKILL.md`:

```yaml
---
name: gh-cli
description: GitHub CLI (gh) reference for repositories, pull requests, Actions, organizations, and search from the command line.
---
```

The resulting table should be:

```markdown
| `ast-grep` | Guide for writing ast-grep rules to perform structural code search and analysis. | [SKILL.md](.opencode/skills/ast-grep/SKILL.md) |
| `gh-cli` | GitHub CLI (gh) reference for repositories, pull requests, Actions, organizations, and search from the command line. | [SKILL.md](.opencode/skills/gh-cli/SKILL.md) |
```

## Step 6: Report Changes

After updating, report:
- Number of skill files found
- Any excluded `SKILL.md` files, listed by folder name with the reason (e.g., missing `name`, missing `description`)
- Number of rows added/updated/removed
- Final state of the table

If no changes were needed, confirm that `AGENTS.md` is already in sync.

</instructions>
