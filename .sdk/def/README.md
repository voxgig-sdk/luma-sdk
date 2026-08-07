# API Definition

`luma-openapi.json` is Luma's own public OpenAPI specification, vendored here
**byte-for-byte unmodified** so that generation is reproducible and the source is verifiable.

| | |
|---|---|
| Source | https://public-api.luma.com/openapi.json |
| Fetched | 2026-08-07 |
| Size | 322,686 bytes |
| sha256 | `ce307166fa60d3bf3a9c7a18ce60c0076d5460a210c8c307aae0cf11bc6dab14` |
| Format | OpenAPI 3.1.0 — 66 paths, 66 operations, 9 webhook definitions |

Luma's developer docs publish this spec URL and state that you can use it "to generate client
libraries, import into tools like Postman, or build integrations with AI agents".

**Licensing.** The spec declares `info.license: {"name": "Proprietary"}`. It is reproduced here
for attribution and reproducibility, and it is **not** relicensed. The MIT licence at the root
of this repository covers the generated code only, not this file. Luma and LUMA are trademarks
of Luma Labs, Inc.; this repository is not affiliated with, endorsed by, or sponsored by Luma.

To re-verify:

```bash
curl -s https://public-api.luma.com/openapi.json | sha256sum
```
