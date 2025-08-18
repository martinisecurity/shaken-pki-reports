# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InteractiveTel, LLC 920J

Tested At: 18 Aug 25 20:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -12 day(s)\
Subject: CN=SHAKEN InteractiveTel\\, LLC 920J, O=InteractiveTel\\, LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/920J/429C7C70711E3820F0B8E1DEAE6FF3262264D38B.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJk04swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI1MDcwNzEzMTIyNFoXDTI1MDgwNjEzMTIyNFowZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRwwGgYDVQQKDBNJbnRlcmFjdGl2ZVRlbCwgTExDMSgwJgYDVQQDDB9TSEFLRU4gSW50ZXJhY3RpdmVUZWwsIExMQyA5MjBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEjHKQ3%2FT360Cb2diP9uObh0JqxxA1R5Dzno2vLhHsihhkJgjuWfiMnnNcp05LFh73ggwquASoWuxbeLEcecZi1KOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTIwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFEzpkbnSJVNZqafxkTPtmsYetg7aMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAj1BmGArtBz8zEEGITdGgx1%2BFTCT337rd1Ztk%2B5ri7qcCIBVdGnrkxZmsrRUMnA2wxYyB3vtUAZW1PrsjoIXaQXHc)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 920J', but common name is 'SHAKEN InteractiveTel, LLC 920J' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 18 Aug 25 21:13 UTC