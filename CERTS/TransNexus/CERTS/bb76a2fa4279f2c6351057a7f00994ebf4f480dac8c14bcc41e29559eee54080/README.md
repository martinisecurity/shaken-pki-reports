# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 495J

Tested At: 28 Nov 23 10:23 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -429 day(s)\
Subject: CN=SHAKEN 495J, OU=SHAKEN, O=Stratus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/94c136bd-6438-40a6-b740-7765d5ead597/8b3ee25d58998b94749600fcd5ec8b75.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQYu%2B%2BAXM8sZM9yhczyGJqdzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDE3NTBaFw0yMjA5MjQyMDE3NDlaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdTdHJhdHVzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0OTVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEw6WNBNlcWFKKz4AqnZ2usZtgBgCL8wZKeGYoItxgde1D116drYbd%2FZukqzX1JsK%2FOmA%2FvkGie5b31nTMbmkeCKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQPhgmSedjwgkJg%2FdRJy3nvY9wd5TAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTVKMAoGCCqGSM49BAMCA0gAMEUCIEPGKFvy7g2D1YHh6DvWv%2Fk02QJVNhJDRVR%2FNBW54MBzAiEAoorzXbkP8O4JfHqYjgNmUGgiYsmBJcE6MzaUmhZt2VI%3D)

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


Generated: 28 Nov 23 10:53 UTC