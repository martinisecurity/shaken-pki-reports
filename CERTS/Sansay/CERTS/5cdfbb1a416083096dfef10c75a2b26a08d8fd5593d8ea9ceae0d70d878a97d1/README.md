# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Every1 Telecom 486K

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 124 day(s)\
Subject: CN=SHAKEN Every1 Telecom 486K, OU=President, O=Every1 Telecom, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cdn.cnxcdn.com/shaken/69.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC1DCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZaYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxOTAwMzExM1oXDTI0MDMxODAwMzExM1owbzELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRcwFQYDVQQKDA5FdmVyeTEgVGVsZWNvbTESMBAGA1UECwwJUHJlc2lkZW50MSMwIQYDVQQDDBpTSEFLRU4gRXZlcnkxIFRlbGVjb20gNDg2SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCjLZj8zhGPkCC0iLijO1zTqteJccJkiwAiQ%2BIErNbSmqiuQOYLT%2Bkrqj0ud%2FzMnOkEZLJepZVeKLXh9l7XJMzqjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ4NkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTdImcpMSSNy0%2B6YxbatnI1MMlDxDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAM6MsGxxFs55S67JSACoq5ks4Vwa71t419JwoW4MG9sXAiEAyVUttj6wnF%2BZpUTpe6g75FoRYFAWBq4IVMtwmGGRWXw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 486K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 16:51 UTC