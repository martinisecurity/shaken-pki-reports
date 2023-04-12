# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 9157

Tested At: 12 Apr 23 21:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN Momentum Telecom 9157, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MomentumTelecom_9157

[View certificate details](https://understandingwebpki.com/?cert=MIIC0jCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkaiwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQxMTEzNDY1NloXDTIzMDUxMTEzNDY1NlowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gOTE1NzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMAoR2hhlAdFCEs0FU4ijCdcMKOr8L38LGFGNqTxndGpUVfQTjvjgl1cFmyFxp2fb1vOAVoMqT4Q5Tz8Br3O%2BQejgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkxNTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQ%2Fof%2BXyLnO4yitFukJpfvLLFMnQzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgFJvA%2BaCURpwQLDtO4PNSvg%2FmX3zYDekgI0whdmfPOSMCIAngXHxvnvguSbHv0hay6ldBKWf5A%2F8R%2FkZjz76U87gL)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9157' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC