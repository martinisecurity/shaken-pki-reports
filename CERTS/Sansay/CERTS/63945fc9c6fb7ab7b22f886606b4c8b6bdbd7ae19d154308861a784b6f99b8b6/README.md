# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN VOIP OFFICE.COM LLC 389K

Tested At: 15 Nov 23 16:36 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -19 day(s)\
Subject: CN=SHAKEN VOIP OFFICE.COM LLC 389K, OU=Communications, O=VOIP OFFICE, ST=Michigan, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://nguc.voipoffice.com/stirshaken/VOIP_OFFICE.COM_LLC_389K

[View certificate details](https://understandingwebpki.com/?cert=MIIC3jCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU1kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5NTQwNVoXDTIzMTAyNjE5NTQwNVoweTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE1pY2hpZ2FuMRQwEgYDVQQKDAtWT0lQIE9GRklDRTEXMBUGA1UECwwOQ29tbXVuaWNhdGlvbnMxKDAmBgNVBAMMH1NIQUtFTiBWT0lQIE9GRklDRS5DT00gTExDIDM4OUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASyNJX%2FUt8BC24kxUNOea06PnB5syn2R%2FNwkXPl1D6C1RrD5Y1QmvcytXYJGLBp9g8A7NrZr6YDRFOGY9Jrk9J4o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzODlLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpoZG1j2PdPoEsNQ5EWxpuSYUQ1YwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQD0WtbMmhfRuuv1Bne4IS7hQsaQ9y5HLHxHQw7qU7zr0AIhANnFJYjB4ks54Hsvtpc1ydUbrbFwicawXUkhuGON7XOK)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 389K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 17:17 UTC