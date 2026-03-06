# app-store-sdk-go

Go SDK for the Apple App Store Connect API (v3.8.0), generated with OpenAPI Generator and extended with custom JWT auth.

- Repo: `git@github.com:fintreal/app-store-sdk-go.git`
- Default branch: main

## Tech Stack

- Go 1.23+ (CI uses 1.24.1)
- OpenAPI Generator 7.16.0 (code generation)
- `github.com/golang-jwt/jwt/v5` (JWT auth)
- `gopkg.in/validator.v2` (validation)
- GitHub Actions (CI/CD)

## Project Structure

```
*.go              # Generated API client, models, and utils (package: openapi)
auth.go           # Custom JWT token generation (GenerateToken)
configuration.go  # Custom (protected from regeneration via .openapi-generator-ignore)
api/              # OpenAPI spec (openapi.json, original.json) and generate.sh
test/             # Integration tests (run against real App Store Connect API)
docs/             # Generated API documentation
.github/workflows # test.yaml (PR + nightly), release.yaml (semver on merge to main)
```

## API Coverage

- Apps (list, get, update)
- Bundle IDs (CRUD + relationships)
- Bundle ID Capabilities (CRUD)
- Certificates (CRUD)
- Provisioning Profiles (CRUD)

## Key Commands

```sh
# Regenerate client from OpenAPI spec
cd api && bash generate.sh

# Run integration tests (requires KEY_DATA, KEY_ID, ISSUER_ID env vars)
go test ./test/... -count=1 -cover
```

## Conventions

- Package name is `openapi`; import as `openapi "github.com/fintreal/app-store-sdk-go"`
- Most code is auto-generated -- do not edit `model_*.go` or `api_*.go` files directly
- Hand-maintained files protected in `.openapi-generator-ignore`: `configuration.go`, `test/*`, `api/openapi.yaml`, `.gitignore`
- `auth.go` is hand-written (JWT generation using ES256 with ECDSA private key)
- `generate.sh` post-processes generated code: replaces module path, removes `DisallowUnknownFields`, fixes enum conflicts
- Releases are auto-tagged with semver on merge to main
- Tests are integration tests that hit the real API -- they need App Store Connect credentials as env vars
