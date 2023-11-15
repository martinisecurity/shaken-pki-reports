# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Talk IT Pro 321K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 338 day(s)\
Subject: CN=SHAKEN Talk IT Pro 321K, OU=Talk IT Pro, O=Talk IT Pro, ST=Michigan, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Talk_IT_Pro_321K102522

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhMQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxOTExNDM1MloXDTI0MTAxODExNDM1MlowbjELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE1pY2hpZ2FuMRQwEgYDVQQKDAtUYWxrIElUIFBybzEUMBIGA1UECwwLVGFsayBJVCBQcm8xIDAeBgNVBAMMF1NIQUtFTiBUYWxrIElUIFBybyAzMjFLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEegO82HZ0pOrBUxOmMSn45mfmADCPaSCy32xBAMjc1QQuBNGWZvnAg2XL8Wr5zSnMFJ6B4LMs0WzEyYvvzqWOr6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzIxSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHO%2FGMfKSzzhjVVM%2FtsHwAL%2FFnIEMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAtv2OR3qzEj5nchrUv%2FFgkorPgq8TPsmqeGoC5qIs2zMCIQC8dm7Y8DYGl8vPez7YU0yaxvVnZOmNzoOCHGyDvTrykA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 321K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 16:51 UTC