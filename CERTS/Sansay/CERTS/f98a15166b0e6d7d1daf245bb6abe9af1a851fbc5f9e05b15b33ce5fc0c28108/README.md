# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 1417

Tested At: 11 Jan 23 21:50 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: CN=SHAKEN Momentum Telecom 1417, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MomentumTelecom_1417

[View certificate details](https://understandingwebpki.com/?cert=MIIC0jCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWrswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEwMzIxMDI0NVoXDTIzMDIwMjIxMDI0NVowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gMTQxNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABII7HUTol1f%2FqSE%2B%2Faw3uBpttsdkbcNLNe4GTT6QNc1p5xhfUjCm8EJ1xtaeemvakN4DOVzQkrxtJiYyjdmnnLSjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE0MTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT5vvlSdrpNc3W4WbWUJC6Zwu8MZDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgE%2FoFRQxm%2BAoKO%2FLhwvZJR2%2FpvVi8AXKmxyyiYWPQpQoCIDoLgqyeThu7lxQrkfADwK8XwbFMJoPMQe4xmJ%2BcDzsj)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1417' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 11 Jan 23 21:59 UTC