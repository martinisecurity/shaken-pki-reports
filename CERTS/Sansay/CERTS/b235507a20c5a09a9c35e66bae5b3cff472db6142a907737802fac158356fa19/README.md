# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Convoso 758J

Tested At: 04 Oct 24 16:26 UTC\
Initial Validity Period: 35 day(s)\
Remaining Validity Period: 31 day(s)\
Subject: CN=SHAKEN Convoso 758J, O=Convoso, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://stirshaken.s3.us-west-1.amazonaws.com/stirshaken.pem

[View certificate details](https://x509.io/?cert=MIICtjCCAlygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkrAswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDkzMDExMDg1MVoXDTI0MTEwNDExMDg1MVowUjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEDAOBgNVBAoMB0NvbnZvc28xHDAaBgNVBAMME1NIQUtFTiBDb252b3NvIDc1OEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQr0l5ZxKiYcoEQtP8DiZgZ22gPBvqCK41AW85shlZGPWjj6HD%2Ffq0GCeb9vaOdVm4VW%2FoTE03pP2jgHWjNgSfco4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NThKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAdBgNVHQ4EFgQU%2FUmADFIbdJOik3WEEsHeQZ5jFlkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDHKsGEJ988jss%2FVZyNDoxyXXxD0JLEfQs3FHt7qNghsgIga4ybMYDVs%2BL3OeKZUaMk%2FRoTFNNGJix9OSGAOSCN9Rk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 758J', but common name is 'SHAKEN Convoso 758J' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC