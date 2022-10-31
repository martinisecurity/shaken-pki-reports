# STIR/SHAKEN CA Ecosystem Compliance

## Certificate BigRiverTelephoneCompany

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 449 day(s)\
Subject: C=US, ST=MO, L=Girardeau, O=Big River Telephone Company, OU=BRCNG, CN=BigRiverTelephoneCompany\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11137.10171.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIEEzCCAvugAwIBAgIUANhskzwFqmygO%2B6nOUBtHQJcEwQwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAxMjExODQ5NThaFw0yNDAxMjIxODQ5NThaMIGHMSEwHwYDVQQDDBhCaWdSaXZlclRlbGVwaG9uZUNvbXBhbnkxDjAMBgNVBAsMBUJSQ05HMSQwIgYDVQQKDBtCaWcgUml2ZXIgVGVsZXBob25lIENvbXBhbnkxEjAQBgNVBAcMCUdpcmFyZGVhdTELMAkGA1UECAwCTU8xCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmlngWIOLevw54sn3sb%2BTj3p%2BI0IVi5qdq8ijX%2BX6gGFaDF0ZEKpWkwnIEjgBIVUPL6rTzYSWMYM3sMwcW2dW16OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUwv7UK2fclkaYBWZlYtgPN7FOWsowDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQwMjNiMA0GCSqGSIb3DQEBCwUAA4IBAQA5w4jR646%2FEON5U0UqN74310EQrK4xglaRJew08tNMhrPTVDo5FeZskI57FYs4hJf1xY3SzwJcGWxbGdkssdVDuVNrVxco76QFfPpdOi84waK8rMOsrkGCiYENXemwOfXonC9Aga0wqHqk%2B9zfb%2F2CZh1noguVeOOd%2BBUqRnCIbwPnTYs3s4NpTLwrWpm6%2BCfHS%2FvHO27tq0pqdn%2BAIU%2FU3Q1UuO%2FjbRLR4XcGmzHO6FE6QuM%2B6DpYe2zzNIXcAFoXBHpqnXIjsBL8QWAVRZ1Cw0QWrdc5Il6XPjL3gipTer4P0xFjnOHNVJ2IcOQKlywRhJi3XthqDK9lD7LbY4D8)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_sti_crl_distribution_not_reachable](../../ISSUES/e_sti_crl_distribution_not_reachable/README.md) | error | ATIS&#x2011;1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22