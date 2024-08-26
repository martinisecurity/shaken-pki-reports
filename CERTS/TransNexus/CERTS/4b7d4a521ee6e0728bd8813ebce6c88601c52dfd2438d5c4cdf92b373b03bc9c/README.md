# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 819H

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -671 day(s)\
Subject: CN=SHAKEN 819H, OU=SHAKEN, O=BluIP Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3c4ce448-386b-4d47-a276-7fe32e380a83/ffef99cc09645bfeb327380d6b236a5f.pem

[View certificate details](https://x509.io/?cert=MIIC7DCCApKgAwIBAgIQXKeoeJUMnbmkBNdE5XmyaTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDE5NDFaFw0yMjEwMjQyMDE5NDBaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpCbHVJUCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4MTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE21uIAcwQUiwSApH5lLCbE5xO9DkDh2cU6goZUiUOHazzglbeE0dMzh4RKgdb1WNzK5QHBjWOXqRB5inHveAD66OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTQGDVPDmh31Y8qRU%2Bp0nrrpQDFFzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTlIMAoGCCqGSM49BAMCA0gAMEUCIDh1xQw2uNx5qPYpNc0QaJUKM%2FqpUQhNmerE5TZ9RKbvAiEA28%2FSJgasvQH4KFoABoduhHbfLsnHp6yhzPHnoxIxjPU%3D)

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


Generated: 26 Aug 24 18:03 UTC