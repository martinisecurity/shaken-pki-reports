# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 735J

Tested At: 27 Nov 23 22:22 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -399 day(s)\
Subject: CN=SHAKEN 735J, OU=SHAKEN, O=AVOXI Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/11124177-79f3-48b2-867a-386d4dc61c99/a094dbb369b0daa6b826aec6227eef7d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApKgAwIBAgIQc%2FNJOMEjy7rjruwsHS51OzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDE3NThaFw0yMjEwMjQyMDE3NTdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpBVk9YSSBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJUY%2B5zIUpeVNAIi6OK1orE3lI%2Fc6jBPK7elsughH30c9O1xWTj%2B67VlRPh%2BJ9AtNZy1ue2INL%2FLytw5u%2FGQcl6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSOP%2BEWcwezRcMTe7VWUjJ0iYn7VjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzVKMAoGCCqGSM49BAMCA0cAMEQCID0yFIyw2rDAI0lcsFWKbWLOutJGBCGLBYYuwNWsuXbZAiB5NPbedOLbFs97Y3YkH0MIVkgyBZuJSWpxRaBK4PwmGA%3D%3D)

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


Generated: 27 Nov 23 22:56 UTC