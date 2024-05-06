# Build binary executables 
go mod init ezfaas
goreleaser init
goreleaser release --snapshot --clean
goreleaser check
goreleaser build --single-target

# Release binary to your github repository
# $env:GITHUB_TOKEN="YOUR_GH_TOKEN"
# git tag -a v0.1.0 -m "First release"
# git push origin v0.1.0
# goreleaser release
