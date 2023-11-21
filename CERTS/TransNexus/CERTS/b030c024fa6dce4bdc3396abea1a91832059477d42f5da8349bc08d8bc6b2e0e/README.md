# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 060K

Tested At: 21 Nov 23 01:23 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -377 day(s)\
Subject: CN=SHAKEN 060K, OU=SHAKEN, O=Telware, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0c4e235f-3e3e-4dd3-bbfc-2f15badab180/9719c0e4dbfe850baf03f6c3b5cb85b6.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIQSNf7OoaMSyTvAVl%2B0S9SyTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDExODE0MDVaFw0yMjExMDgxODE0MDRaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdUZWx3YXJlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNjBLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEM%2FN43HcA3aBnvE4RZR1FmF4e3dkT%2F7FoJpI27RuwB36qcwh8cUbgvVb3oY9st3ea8uyIFx9AwqfjfnYYv1DLEqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTe8Cyyot9ChdoRL5GjXJal2VO%2FgTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNjBLMAoGCCqGSM49BAMCA0kAMEYCIQChEjnSDKINt5CQBxPxKw%2BEuKiqfOWlQwFwZsRU%2BJ1bWwIhAKQMi4vQ6kcK2Arryqx2%2FNQSgNPmiqSikYjOLm7NU3Ek)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC