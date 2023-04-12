# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Rayfield Communications, Inc. 006K

Tested At: 12 Apr 23 01:38 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 198 day(s)\
Subject: CN=SHAKEN Rayfield Communications\\, Inc. 006K, OU=VoIP, O=Rayfield Communications\\, Inc., ST=Missouri, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Rayfield_Communications_Inc._006K

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU0wwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5NDQwNloXDTIzMTAyNjE5NDQwNlowgYsxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNaXNzb3VyaTEmMCQGA1UECgwdUmF5ZmllbGQgQ29tbXVuaWNhdGlvbnMsIEluYy4xDTALBgNVBAsMBFZvSVAxMjAwBgNVBAMMKVNIQUtFTiBSYXlmaWVsZCBDb21tdW5pY2F0aW9ucywgSW5jLiAwMDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4hYpO8vVQiNpQnjWHQoE4hVTe1rj2SEhJKqNqLAnoFtLaBh9lMy9a58h3Rz0Tr%2FkFcvyaz%2FuTsiPDDn8lBVZL6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFCxjV%2Fl%2BVxoY14LRjiqPwwETJtBQMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAoy4%2FouWl2V7zYRwvfUZBvA9YuLik5dcUpuytt3Ygnv8CICpy1Rlfy1mABjyP%2FkMwtE1OT067jZZPCaFkxcg0OIjH)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 006K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 12 Apr 23 01:46 UTC