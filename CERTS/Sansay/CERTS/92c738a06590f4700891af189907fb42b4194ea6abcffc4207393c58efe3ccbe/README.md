# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Convoso 758J

Tested At: 05 Apr 24 19:00 UTC\
Initial Validity Period: 35 day(s)\
Remaining Validity Period: 12 day(s)\
Subject: CN=SHAKEN Convoso 758J, O=Convoso, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://stirshaken.s3.us-west-1.amazonaws.com/stirshaken.pem

[View certificate details](https://x509.io/?cert=MIICtTCCAlygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJklUUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDMxMzA2MDAxMVoXDTI0MDQxNzA2MDAxMVowUjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEDAOBgNVBAoMB0NvbnZvc28xHDAaBgNVBAMME1NIQUtFTiBDb252b3NvIDc1OEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQr0l5ZxKiYcoEQtP8DiZgZ22gPBvqCK41AW85shlZGPWjj6HD%2Ffq0GCeb9vaOdVm4VW%2FoTE03pP2jgHWjNgSfco4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NThKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAdBgNVHQ4EFgQU%2FUmADFIbdJOik3WEEsHeQZ5jFlkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBr3xMjq8GP4obgGbd6fMZW%2F8DcVvmkhCHbzCG%2FGWHOyAiAvcDj%2FGegjfzrwZTmhXZq7ontvROTfA78U8x7fAIxJ7w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 758J', but common name is 'SHAKEN Convoso 758J' |


Generated: 05 Apr 24 19:04 UTC