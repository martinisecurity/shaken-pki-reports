# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 735J

Tested At: 21 Nov 23 01:23 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -407 day(s)\
Subject: CN=SHAKEN 735J, OU=SHAKEN, O=AVOXI Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/11124177-79f3-48b2-867a-386d4dc61c99/60797bd734fa6f7e6aa28e23adf7e4d0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQSayep0mEyMUz%2BI0r3ezfhDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIyMDE2NDZaFw0yMjEwMDkyMDE2NDVaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpBVk9YSSBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEw18ZpsBoNr482NR6nvq9eHlnGyW1muw2Gs86yMtv6n00xph0%2B4B%2F8KbNn8BsbgxyZ7sjA60BCFHNL%2B9Fm7GfG6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQfMNSw55SnoleljyHkZhJFV8y9zzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzVKMAoGCCqGSM49BAMCA0gAMEUCIFATjZMBPXVp7SIS8pAMX9rY%2B8gJtTKP7Ue8pfJNW7a5AiEAr3WnTcXE0xA8J94FjS9zT6kYYcBu%2BksPs2%2FHv33QuCo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC