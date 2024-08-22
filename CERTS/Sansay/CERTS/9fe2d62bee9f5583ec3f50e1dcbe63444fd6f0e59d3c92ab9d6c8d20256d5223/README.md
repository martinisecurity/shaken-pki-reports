# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 22 Aug 24 15:32 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -508 day(s)\
Subject: CN=SHAKEN BareTelecom 864J, OU=BareTelecomLLC, O=BareTelecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/BareTelecom_864J_0302

[View certificate details](https://x509.io/?cert=MIIC1jCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYxswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMwMjE3NDExNVoXDTIzMDQwMTE3NDExNVowczELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFDASBgNVBAoMC0JhcmVUZWxlY29tMRcwFQYDVQQLDA5CYXJlVGVsZWNvbUxMQzEgMB4GA1UEAwwXU0hBS0VOIEJhcmVUZWxlY29tIDg2NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQw0u5dL2msUaOT7g%2Fj6lYVb3ul3h1hjK065leet4kb2sx2sx5qDAAQFCjA9BYE0LbIeUyqVEHbyb5ycAclaAmho4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4NjRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUSHVBZeir0%2FV8iGj5J5nGJ1iCgm8wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIF3HdALWAVJzFexarMLqhqdrtKGJuyAXIiX7ydH5TVvjAiB8pXBXWhWJ4OVg3JFtHh9DDJuc5jOTPuh3zR0ds3qVvw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J', but common name is 'SHAKEN BareTelecom 864J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC