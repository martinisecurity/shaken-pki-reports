module github.com/martinisecurity/shaken-pki-reports

go 1.18

require (
	github.com/zmap/zcrypto v0.0.0-20231106212110-94c8f62efae4
	github.com/zmap/zlint/v3 v3.5.0
	golang.org/x/net v0.18.0
)

require (
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/weppos/publicsuffix-go v0.30.2-0.20230730094716-a20f9abcc222 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/zmap/zlint/v3 => github.com/martinisecurity/zlint/v3 v3.0.0-20231127163917-adf25183c08f
