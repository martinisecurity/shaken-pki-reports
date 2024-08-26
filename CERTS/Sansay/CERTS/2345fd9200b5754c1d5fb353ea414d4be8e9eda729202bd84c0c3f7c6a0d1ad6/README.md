# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 26 Aug 24 17:56 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 61 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://primodialer.46labs.com

[View certificate details](https://x509.io/?cert=MIIC4jCCAoigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkha8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAyNzA5MzMyMFoXDTI0MTAyNjA5MzMyMFowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIHByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEInc%2B60%2Bo8RenjFw9K3CfIdS2A%2BZoSiWO1RpHkajruzFO%2Fb%2FT5q3Dyz1v8JddrviJEvBzj5A6jsHT9yKndb96qjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDI0OUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTKyc3rdvvhHI6JRj4QJ4hxKf8GpTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMr7EEoEFZPtO2C%2BzwEwBTDi0uF%2FSRokoPJqEEXZH1UOAiBUKZYoXRu2yojHN7frsWwQsYD0cpu3Uq6KkxMHouTL1Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K', but common name is 'SHAKEN Primo Dialler LLC 249K' |


Generated: 26 Aug 24 18:03 UTC