# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 22 Aug 24 15:31 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -12 day(s)\
Subject: CN=SHAKEN Lightspeed Voice 557F, O=Lightspeed Voice, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/557F/429C7C70711E3820F0B8E1DEAE6FF3262264A2FD.pem

[View certificate details](https://x509.io/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkov0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcxMDE3NDA1MVoXDTI0MDgwOTE3NDA1MVowYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGTAXBgNVBAoMEExpZ2h0c3BlZWQgVm9pY2UxJTAjBgNVBAMMHFNIQUtFTiBMaWdodHNwZWVkIFZvaWNlIDU1N0YwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQFeXcYK5r6oxenfVjyTh%2FSc3iAftUx8biwO%2F%2BKJyKdWCfK5UxUobETqiBvKZG4jfV5Kw9rCOtIi2m%2B3ICDJQZwo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTdGMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAdBgNVHQ4EFgQUVyKFWLDI74e8KWf8JuXGw9Db0f0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIHz%2Fw2JOFGEz%2BCRswduZRdM%2FDgPMx%2B0VqSNNreZXkVtBAiEAx3qNCvydRP4YNUSCvEz2zIuaubKT%2FqKdBNwc2S%2Fxm7M%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F', but common name is 'SHAKEN Lightspeed Voice 557F' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 22 Aug 24 15:44 UTC