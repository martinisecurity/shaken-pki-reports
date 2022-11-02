# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Five9

Tested At: 02 Nov 22 17:23 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 532 day(s)\
Subject: C=US, ST=CA, L=SanRamon, O=Five9 Inc, OU=Five9, CN=Five9\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11200.10185

[View certificate details](https://understandingwebpki.com/?cert=MIID7DCCAtSgAwIBAgIUYcNE88r%2BDL6q%2Fzei%2FaP5iTJtVKswDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA0MTYyMDI2MzdaFw0yNDA0MTYyMDI2MzdaMGExDjAMBgNVBAMMBUZpdmU5MQ4wDAYDVQQLDAVGaXZlOTESMBAGA1UECgwJRml2ZTkgSW5jMREwDwYDVQQHDAhTYW5SYW1vbjELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYe6fU5sFNQH4EyFVSE4QDVW2t%2BzzrRM4qx4ES%2BMXxtYaQ0CqNWpbUOKOSzqgoEH9fTG4vM2odQFMAJf5BQmSpaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUAOiUdNHoS9wFZisWmUrHelF3LbowDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ1OThKMA0GCSqGSIb3DQEBCwUAA4IBAQBE7lue81wzFqwsX%2B7zBcLJ1PEXX8pnf%2FmHnQ4Uc%2Bq6VqgXpDLHMsAedeO57XdpZJ%2FtsItF9ejYhTMBlwWFOn7XW2QtG34cp%2FEiuNY82Loyuv7nGlfEpYqxTLf1IAOvF%2FJdEb%2By0ZOxqI13tYZZT4bmxvqQEH8CErLGM8m28Hw5onNXaYixexKKjY17NxRqIlIQ9o5QB85ymEE7FtNxLV99B0O4kfPj%2F8ryYrIcjzBEQyELEvmEwhqbfTirTyldF19glBlf%2B7%2FmpABlQ1aJyRRnyRKlFxkfWo4gJz%2B7zOtxjd6sDld%2FAMCc7A5jkUUOaBmzxGcBCERFwj0TXF%2Fp3FQ7)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Nov 22 17:25 UTC