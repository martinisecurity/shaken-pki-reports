# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ReInvent

Tested At: 31 Oct 22 20:25 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 568 day(s)\
Subject: C=US, ST=AZ, L=Scottsdale, O=ReInvent Telecom, OU=ReInvent, CN=ReInvent\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11292.10180

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUMqB2kp5Yzq1lQdLtLKQ1tlwehy0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MjExNjM4MzVaFw0yNDA1MjExNjM4MzVaMHAxETAPBgNVBAMMCFJlSW52ZW50MREwDwYDVQQLDAhSZUludmVudDEZMBcGA1UECgwQUmVJbnZlbnQgVGVsZWNvbTETMBEGA1UEBwwKU2NvdHRzZGFsZTELMAkGA1UECAwCQVoxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2Bqkd5rMIyVyhJWXkLn%2FZRDT08aK3gZQKMyVp4zDgnl%2B2tlYJw6oAJCAUaDIdprYYEPaumDLXP1SAeStQw2K9g6OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUtfYGF6BVSbB4Ianeekjzpqipvx8wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ2NjNHMA0GCSqGSIb3DQEBCwUAA4IBAQAW7A5FhzUsWn2fX5n%2B3pvdzwPUry59qgthua%2BiXohWic0XFXuhlZNKf1Np4fP4LaV2a%2FQr3L0W%2F%2FMfsGxEmQHkjrdexUeOJ4HTZMMGHVWl6IFExPmYblto%2FhCnyM6jrSV4Eeas2o60qicErWGqsbxZqHmUB0nupUoz3Wscv9x68BCPiNM1fHiXpSW8SBYaJIlN6trNmVVJZEcz81fp3JWUpG94sHAw1g%2BmNOhK5pee7SU8bN8%2BIJNHWlxb4IJuaXvupFyaMgEFu%2BkFVu0dgcRjLJRE42jAgO4uHuXPB%2Fsky29Do4h8PymgKfgZ3DeD6UHSLDB1wmh7pvjKxCY9Vg3o)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 20:32:42