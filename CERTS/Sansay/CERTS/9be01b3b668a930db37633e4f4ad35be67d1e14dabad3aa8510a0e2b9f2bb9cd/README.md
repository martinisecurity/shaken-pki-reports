# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 3058

Tested At: 21 Nov 22 20:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN Matrix 3058, OU=NOC, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Lingo-3058

[View certificate details](https://understandingwebpki.com/?cert=MIICvjCCAmOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVZEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMTM1NVoXDTIyMTIxNTIxMTM1NVowWTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxDDAKBgNVBAsMA05PQzEbMBkGA1UEAwwSU0hBS0VOIE1hdHJpeCAzMDU4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzUotEqptZc4M0%2Fx0MfF15yWz8%2Fntg4VDtUOPInSqs1AZUeQf2VfshGFf2LsxKq6gTQRS2LkM%2B79lJYpEztDLyKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzA1ODAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFAJFOlo0TAnsTivHURmXKZS0qDDWMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAp4xkC0lD%2FE5nSLZzrH%2BaljCTQqCA%2BwHHUB220LiJ90kCIQDJQH3M0bq5qaAdg%2BOR%2BpFbyRGpqG2evpNHFqzLWOta6w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3058' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Nov 22 20:33 UTC