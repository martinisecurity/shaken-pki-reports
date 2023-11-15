# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ACS Technologies 488K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 79 day(s)\
Subject: CN=SHAKEN ACS Technologies 488K, OU=ACS Technologies, O=ACS Technologies, ST=South Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ACS_Technologies_488K

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYEEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIwMTIwNTEyOFoXDTI0MDIwMTIwNTEyOFowgYMxCzAJBgNVBAYTAlVTMRcwFQYDVQQIDA5Tb3V0aCBDYXJvbGluYTEZMBcGA1UECgwQQUNTIFRlY2hub2xvZ2llczEZMBcGA1UECwwQQUNTIFRlY2hub2xvZ2llczElMCMGA1UEAwwcU0hBS0VOIEFDUyBUZWNobm9sb2dpZXMgNDg4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFMe9YwMofpHj%2F50MuZFl7yLQATMdgdAhuqq2qwLR19fK2UNrCkFqpxTLfbOBgBLlNAVLCyyePfpFk07nFEgbSOjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ4OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTbQUsyrlcmqIln2ps%2B8zHK2r0EpTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAPJiFVf8VI5Gk42LeSziSs1stfikG3mBcR1OGLbNzq5gAiA6m7W0915ILnNCXNtAVHvwPLzZBtL1O6Rw4IJjaU%2FMCQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 488K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC