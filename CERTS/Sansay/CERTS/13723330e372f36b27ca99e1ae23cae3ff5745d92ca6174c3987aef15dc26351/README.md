# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN FastCast Networks 318K

Tested At: 15 Nov 23 16:36 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 220 day(s)\
Subject: CN=SHAKEN FastCast Networks 318K, OU=Fastcast, O=FastCast Networks, ST=Nevada, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://fastcast.46labs.com

[View certificate details](https://understandingwebpki.com/?cert=MIIC2TCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdjcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyMjE4NDI1NloXDTI0MDYyMTE4NDI1NlowdTELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk5ldmFkYTEaMBgGA1UECgwRRmFzdENhc3QgTmV0d29ya3MxETAPBgNVBAsMCEZhc3RjYXN0MSYwJAYDVQQDDB1TSEFLRU4gRmFzdENhc3QgTmV0d29ya3MgMzE4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC0JMGhT9Z87iOvAFQAiHdcXrjRgLjC1uSe0jWuMGaeoW29edxhbx03ROnTNkqx7pGEEB0L0jt158yQDBN5NsmijgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxOEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTqQDe%2B4oToeA7p72aQAWka9pdYTDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgZ0PTv%2BoU6k%2F40F%2FU%2FcQQ5u7RJs1tugeeFX7fKM9%2F6zsCIQCeoMBfM5%2F4aBU9amGLKb3tURB46UXcPrxOcJo1s9ZLzw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 318K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC