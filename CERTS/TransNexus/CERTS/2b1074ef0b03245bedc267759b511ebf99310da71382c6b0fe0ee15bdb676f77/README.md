# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 21 Nov 23 16:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -374 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SIP.US LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2185f3c0-09fe-4171-ac47-5f28eb0ef48e/c1df43e07ecff89b5a07e9842f5a4b2e.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQYdy27VW2Zp8ZuTNCK6uf7TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDQyMDE5NDFaFw0yMjExMTEyMDE5NDBaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpTSVAuVVMgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkjRjiGbzvwRnujkkp7CpQ%2BprOwmNMzsczKzAIsDOjYaMgKEwd5OQ%2Bo0x5cMHHdUDlZuYxBuu9waGaWwDF7GcyKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQZToQrnVVP%2FNPwgpW3nV%2BkNiGGPTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzhKMAoGCCqGSM49BAMCA0kAMEYCIQCtpy2PzOnA7Wei4SuP3pQH7%2B9VcU%2BFChszQWHsY0D7YQIhAKObce1%2B8jsMv7mxqDT2FtwaNkREuMw8ulVFFSCZZ%2B1m)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC