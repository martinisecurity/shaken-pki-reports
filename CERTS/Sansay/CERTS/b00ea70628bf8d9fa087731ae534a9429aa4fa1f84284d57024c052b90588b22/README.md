# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ONE OWL TELECOM INC 412K

Tested At: 12 Apr 23 01:38 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -23 day(s)\
Subject: CN=SHAKEN ONE OWL TELECOM INC 412K, OU=ONEOWLTELECOM, O=ONE OWL TELECOM INC, ST=Maryland, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ONE_OWL_TELECOM_INC_412K

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYb4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxNzEwMDA0NVoXDTIzMDMxOTEwMDA0NVowgYAxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNYXJ5bGFuZDEcMBoGA1UECgwTT05FIE9XTCBURUxFQ09NIElOQzEWMBQGA1UECwwNT05FT1dMVEVMRUNPTTEoMCYGA1UEAwwfU0hBS0VOIE9ORSBPV0wgVEVMRUNPTSBJTkMgNDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABILltt1rRucTYd1VEjaGA689%2Bhvtph8O496CsQ2m7YmUbAGOG07vJDwBHhtG68Ql6I33Gs4fXjY4aN1MQyKX01CjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQxMkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQ53VJLBv%2FgIdSzL42m4k3O%2F0kxfjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAOytvtYyqbV0nTrKiAM9b%2BBauB41q%2BvDO67XPyhy7abcAiAFv2tBEY6TL4YB8Cg6Fk12YZx4UN8D3FIhhaL6geWMSg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 412K' |


Generated: 12 Apr 23 01:46 UTC