# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ReInvent

Tested At: 02 Nov 22 15:08 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 567 day(s)\
Subject: C=US, ST=AZ, L=Scottsdale, O=ReInvent Telecom, OU=ReInvent, CN=ReInvent\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11292.10180

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUMqB2kp5Yzq1lQdLtLKQ1tlwehy0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MjExNjM4MzVaFw0yNDA1MjExNjM4MzVaMHAxETAPBgNVBAMMCFJlSW52ZW50MREwDwYDVQQLDAhSZUludmVudDEZMBcGA1UECgwQUmVJbnZlbnQgVGVsZWNvbTETMBEGA1UEBwwKU2NvdHRzZGFsZTELMAkGA1UECAwCQVoxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2Bqkd5rMIyVyhJWXkLn%2FZRDT08aK3gZQKMyVp4zDgnl%2B2tlYJw6oAJCAUaDIdprYYEPaumDLXP1SAeStQw2K9g6OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUtfYGF6BVSbB4Ianeekjzpqipvx8wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ2NjNHMA0GCSqGSIb3DQEBCwUAA4IBAQAW7A5FhzUsWn2fX5n%2B3pvdzwPUry59qgthua%2BiXohWic0XFXuhlZNKf1Np4fP4LaV2a%2FQr3L0W%2F%2FMfsGxEmQHkjrdexUeOJ4HTZMMGHVWl6IFExPmYblto%2FhCnyM6jrSV4Eeas2o60qicErWGqsbxZqHmUB0nupUoz3Wscv9x68BCPiNM1fHiXpSW8SBYaJIlN6trNmVVJZEcz81fp3JWUpG94sHAw1g%2BmNOhK5pee7SU8bN8%2BIJNHWlxb4IJuaXvupFyaMgEFu%2BkFVu0dgcRjLJRE42jAgO4uHuXPB%2Fsky29Do4h8PymgKfgZ3DeD6UHSLDB1wmh7pvjKxCY9Vg3o)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Nov 22 15:10 UTC