# STIR/SHAKEN CA Ecosystem Compliance

## Certificate BigRiverTelephoneCompany

Tested At: 02 Dec 22 07:45 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 417 day(s)\
Subject: C=US, ST=MO, L=Girardeau, O=Big River Telephone Company, OU=BRCNG, CN=BigRiverTelephoneCompany\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11133.10171.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIEEzCCAvugAwIBAgIUANhskzwFqmygO%2B6nOUBtHQJcEwQwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAxMjExODQ5NThaFw0yNDAxMjIxODQ5NThaMIGHMSEwHwYDVQQDDBhCaWdSaXZlclRlbGVwaG9uZUNvbXBhbnkxDjAMBgNVBAsMBUJSQ05HMSQwIgYDVQQKDBtCaWcgUml2ZXIgVGVsZXBob25lIENvbXBhbnkxEjAQBgNVBAcMCUdpcmFyZGVhdTELMAkGA1UECAwCTU8xCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmlngWIOLevw54sn3sb%2BTj3p%2BI0IVi5qdq8ijX%2BX6gGFaDF0ZEKpWkwnIEjgBIVUPL6rTzYSWMYM3sMwcW2dW16OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUwv7UK2fclkaYBWZlYtgPN7FOWsowDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQwMjNiMA0GCSqGSIb3DQEBCwUAA4IBAQA5w4jR646%2FEON5U0UqN74310EQrK4xglaRJew08tNMhrPTVDo5FeZskI57FYs4hJf1xY3SzwJcGWxbGdkssdVDuVNrVxco76QFfPpdOi84waK8rMOsrkGCiYENXemwOfXonC9Aga0wqHqk%2B9zfb%2F2CZh1noguVeOOd%2BBUqRnCIbwPnTYs3s4NpTLwrWpm6%2BCfHS%2FvHO27tq0pqdn%2BAIU%2FU3Q1UuO%2FjbRLR4XcGmzHO6FE6QuM%2B6DpYe2zzNIXcAFoXBHpqnXIjsBL8QWAVRZ1Cw0QWrdc5Il6XPjL3gipTer4P0xFjnOHNVJ2IcOQKlywRhJi3XthqDK9lD7LbY4D8)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Dec 22 07:46 UTC