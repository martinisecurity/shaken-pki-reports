# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Doylestown Communications, Inc 849C

Tested At: 04 Oct 24 16:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -107 day(s)\
Subject: CN=SHAKEN Doylestown Communications\\, Inc 849C, O=Doylestown Communications\\, Inc, ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/849C/429C7C70711E3820F0B8E1DEAE6FF32622649CE2.pem

[View certificate details](https://x509.io/?cert=MIIC3TCCAoSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknOIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUxOTE4MTEzNVoXDTI0MDYxODE4MTEzNVowejELMAkGA1UEBhMCVVMxDTALBgNVBAgMBE9oaW8xJzAlBgNVBAoMHkRveWxlc3Rvd24gQ29tbXVuaWNhdGlvbnMsIEluYzEzMDEGA1UEAwwqU0hBS0VOIERveWxlc3Rvd24gQ29tbXVuaWNhdGlvbnMsIEluYyA4NDlDMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAES7%2FVcVi5S%2FfjFrz9TTX%2Faq3nf7Xx4I1tRKO5t%2BTzny%2BLgzMfAEnIJykJyJarb3HdCsrefb1mJ1HvloISjRemdKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODQ5QzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFMsEDZ37F%2BvjOHVFoUTu3q5TinzXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAtuzuwvMJYQUgguX609VBXPuKXNB%2Fh3Bemw7dk42abQgIgK%2F0cYHaNBVPpH9MxSsMN%2BblCqIQmXsGANcLI81%2BD92w%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 849C', but common name is 'SHAKEN Doylestown Communications, Inc 849C' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC