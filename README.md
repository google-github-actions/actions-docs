# actions-docs

A simple CLI to generate input/output documentation for your GitHub Actions and optionally insert it into a Markdown file.

## Installation

## Usage

```
  -action-metadata string
        Path to action metadata file (action.yml) (default "action.yml")
  -dry-run
        Print to stdout instead of overwriting readme
  -readme string
        Path to README.md (default "README.md")
```

Replacing content in a readme Markdown file makes use of HTML comments as markers to determine insert position.

For example:
```markdown
# foo
hello world

## Docs
<!-- BEGINNING OF PRE-COMMIT-ACTION DOCS HOOK -->
<!-- END OF PRE-COMMIT-ACTION DOCS HOOK -->

# Bar
```
This will insert the docs between the two markers.
## Examples

A GitHub Action metadata file like
```yaml
name: foo
author: bar
description: baz
inputs:
  input1:
    description: foo
    default: bar
  input2:
    description: baz
    default: qux
    required: true
outputs:
  output1:
    description: foo
```

will generate docs like
```markdown
| INPUT ID | DESCRIPTION | REQUIRED | DEFAULT |
|----------|-------------|----------|---------|
| `input1` | foo         | false    | bar     |
| `input2` | baz         | true     | qux     |

| OUTPUT ID | DESCRIPTION |
|-----------|-------------|
| `output1` | foo         |
```
