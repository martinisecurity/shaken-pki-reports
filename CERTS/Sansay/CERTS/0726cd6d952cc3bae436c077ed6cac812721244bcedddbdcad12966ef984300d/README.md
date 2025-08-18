# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN FastCast Networks 318K

Tested At: 18 Aug 25 20:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN FastCast Networks 318K, OU=Fastcast, O=FastCast Networks, ST=Nevada, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://fastcast.46labs.com

[View certificate details](https://x509.io/?cert=MIIC2TCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkp6swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDgyMjE0MjEzOVoXDTI1MDgyMjE0MjEzOVowdTELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk5ldmFkYTEaMBgGA1UECgwRRmFzdENhc3QgTmV0d29ya3MxETAPBgNVBAsMCEZhc3RjYXN0MSYwJAYDVQQDDB1TSEFLRU4gRmFzdENhc3QgTmV0d29ya3MgMzE4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC0JMGhT9Z87iOvAFQAiHdcXrjRgLjC1uSe0jWuMGaeoW29edxhbx03ROnTNkqx7pGEEB0L0jt158yQDBN5NsmijgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxOEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTqQDe%2B4oToeA7p72aQAWka9pdYTDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgezb4KUGgVM%2Blt8eiblMZ6j5NYtw09q0buL260d09RXUCIQDzCvywbW07Rb4P7AKO%2F4VfrNL9xQy0elMRDvVjXjRs9Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318K', but common name is 'SHAKEN FastCast Networks 318K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC