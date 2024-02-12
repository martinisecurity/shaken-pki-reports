# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 518J

Tested At: 12 Feb 24 16:31 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -478 day(s)\
Subject: CN=SHAKEN 518J, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/transnexus/028fea0e03c28db631b53e6ebffddbe2.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIQay2ftvxH4JXwB815NeObwTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDQ4MzlaFw0yMjEwMjEyMDQ4MzhaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1MThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKu31Ys%2F%2BDxW7GFIU033kyQu8XHyD3%2BD%2BhbpQye2F2Nkjqa2x42Ls0D52MxlXxDJoOr9uwyUmyaB2Q2x1qATKOaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQMNWg%2FXxZpas%2Fj7eFhfYv3qNK%2F8DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1MThKMAoGCCqGSM49BAMCA0kAMEYCIQCe88puSJhFUesvA88jUrRxgzIVswb49Qmu%2FQ5J8AY3DAIhALL5vs9xtYGsmI9We34rbyVmEfibJFI6tR8UGlG%2FCWQf)

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


Generated: 12 Feb 24 17:02 UTC