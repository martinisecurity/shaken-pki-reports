# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 12 Apr 23 21:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -42 day(s)\
Subject: CN=SHAKEN BareTelecom 864J, OU=BareTelecomLLC, O=BareTelecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/BareTelecom_864J

[View certificate details](https://understandingwebpki.com/?cert=MIIC1zCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkX8EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEzMDAyNTUwOFoXDTIzMDMwMTAyNTUwOFowczELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFDASBgNVBAoMC0JhcmVUZWxlY29tMRcwFQYDVQQLDA5CYXJlVGVsZWNvbUxMQzEgMB4GA1UEAwwXU0hBS0VOIEJhcmVUZWxlY29tIDg2NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQw0u5dL2msUaOT7g%2Fj6lYVb3ul3h1hjK065leet4kb2sx2sx5qDAAQFCjA9BYE0LbIeUyqVEHbyb5ycAclaAmho4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4NjRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUSHVBZeir0%2FV8iGj5J5nGJ1iCgm8wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB%2B6DLUR3SANoqQn5aYiDYrqXYoyYYiGx39v2Bq7uz83AiEA1I1dd%2FumXlzDUyKGwYkhr%2Fcx7IsewNSfmMz7e0UvZ94%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC