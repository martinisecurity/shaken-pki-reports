# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 736J

Tested At: 27 Nov 23 22:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -151 day(s)\
Subject: CN=SHAKEN 736J, OU=SHAKEN, O=Masergy Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/736J/27cf1c16-f0f0-41fa-a0c1-6c167396fe34.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQWRQ2JbMmoCW2ZxrIW1Oa2zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjkyMDI0MzhaFw0yMzA2MjkyMDI0MzdaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZNYXNlcmd5IENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYKlJkB%2BmZz1ZOnkBLs9aVrtmGG7BJm0QKG6NbbjOQd0RzU7PF6UGW7T5oSbyFNHoM%2B0n%2Bc8GFCSUEiyMbFGjM6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQcKm6v0ilnb5qwHMq%2BkpidObWlPjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzZKMAoGCCqGSM49BAMCA0gAMEUCIQCUvrdfgejPV9Kr6huNG7OByuX8CO%2FYVYmEhrHjqE84cgIgOjoQsGgc7%2Bqfnqlf2C5cr6f145EUxDzoOTyujQ9%2BX%2F0%3D)

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