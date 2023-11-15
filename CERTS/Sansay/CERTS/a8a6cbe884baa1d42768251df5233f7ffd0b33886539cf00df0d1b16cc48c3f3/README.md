# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ONE OWL TELECOM INC 412K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 364 day(s)\
Subject: CN=SHAKEN ONE OWL TELECOM INC 412K, OU=ONEOWLTELECOM, O=ONE OWL TELECOM INC, ST=Maryland, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/OneOwlTelecom_412K

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhkAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzgxMloXDTI0MTExMzEzMzgxMlowgYAxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNYXJ5bGFuZDEcMBoGA1UECgwTT05FIE9XTCBURUxFQ09NIElOQzEWMBQGA1UECwwNT05FT1dMVEVMRUNPTTEoMCYGA1UEAwwfU0hBS0VOIE9ORSBPV0wgVEVMRUNPTSBJTkMgNDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKEvu4fMwn4fvF%2B%2BkuqZFBavh4MF7%2BTI1J18I7lxtsDUcIqikGFk5PHCAuGheBcM9J4nUUCi%2F2qkVcYl0vvsEE2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQxMkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRr1XzKQRLwXA4AN%2B3CbBTFpb5cADAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgXoGNa%2FAQcu6n1LfKRS3omFNoRQs%2Bg5TVC0iM2IsXHOoCIC3YUTMeRAQwcEuixSKPayvLlgWdgK8i4CSChrPNV1ts)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 412K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC