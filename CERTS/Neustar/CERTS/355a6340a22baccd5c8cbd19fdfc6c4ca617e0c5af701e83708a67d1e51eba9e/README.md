# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ***SHAKEN***464D

Tested At: 15 Nov 23 17:15 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: -131 day(s)\
Subject: CN=***SHAKEN***464D, OU=VOIP, O=Fibernetics, L=Cambridge, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://stir.fibernetics.ca/prod-cert2022.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDNzCCAt2gAwIBAgIUC4SNyy74%2Fk04h6i14xB%2FA6tSWrgwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMTAeFw0yMjA2MDMxODI0NDNaFw0yMzA3MDYxODI0NDNaMG4xCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjESMBAGA1UEBwwJQ2FtYnJpZGdlMRQwEgYDVQQKDAtGaWJlcm5ldGljczENMAsGA1UECwwEVk9JUDEZMBcGA1UEAwwQKioqU0hBS0VOKioqNDY0RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEUTvm%2BKH1Lt6cZ7xQHlRQ4cnwKuVhpiWB%2BHDyMXCpsTH6b6mxZFmq5DJ1Pm7GZABYCjVDfvtYaKbdwGXqBuNBGjggE1MIIBMTAWBggrBgEFBQcBGgQKMAigBhYENDY0RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFFPuVX0JqdTNaA6xmavSADV%2FCEq4MGAGCCsGAQUFBwEBBFQwUjBQBggrBgEFBQcwAoZEaHR0cDovL2NhY2VydHMtY2EuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENhbmFkYUNBMS5jcnQwFQYDVR0gBA4wDDAKBggrBgEEAYO3HzBABgNVHR8EOTA3MDWgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybDAdBgNVHQ4EFgQU8NH4cyPBI3VtI7kikPU%2Fzbz444EwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCID%2FJ9rPrOWkf2%2FAuS8RVdlXlI7QQRm63JAzjzTGK9uklAiEAloD%2BBaqRgqROWrTp2HJCc%2By8O4ThZ6fAKt4Psm5JVB8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 464D' |


Generated: 15 Nov 23 17:17 UTC