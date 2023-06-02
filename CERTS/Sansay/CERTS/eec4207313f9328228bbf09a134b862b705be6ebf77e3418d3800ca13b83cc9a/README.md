# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InPhonex 1821

Tested At: 02 Jun 23 01:08 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: 79 day(s)\
Subject: CN=SHAKEN InPhonex 1821, OU=NOC Team, O=InPhonex, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/InPhonex_1821

[View certificate details](https://understandingwebpki.com/?cert=MIICyDCCAm6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYh8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIyMDE3MzkxM1oXDTIzMDgxOTE3MzkxM1owZDELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExETAPBgNVBAoMCEluUGhvbmV4MREwDwYDVQQLDAhOT0MgVGVhbTEdMBsGA1UEAwwUU0hBS0VOIEluUGhvbmV4IDE4MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQy4FGoWngM6f0PzGtJtO4jMAWH9FA1mXNYOsre99xmVTb%2BnAn%2BQO2gYbqz5G4zJPk5Nk73CeWyCMHAaKkSj5Gdo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQxODIxMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUJ5NsBUXCUPYHrvfJNlWFL3enjwYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIENyyT9RidZIyU90VpgHnpN4l4wPDq25TqGCIIjXbGV3AiEAsxITSzGgwcZd0hij2xZNEEkgH%2FCiZiT0ef1lRPyY8Cc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1821' |


Generated: 02 Jun 23 01:12 UTC