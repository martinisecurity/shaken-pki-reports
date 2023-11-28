# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 594J

Tested At: 28 Nov 23 10:19 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -429 day(s)\
Subject: CN=SHAKEN 594J, OU=SHAKEN, O=Carolina Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/14652271-f2ba-4197-87c0-704f8c618e4b/3f9a5e102a5aea8ae1b91aa8c557c3b3.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApigAwIBAgIQR26H7a0c2q3CDUAuqzHFjDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDE2MDlaFw0yMjA5MjQyMDE2MDhaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBDYXJvbGluYSBEaWdpdGFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1OTRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtiTvOkkM9LL02CAmGtfctBSi3JWME2Wx8mxDTgA17TvztaW4ViNiHbyqZYpAzIH%2Fx4UnV047SFN%2F1UAntHhG86OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTAAdyn%2BrYheWIM6V%2FgGzVn85xBODAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1OTRKMAoGCCqGSM49BAMCA0cAMEQCICLSU%2BfwAk7zX0lB8aJF4lvs6MA7P%2FlS1BtcL7IEngfVAiBD2ZI1SD%2BlqEI2tJMf%2BlE6dUkLcuXrPvUktm9M0Ydpfw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC