# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cascabel Networks 536K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 226 day(s)\
Subject: CN=SHAKEN Cascabel Networks 536K, OU=Cascabel, O=Cascabel Networks, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Cascabel_Networks_536K

[View certificate details](https://understandingwebpki.com/?cert=MIIC2DCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkduQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyOTAwMTUyMVoXDTI0MDYyODAwMTUyMVowdDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRowGAYDVQQKDBFDYXNjYWJlbCBOZXR3b3JrczERMA8GA1UECwwIQ2FzY2FiZWwxJjAkBgNVBAMMHVNIQUtFTiBDYXNjYWJlbCBOZXR3b3JrcyA1MzZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkh05tl8Gv96TLqiFj3Dpu3Eydb%2FKM7lNa4826Yyi1ClbtWDsGdtj6No30kMv3m15Ujx32mhy3WG2WrZaY4x2M6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTM2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFLXCXl3wxIGM6ePCNXL%2FzGL3WcIvMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiA90opUBoQ%2BQ2JoUd1qxd7WzoXiWQgZwKLND5MnmWH72QIhAJV5iAwDqrkZKxJ13BKMK8p%2FfP0WcM8wTLlHObK5vnba)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 536K' |


Generated: 15 Nov 23 16:51 UTC