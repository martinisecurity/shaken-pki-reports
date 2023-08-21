# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN GIP Technology 434K

Tested At: 21 Aug 23 20:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 247 day(s)\
Subject: CN=SHAKEN GIP Technology 434K, OU=GIPT, O=GIP Technology, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://gip.46labs.com

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkbKowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQyNTE2NDM0MloXDTI0MDQyNDE2NDM0MlowbzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFzAVBgNVBAoMDkdJUCBUZWNobm9sb2d5MQ0wCwYDVQQLDARHSVBUMSMwIQYDVQQDDBpTSEFLRU4gR0lQIFRlY2hub2xvZ3kgNDM0SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBT4LNy3yTOc1Djbu5Wr%2FxbWVEj3HfO5hrPoISMOIElWo%2BUMeaOMbQLNqiKqUpOFMle1YpulbUGdGnWNiLr9alyjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQzNEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRVTHMXP2axiTd0Q3OP%2FQDNwZtk6zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAI4GeYvkfjoDWuLmGaSxj8%2FJ7hCtKhdc7rmeaewhjfQ2AiBCRqFbD1mOJquCnyjl1Sxpq3svjSswAIOGSpX4M60vIA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 434K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC