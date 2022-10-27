# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 9ab28ecee0c6785a7aa07eac83709ae3b5cc6116
Tested At: 2022-10-27 18:56:34 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 349 day(s)\
Subject: CN=SHAKEN Technology Innovation Lab 599J, OU=Research and Development, O=Technology Innovation Lab, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Technology_Innovation_Lab_599J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FTCCAqKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTxIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTQyNloXDTIzMTAxMTE3MTQyNlowgZcxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEiMCAGA1UECgwZVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYjEhMB8GA1UECwwYUmVzZWFyY2ggYW5kIERldmVsb3BtZW50MS4wLAYDVQQDDCVTSEFLRU4gVGVjaG5vbG9neSBJbm5vdmF0aW9uIExhYiA1OTlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEp%2FXuIMYmugxcCt65pBaM1dcrsbEtioYpHt3nRkRBelXmNqq%2FjyUOMMiocVy2sYCLecTaT48odMFBVHHx3pa0VaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk5SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFJ04%2F3tOUN8HAfR7HWWhha7kRtsFMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA1x%2BIMfol09%2FR8AzeD2hm2SV%2FB6%2FnlBruckEu3uRowucCIQCynbPty%2BhtRdbKituQzpeS9xpuukcQQIggXwkKPxADbw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 599J' |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 18:57:26