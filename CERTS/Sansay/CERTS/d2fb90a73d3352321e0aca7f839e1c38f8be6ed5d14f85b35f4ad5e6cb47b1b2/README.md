# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1stPoint Communications, LLC 463G

Tested At: 04 Oct 24 16:02 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN 1stPoint Communications\\, LLC 463G, O=1stPoint Communications\\, LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/463G/429C7C70711E3820F0B8E1DEAE6FF3262264A94C.pem

[View certificate details](https://x509.io/?cert=MIIC3jCCAoSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkqUwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDkwNjAzMzkwNVoXDTI0MTAwNjAzMzkwNVowejELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMSUwIwYDVQQKDBwxc3RQb2ludCBDb21tdW5pY2F0aW9ucywgTExDMTEwLwYDVQQDDChTSEFLRU4gMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQyA0NjNHMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAELs6zRwpFnh1RXoHhA%2Fxf4rfKjbd9Nfq3eUs2D8NO%2Fad5EI%2BrHaj5kzwlXBLGPguDJcIGUC4kVK9Kg1HavwIwRqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENDYzRzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFM81qK%2Bd66zq6IUVZdoP8XnclRGbMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAucuUYvzuddEKR23ES8xB2%2BB1WIjXcbqY9xY6aKTbMYACIHm1Wk2D9hZW8XA4O6qA5N8eRNVG1rTR%2FzDGWqcJm%2FUC)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463G', but common name is 'SHAKEN 1stPoint Communications, LLC 463G' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC