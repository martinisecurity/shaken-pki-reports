# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 1417

Tested At: 31 Oct 22 20:40 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Momentum Telecom 1417, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MomentumTelecom_1417

View: [Click to view](https://understandingwebpki.com/?cert=MIIC0zCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTzMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjEyOFoXDTIyMTExMDE3MjEyOFowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gMTQxNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABII7HUTol1f%2FqSE%2B%2Faw3uBpttsdkbcNLNe4GTT6QNc1p5xhfUjCm8EJ1xtaeemvakN4DOVzQkrxtJiYyjdmnnLSjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE0MTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT5vvlSdrpNc3W4WbWUJC6Zwu8MZDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgBV8zkNc6Goyvh6bcgGnoQNJMO2nJ1EsBGUlqQyUsVXICIQDaVA9zrGdNxrKueMcat%2BxFlx%2FgdyyeuKU7kcgFrJr0tA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 1417' |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31/10/2022 at 20:47:45