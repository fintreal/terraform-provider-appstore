openapi-generator generate -i api/openapi.json -g go -o .
LC_ALL=C find . -path ./.git -prune -o -type f ! -path './api/generate.sh'  -exec sed -i '' 's|GIT_USER_ID/GIT_REPO_ID|fintreal/app-store-sdk-go|g' {} +
LC_ALL=C find . -path ./.git -prune -o -type f -exec sed -i '' '/decoder\.DisallowUnknownFields()/d' {} +
LC_ALL=C sed -i '' 's/APPLE_PAY CertificateType = "APPLE_PAY"/APPLE_PAY_ CertificateType = "APPLE_PAY"/' model_certificate_type.go
LC_ALL=C sed -i '' '/dec.DisallowUnknownFields()/d' utils.go
go mod tidy