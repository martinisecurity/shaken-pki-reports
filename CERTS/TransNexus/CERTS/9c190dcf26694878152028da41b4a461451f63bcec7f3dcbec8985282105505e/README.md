# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 952J

Tested At: 04 Oct 24 15:50 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -709 day(s)\
Subject: CN=SHAKEN 952J, OU=SHAKEN, O=Syndeo LLC dba Broadvoice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/952J/7787f5d6-eb67-4354-b625-597d01705d8b.pem

[View certificate details](https://x509.io/?cert=MIIC%2BjCCAqGgAwIBAgIQZJHdBgTOLUEYkScs4FjeTTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjUxNzU0MzhaFw0yMjEwMjUxNzU0MzdaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlTeW5kZW8gTExDIGRiYSBCcm9hZHZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA5NTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnrknl5CabCS7bMbBU6hrZiJhSY%2BeeSfyRuM3nhFzM9p3GsBXWjYu4eY9RsY41SvsXKl6Y%2FFJzQw75IaNTLtZBKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSnDlmcb2HeAONNI62tp2d2a62T5DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTJKMAoGCCqGSM49BAMCA0cAMEQCIDmsQ4MCD4jRWUlYxUBsae8ojQ3Yp6MgtPZHY3vKPnYhAiAtCxZoh5fPRbFKNVMXKbZraLGLAtT%2FAi%2FloUuI1aLsAg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC