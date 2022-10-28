# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate b9fc91972f86a60e6b515cb176b9616851ffb0ab
Tested At: 2022-10-28 18:15:36 +0000 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 252 day(s)\
Subject: CN=***SHAKEN***464D, OU=VOIP, O=Fibernetics, L=Cambridge, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA

Link: https://stir.fibernetics.ca/prod-cert2022.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDNzCCAt2gAwIBAgIUC4SNyy74%2Fk04h6i14xB%2FA6tSWrgwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMTAeFw0yMjA2MDMxODI0NDNaFw0yMzA3MDYxODI0NDNaMG4xCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjESMBAGA1UEBwwJQ2FtYnJpZGdlMRQwEgYDVQQKDAtGaWJlcm5ldGljczENMAsGA1UECwwEVk9JUDEZMBcGA1UEAwwQKioqU0hBS0VOKioqNDY0RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEUTvm%2BKH1Lt6cZ7xQHlRQ4cnwKuVhpiWB%2BHDyMXCpsTH6b6mxZFmq5DJ1Pm7GZABYCjVDfvtYaKbdwGXqBuNBGjggE1MIIBMTAWBggrBgEFBQcBGgQKMAigBhYENDY0RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFFPuVX0JqdTNaA6xmavSADV%2FCEq4MGAGCCsGAQUFBwEBBFQwUjBQBggrBgEFBQcwAoZEaHR0cDovL2NhY2VydHMtY2EuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENhbmFkYUNBMS5jcnQwFQYDVR0gBA4wDDAKBggrBgEEAYO3HzBABgNVHR8EOTA3MDWgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybDAdBgNVHQ4EFgQU8NH4cyPBI3VtI7kikPU%2Fzbz444EwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCID%2FJ9rPrOWkf2%2FAuS8RVdlXlI7QQRm63JAzjzTGK9uklAiEAloD%2BBaqRgqROWrTp2HJCc%2By8O4ThZ6fAKt4Psm5JVB8%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_extension_unknown | error | ATIS-1000080 | STI certificate shall not include extensions that are not specified |
| e_sti_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 464D' |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 28/10/2022 at 18:15:47