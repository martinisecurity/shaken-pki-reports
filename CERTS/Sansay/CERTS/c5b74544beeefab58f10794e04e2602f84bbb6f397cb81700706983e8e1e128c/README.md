# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InPhonex 1821

Tested At: 12 Feb 24 16:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 191 day(s)\
Subject: CN=SHAKEN InPhonex 1821, OU=NOC Team, O=InPhonex, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/InPhonex_1821

[View certificate details](https://understandingwebpki.com/?cert=MIICyDCCAm6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfc8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgyMjAwMDYxNVoXDTI0MDgyMTAwMDYxNVowZDELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExETAPBgNVBAoMCEluUGhvbmV4MREwDwYDVQQLDAhOT0MgVGVhbTEdMBsGA1UEAwwUU0hBS0VOIEluUGhvbmV4IDE4MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQy4FGoWngM6f0PzGtJtO4jMAWH9FA1mXNYOsre99xmVTb%2BnAn%2BQO2gYbqz5G4zJPk5Nk73CeWyCMHAaKkSj5Gdo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQxODIxMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUJ5NsBUXCUPYHrvfJNlWFL3enjwYwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDj83WRAMcMnrR%2FDlBjnqfIS%2B%2BXiXc8zhS2M2krnhqPnQIgBFp3kWJTt4lJJhElsvXyDG3QICozbweqbb4djPZqRpQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1821', but common name is 'SHAKEN InPhonex 1821' |


Generated: 12 Feb 24 17:02 UTC