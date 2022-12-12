# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Technology Innovation Lab 599J

Tested At: 12 Dec 22 23:38 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 303 day(s)\
Subject: CN=SHAKEN Technology Innovation Lab 599J, OU=Research and Development, O=Technology Innovation Lab, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Technology_Innovation_Lab_599J

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FTCCAqKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTxIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTQyNloXDTIzMTAxMTE3MTQyNlowgZcxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEiMCAGA1UECgwZVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYjEhMB8GA1UECwwYUmVzZWFyY2ggYW5kIERldmVsb3BtZW50MS4wLAYDVQQDDCVTSEFLRU4gVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYiA1OTlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEp%2FXuIMYmugxcCt65pBaM1dcrsbEtioYpHt3nRkRBelXmNqq%2FjyUOMMiocVy2sYCLecTaT48odMFBVHHx3pa0VaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFJ04%2F3tOUN8HAfR7HWWhha7kRtsFMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA1x%2BIMfol09%2FR8AzeD2hm2SV%2FB6%2FnlBruckEu3uRowucCIQCynbPty%2BhtRdbKituQzpeS9xpuukcQQIggXwkKPxADbw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 599J' |


Generated: 12 Dec 22 23:45 UTC