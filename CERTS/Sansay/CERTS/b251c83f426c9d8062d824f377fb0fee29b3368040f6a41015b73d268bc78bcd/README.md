# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Telxio Networks 492K

Tested At: 12 Apr 23 21:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=SHAKEN Telxio Networks 492K, OU=development, O=Telxio Networks, ST=Nevada, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ss-telxio

[View certificate details](https://understandingwebpki.com/?cert=MIIC2DCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZecwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyMDA0NTMwN1oXDTIzMDQxOTA0NTMwN1owdDELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk5ldmFkYTEYMBYGA1UECgwPVGVseGlvIE5ldHdvcmtzMRQwEgYDVQQLDAtkZXZlbG9wbWVudDEkMCIGA1UEAwwbU0hBS0VOIFRlbHhpbyBOZXR3b3JrcyA0OTJLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnF9Kveo1Pj4vdV%2BuN8zzNSzTT1SGTakC%2BbLBKaHPhI5RdhRtF4MJ%2FiuS3tEvS5mv5Hos1DSD1GPTs2k6pN44EqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENDkySzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFGmk0HW27VrCfU%2FRsi6ARZBbnNpUMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBNsqudvxBjqAVKZ4dIA2VFJFH%2FuFGrwtJMJQ%2BJEr5i1gIhALvq%2F8%2Bkj4sfK%2Bk8xHM9iPp6VQ96BFkmKfJc7m1gPieB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 492K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 12 Apr 23 22:02 UTC