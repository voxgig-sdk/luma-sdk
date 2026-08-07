#!/usr/bin/env bash
# Rebuild the root README from .sdk/readme/{head,foot}.md, keeping the generator's own
# Entities table and How-it-works section, and correct the generated CHANGELOG's
# hardcoded language list.
#
# `npm run generate` in .sdk/ rewrites README.md and CHANGELOG.md from sdkgen templates,
# so run this afterwards. Editing README.md by hand does not survive a regenerate; edit
# .sdk/readme/head.md or foot.md instead.
set -euo pipefail
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
python3 - "$root" <<'PY'
import sys, pathlib, re
root = pathlib.Path(sys.argv[1])

# --- README -------------------------------------------------------------
gen = (root / 'README.md').read_text()

def slice_between(start_re, end_re, label):
    m1 = re.search(start_re, gen, re.M)
    if not m1:
        raise SystemExit(f'apply.sh: could not find start of {label} in generated README')
    m2 = re.search(end_re, gen[m1.start():], re.M)
    if not m2:
        raise SystemExit(f'apply.sh: could not find end of {label} in generated README')
    return gen[m1.start(): m1.start() + m2.start()].rstrip()

# Anchor on the entity-table heading specifically, NOT the prose section that is
# also called "## Entities, not endpoints".
entities = slice_between(r'^## Entities\n\nThe API exposes', r'^## Quickstart in other languages', 'Entities table')
# The generated per-language quickstarts are executed by each target's own
# README-example gate (py/php/ts/lua), so keep them rather than hand-writing snippets.
otherlangs = slice_between(r'^## Quickstart in other languages', r'^## Advanced', 'per-language quickstarts')
advanced = slice_between(r'^## Advanced\n', r'^## Per-language documentation', 'How it works')
advanced = advanced.replace('## Advanced', '## How it works', 1)

head = (root / '.sdk/readme/head.md').read_text()
foot = (root / '.sdk/readme/foot.md').read_text()
out = head + otherlangs + '\n\n' + entities + '\n\n' + advanced + '\n' + foot
(root / 'README.md').write_text(out)

for bad in ('{{', '## Packages', '## Surfaces', 'Ruby'):
    if bad in out:
        raise SystemExit(f'apply.sh: unexpected {bad!r} left in README.md')
print('README.md rebuilt (%d bytes)' % len(out))

# --- CHANGELOG ----------------------------------------------------------
# sdkgen hardcodes the language list in its Changelog component (it says "Ruby"),
# which is wrong for any SDK whose target set is not the default. Correct it here.
cl_path = root / 'CHANGELOG.md'
cl = cl_path.read_text()
fixed = cl.replace(
    'Initial generated release of the Luma SDK (TypeScript, Python, PHP, Go,\n  Ruby, and Lua, plus CLI and MCP surfaces)',
    'Initial generated release of the Luma SDK (TypeScript, JavaScript, Go,\n  Python, PHP, and Lua, plus CLI and MCP surfaces)')
if 'Ruby' in fixed:
    raise SystemExit('apply.sh: CHANGELOG still mentions Ruby - the sdkgen boilerplate changed, update this patch')
cl_path.write_text(fixed)
print('CHANGELOG.md language list corrected')
PY
