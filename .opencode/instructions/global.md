# Global Instructions

## General Best Practices

- Follow consistent coding style and conventions across the project.
- Use clear, descriptive naming — names should convey intent without needing a comment.
- Keep functions small and focused on a single responsibility.
- Follow the DRY principle — don't repeat yourself.
- Write unit tests for all non-trivial functionality.
- Be proactive in identifying and resolving issues.

## Formatting: `.editorconfig` Is Law

The `.editorconfig` file is the single source of truth for all formatting decisions — indentation style, indent size, line endings, trailing whitespace, charset, and any other property it defines.

**Rules:**
- Always respect `.editorconfig` settings above all other formatting preferences, tool defaults, or language conventions.
- If `.editorconfig` conflicts with a linter or formatter config, `.editorconfig` wins.
- Never override or work around `.editorconfig` settings without explicit user approval.
- When creating new files, apply the matching `.editorconfig` rules from the start.

## Code Comments: Explain Why, Not What

Write code that speaks for itself. Comment only when necessary — and when you do, explain **why**, not **what**.

### When NOT to Comment

- **Obvious statements** — don't narrate what the code already says (`counter++ // increment counter`).
- **Redundant descriptions** — if the function name and signature make the intent clear, a comment adds noise.
- **Dead code** — delete it, don't comment it out. That's what version control is for.
- **Changelogs** — don't track modification history in comments. Use git.
- **Decorative dividers** — no ASCII art section separators.

### When to Comment

- **Business logic rationale** — why this specific rule, threshold, or calculation exists.
- **Non-obvious algorithms** — why this approach was chosen over simpler alternatives.
- **External constraints** — API rate limits, library quirks, protocol requirements.
- **Regex patterns** — always explain what a non-trivial regex matches.
- **Configuration values** — document the reasoning behind magic numbers and thresholds.
- **Public API documentation** — document parameters, return values, and edge cases for exported/public interfaces.

### Annotations

Use standard annotations consistently:

- `TODO:` — planned improvement, include context on what and why.
- `FIXME:` — known broken behavior that needs a fix.
- `HACK:` — intentional workaround; document what it works around and when it can be removed.
- `NOTE:` — important assumption or context that isn't obvious from the code.
- `WARNING:` — surprising behavior (e.g., mutation, side effects).

### Before Writing a Comment, Ask:

1. Is the code already self-explanatory? → Skip the comment.
2. Would a better name eliminate the need? → Rename instead.
3. Does this explain **why**, not **what**? → Good comment.
4. Will this stay accurate as the code evolves? → If not, reconsider.
