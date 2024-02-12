# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPSBS Managed Services LLC 828J

Tested At: 12 Feb 24 16:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 19 day(s)\
Subject: CN=SHAKEN IPSBS Managed Services LLC 828J, O=IPSBS Managed Services LLC, ST=Tennessee, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/828J/429C7C70711E3820F0B8E1DEAE6FF3262264901D.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC2zCCAoGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkkB0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEzMTIyMzgyM1oXDTI0MDMwMTIyMzgyM1owdzELMAkGA1UEBhMCVVMxEjAQBgNVBAgMCVRlbm5lc3NlZTEjMCEGA1UECgwaSVBTQlMgTWFuYWdlZCBTZXJ2aWNlcyBMTEMxLzAtBgNVBAMMJlNIQUtFTiBJUFNCUyBNYW5hZ2VkIFNlcnZpY2VzIExMQyA4MjhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDShEx4cnZ31MAysnbxMraQzdzibEL%2BeVsH4G033WsSQe8ZhzcdJbLmeANlhKpvGi34rAYld4bx9TYd4VQ7cxP6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEODI4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwHQYDVR0OBBYEFELysdZtK9UtpJTfucKRMe8KYe8XMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBZ2gVEifUrmrlflaBdocot4DnVnQK0Ez1beCs%2FLbrGpQIhAMH0lKNOSOB48lcNGVXNIm4QCUGnhMbguck%2BTBkFmr39)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 828J', but common name is 'SHAKEN IPSBS Managed Services LLC 828J' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 12 Feb 24 17:02 UTC