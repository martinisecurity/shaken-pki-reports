# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ringfree Communications Inc 317K

Tested At: 20 Nov 22 22:56 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 325 day(s)\
Subject: CN=SHAKEN Ringfree Communications Inc 317K, OU=NOC, O=Ringfree Communications Inc, ST=North Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/RingFree_Communications_317K

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT0EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjM0NloXDTIzMTAxMTE3MjM0NlowgYwxCzAJBgNVBAYTAlVTMRcwFQYDVQQIDA5Ob3J0aCBDYXJvbGluYTEkMCIGA1UECgwbUmluZ2ZyZWUgQ29tbXVuaWNhdGlvbnMgSW5jMQwwCgYDVQQLDANOT0MxMDAuBgNVBAMMJ1NIQUtFTiBSaW5nZnJlZSBDb21tdW5pY2F0aW9ucyBJbmMgMzE3SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABH4R1r%2BgajYy1I3OidPuhqS1MHvRrhq8HC0kc3Qo%2F2gzwdivAutf%2Bo0yFx%2BaVTp4CNgSCvBTKA3berayC9fI3A6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxN0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRvR8WGcaTRo0hvqEbV08tv6vGx%2BDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAKLvwNQww3rjIZsFVms4%2FX0wRVuJOlBYBGubTok8Cq%2B%2FAiEA4YE3vnnD5o1PMjG82NpNVYMTvrNtffLAdJaV0xclBI8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 317K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 20 Nov 22 22:57 UTC