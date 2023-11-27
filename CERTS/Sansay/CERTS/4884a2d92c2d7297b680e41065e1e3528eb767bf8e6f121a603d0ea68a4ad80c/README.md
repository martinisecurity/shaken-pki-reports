# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 27 Nov 23 23:19 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 11 day(s)\
Subject: emailAddress=level5@lightspeedvoice.com, CN=SHAKEN Lightspeed Voice 557F, OU=Operations, O=Lightspeed Voice, ST=Florida, C=US, emailAddress=level5@lightspeedvoice.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/557F/order/15_557F_83

[View certificate details](https://understandingwebpki.com/?cert=MIIDBjCCAqygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhhcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwODAxNTQyMloXDTIzMTIwODAxNTQyMlowgaExCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRkwFwYDVQQKDBBMaWdodHNwZWVkIFZvaWNlMRMwEQYDVQQLDApPcGVyYXRpb25zMSUwIwYDVQQDDBxTSEFLRU4gTGlnaHRzcGVlZCBWb2ljZSA1NTdGMSkwJwYJKoZIhvcNAQkBFhpsZXZlbDVAbGlnaHRzcGVlZHZvaWNlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAV5dxgrmvqjF6d9WPJOH9JzeIB%2B1THxuLA7%2F4onIp1YJ8rlTFShsROqIG8pkbiN9XkrD2sI60iLab7cgIMlBnCjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU1N0YwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRXIoVYsMjvh7wpZ%2Fwm5cbD0NvR%2FTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgbh6KofoIzMF4UjIgw4nCWVVoOsNysMJ4HFoDOBLYIYkCIQDTzxRGDoKldVfWGflVicREpZIBdNzEinXLtrvl%2B8dYHQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F', but common name is 'SHAKEN Lightspeed Voice 557F' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 27 Nov 23 23:28 UTC