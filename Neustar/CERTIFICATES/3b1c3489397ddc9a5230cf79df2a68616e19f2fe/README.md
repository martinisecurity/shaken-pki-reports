# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate 3b1c3489397ddc9a5230cf79df2a68616e19f2fe
Tested At: 2022-10-27 21:25:38 +0000 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 452 day(s)\
Subject: C=US, ST=MO, L=Girardeau, O=Big River Telephone Company, OU=BRCNG, CN=BigRiverTelephoneCompany\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11133.10171.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIEEzCCAvugAwIBAgIUANhskzwFqmygO%2B6nOUBtHQJcEwQwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAxMjExODQ5NThaFw0yNDAxMjIxODQ5NThaMIGHMSEwHwYDVQQDDBhCaWdSaXZlclRlbGVwaG9uZUNvbXBhbnkxDjAMBgNVBAsMBUJSQ05HMSQwIgYDVQQKDBtCaWcgUml2ZXIgVGVsZXBob25lIENvbXBhbnkxEjAQBgNVBAcMCUdpcmFyZGVhdTELMAkGA1UECAwCTU8xCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmlngWIOLevw54sn3sb%2BTj3p%2BI0IVi5qdq8ijX%2BX6gGFaDF0ZEKpWkwnIEjgBIVUPL6rTzYSWMYM3sMwcW2dW16OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUwv7UK2fclkaYBWZlYtgPN7FOWsowDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQwMjNiMA0GCSqGSIb3DQEBCwUAA4IBAQA5w4jR646%2FEON5U0UqN74310EQrK4xglaRJew08tNMhrPTVDo5FeZskI57FYs4hJf1xY3SzwJcGWxbGdkssdVDuVNrVxco76QFfPpdOi84waK8rMOsrkGCiYENXemwOfXonC9Aga0wqHqk%2B9zfb%2F2CZh1noguVeOOd%2BBUqRnCIbwPnTYs3s4NpTLwrWpm6%2BCfHS%2FvHO27tq0pqdn%2BAIU%2FU3Q1UuO%2FjbRLR4XcGmzHO6FE6QuM%2B6DpYe2zzNIXcAFoXBHpqnXIjsBL8QWAVRZ1Cw0QWrdc5Il6XPjL3gipTer4P0xFjnOHNVJ2IcOQKlywRhJi3XthqDK9lD7LbY4D8)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_sti_extension_unknown
- w_cp_1_3_subject_email
- e_cp1_3_ambiguous_identifier
- e_sti_subject_cn
- e_cp1_3_subject_sn
- w_cp1_3_subject_rdn_unknown
- e_sti_serial_number
- e_sti_signature_algorithm

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 21:27:34