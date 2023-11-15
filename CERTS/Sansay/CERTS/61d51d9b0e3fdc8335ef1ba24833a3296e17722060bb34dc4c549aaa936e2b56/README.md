# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: emailAddress=level5@lightspeedvoice.com, CN=SHAKEN Lightspeed Voice 557F, OU=Operations, O=Lightspeed Voice, ST=Florida, C=US, emailAddress=level5@lightspeedvoice.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/557F/order/8_557F_83

[View certificate details](https://understandingwebpki.com/?cert=MIIDBjCCAqygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhd8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwMTAyMjkyMloXDTIzMTIwMTAyMjkyMlowgaExCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRkwFwYDVQQKDBBMaWdodHNwZWVkIFZvaWNlMRMwEQYDVQQLDApPcGVyYXRpb25zMSUwIwYDVQQDDBxTSEFLRU4gTGlnaHRzcGVlZCBWb2ljZSA1NTdGMSkwJwYJKoZIhvcNAQkBFhpsZXZlbDVAbGlnaHRzcGVlZHZvaWNlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAV5dxgrmvqjF6d9WPJOH9JzeIB%2B1THxuLA7%2F4onIp1YJ8rlTFShsROqIG8pkbiN9XkrD2sI60iLab7cgIMlBnCjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU1N0YwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRXIoVYsMjvh7wpZ%2Fwm5cbD0NvR%2FTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgGSzT50hjIBJyrDQ2u4XV%2B%2FOMS0MSBJO5a%2Fivf3nzPNcCIQCnPQKL8qFwhSkNGSjRrSJIZZIhWRWHJ3dUH3ruj8Uekg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |


Generated: 15 Nov 23 17:17 UTC