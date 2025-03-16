go mod init ezfaas
goreleaser init
goreleaser release --snapshot --clean
goreleaser check
goreleaser build --single-target --clean