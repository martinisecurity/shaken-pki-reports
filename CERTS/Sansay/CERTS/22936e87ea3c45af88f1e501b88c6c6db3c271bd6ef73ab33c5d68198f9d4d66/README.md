# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Vinculum Communications, Inc 787J

Tested At: 07 Jan 23 19:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 293 day(s)\
Subject: CN=SHAKEN Vinculum Communications\\, Inc 787J, OU=Engineering, O=Vinculum Communications\\, Inc, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Vinculum_Communications_Inc_787J

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCAp2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUz8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5MzQ0OVoXDTIzMTAyNjE5MzQ0OVowgZIxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSUwIwYDVQQKDBxWaW5jdWx1bSBDb21tdW5pY2F0aW9ucywgSW5jMRQwEgYDVQQLDAtFbmdpbmVlcmluZzExMC8GA1UEAwwoU0hBS0VOIFZpbmN1bHVtIENvbW11bmljYXRpb25zLCBJbmMgNzg3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHh4UmWHkgS57FSPbP6wU9HpEntNVUnHkLBqa6sQxLzCmm3Ik2zip10VJb2d0ts5EbzcblqdSN0uUkXp%2FFiYq%2FKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDc4N0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTf1pldNFylrskTEayXx5bTaq6GnjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgDQL%2BE2mhQJ4yxeCEPSrvBztvxzexNVQTKX0UwBIMzq0CIGEmRXd91AKYLUdEYtirY9NPUcFyA%2Bkp3Mn%2FgeANeBTp)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 787J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 07 Jan 23 19:18 UTC