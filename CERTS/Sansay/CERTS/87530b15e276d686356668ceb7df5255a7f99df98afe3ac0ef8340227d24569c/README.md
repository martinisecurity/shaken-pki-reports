# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN DLS Internet Services 815J

Tested At: 26 Aug 24 18:30 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -12 day(s)\
Subject: CN=SHAKEN DLS Internet Services 815J, O=DLS Internet Services, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/815J/429C7C70711E3820F0B8E1DEAE6FF3262264A37A.pem

[View certificate details](https://x509.io/?cert=MIIC0DCCAnagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJko3owCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcxNTExMzMwN1oXDTI0MDgxNDExMzMwN1owbDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMR4wHAYDVQQKDBVETFMgSW50ZXJuZXQgU2VydmljZXMxKjAoBgNVBAMMIVNIQUtFTiBETFMgSW50ZXJuZXQgU2VydmljZXMgODE1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD8vmDjJBJptvJN6T9pNAzcPO7%2FqgWaA8sqf3L8b%2FMPDM7kz2hvhTae5yKoP0gRR6mGeUntjaMU7kBj7r0O3CSyjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDgxNUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBTTDJEr5UuVXXUerQlSPZ%2FYy7WAgTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMLo4ZnSJyMRfNLAyXdb02dI1iDcqFJ7mANrRl4%2Fai18AiAKx2G%2BWZSADeD4hDyloA1zGIh7t9ao3aM33%2B0Az22n9g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 815J', but common name is 'SHAKEN DLS Internet Services 815J' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 26 Aug 24 18:49 UTC