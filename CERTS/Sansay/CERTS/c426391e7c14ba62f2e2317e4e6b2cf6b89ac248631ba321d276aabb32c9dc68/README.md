# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN SouthPoint Communications 205k

Tested At: 18 Aug 25 20:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -308 day(s)\
Subject: CN=SHAKEN SouthPoint Communications 205k, O=SouthPoint Communications, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/205k/429C7C70711E3820F0B8E1DEAE6FF3262264AA53.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkqlMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDkxNDExMzgxMFoXDTI0MTAxNDExMzgxMFowdjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExIjAgBgNVBAoMGVNvdXRoUG9pbnQgQ29tbXVuaWNhdGlvbnMxLjAsBgNVBAMMJVNIQUtFTiBTb3V0aFBvaW50IENvbW11bmljYXRpb25zIDIwNWswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASfJ3OhB3xnPprzTOKiNwTougqlsCZMeV6M0kcwk8F%2BAkWyorpqFfoEGE65rnxAU2fZJu74W5lCPZtiPEyGdxs1o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMDVrMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAdBgNVHQ4EFgQU9%2BIQO2rwmus8MaQutkdnPhX%2F4aswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCOeO768V04DR8rCosUHGEh72GlBDB89AdpJoPOdaIULwIgedjDWm1uFT9wzHe34SlRiVpfkIXNuoxJZLgdxCt%2BSZw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '205k' contains characters other than uppercase letters and numbers |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 205k', but common name is 'SHAKEN SouthPoint Communications 205k' |


Generated: 18 Aug 25 21:13 UTC