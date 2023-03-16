# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Asia Pacific Network 988J

Tested At: 16 Mar 23 19:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 209 day(s)\
Subject: CN=SHAKEN Asia Pacific Network 988J, OU=APN, O=Asia Pacific Network, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Asia_Pacific_Network_988J

[View certificate details](https://understandingwebpki.com/?cert=MIIC2TCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTzEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjEwMloXDTIzMTAxMTE3MjEwMlowdTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMR0wGwYDVQQKDBRBc2lhIFBhY2lmaWMgTmV0d29yazEMMAoGA1UECwwDQVBOMSkwJwYDVQQDDCBTSEFLRU4gQXNpYSBQYWNpZmljIE5ldHdvcmsgOTg4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGprIYV6mubDNzZG2RPbgooSHOFxVCgV8IjCNPqJmSD47y2kcMzY2zCWJ40QKNL%2B4YYSUYw7WblAWG9N%2B%2FAMtq2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDk4OEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTaJdBkIP5plnGXTtM6JiPD6INMUTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgEEJl2UV3IuuAH%2FjjnBsmt9bW420in4yib6aQYsNaiAcCIQDz4VUJnRCLroUbaiKPKs9il8dxCVc6BrvXqdkKV6SxiQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 988J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 16 Mar 23 19:18 UTC