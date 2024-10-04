# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mango Voice LLC 579K

Tested At: 04 Oct 24 16:06 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -7 day(s)\
Subject: CN=SHAKEN Mango Voice LLC 579K, O=Mango Voice LLC, ST=Utah, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/579K/429C7C70711E3820F0B8E1DEAE6FF3262264A85E.pem

[View certificate details](https://x509.io/?cert=MIICwDCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkqF4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDgyODE1MTQzN1oXDTI0MDkyNzE1MTQzN1owXDELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxGDAWBgNVBAoMD01hbmdvIFZvaWNlIExMQzEkMCIGA1UEAwwbU0hBS0VOIE1hbmdvIFZvaWNlIExMQyA1NzlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEN96Dvg9T%2FYiZ4tCcDVBdaWRfr9zU%2FoiNTef6qXh499UHoCaSIgRwPlHS8kuaEAr3W4pT2vsx%2FtcEJPEuvfvks6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTc5SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFEMNbo0EEdDml2YVgLkg9Zx8259GMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAze3EPpw8duFFy3IHhGhiLQJBEf1D94as%2FDZDNstiCVYCIHgp6mzY3xiTYhnNkM4z9%2B3VPH3E3zxNQu6f0jZg9Mnh)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 579K', but common name is 'SHAKEN Mango Voice LLC 579K' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC