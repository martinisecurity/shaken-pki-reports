# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 107K

Tested At: 28 Nov 23 19:53 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -430 day(s)\
Subject: CN=SHAKEN 107K, OU=SHAKEN, O=Atheral, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4ba6f296-2d7c-42c2-8fe4-12b74f5076df/c1634f554819e62264e1e04f4544a32f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIQQB8oAR1j9NYm68B50rR1pzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcxNzIwNDdaFw0yMjA5MjQxNzIwNDZaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdBdGhlcmFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxMDdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGOHJ110MB4MgrjD94Aq2yL%2BowsMi9dt5xMBE2Pc93H8JlYf1bmc%2FuAlN1Ug1vqiF7siu%2F%2Bg1Xhj6P593ArAkfaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSlq%2BYf%2B17q%2FXlgwYc4WHUI9uxKrTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMDdLMAoGCCqGSM49BAMCA0kAMEYCIQCbAqQA6L1p%2FMAlSqiiGctHEQTV2ST0JH6D50w7Sf3EwgIhAIlSiS%2BTyIaW7FsVOfyayTrbyt%2F7pIkPECKBmAuJfd30)

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


Generated: 28 Nov 23 20:21 UTC