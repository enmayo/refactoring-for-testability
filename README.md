# Intro
Look throug the following article: https://medium.com/@enmayopiniogn/refactoring-go-for-testability-3275a72e8eee

# Mocking
## Generation
### Testify
> mockery -dir . -all -output internal/mocks/testify
### GoMock
> mockgen -destination=internal/mocks/gomock/IoutilPkg.go -source=cmd/phase-4/dependencies.go 
