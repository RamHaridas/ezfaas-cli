# This script can be used to build the ezfaas-cli binary
# Pre-requistites are stated below:
# 1. Delete .goreleaser.yaml
# 2. git commit all the changes
# 3. git tag all your changes as well

go mod init ezfaas
goreleaser init
goreleaser release --snapshot --clean
goreleaser check
goreleaser build --single-target --clean