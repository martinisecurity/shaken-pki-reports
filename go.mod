module github.com/martinisecurity/shaken-pki-reports

go 1.18

require (
	github.com/zmap/zcrypto v0.0.0-20220803033029-557f3e4940be
	github.com/zmap/zlint/v3 v3.3.2-0.20220828143300-44e12c12ca43
)

require (
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/weppos/publicsuffix-go v0.15.1-0.20220724114530-e087fba66a37 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/zmap/zlint/v3 => github.com/martinisecurity/zlint/v3 v3.0.0-20221101162341-50f89e002411
