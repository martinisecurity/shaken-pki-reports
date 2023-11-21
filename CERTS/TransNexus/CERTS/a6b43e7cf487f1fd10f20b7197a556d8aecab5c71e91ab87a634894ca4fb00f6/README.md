# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 031K

Tested At: 21 Nov 23 16:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -376 day(s)\
Subject: CN=SHAKEN 031K, OU=SHAKEN, O=TISD, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/031K/ba0288df-daa5-44b5-a1a3-3270140129ed.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5zCCAoygAwIBAgIQZ4%2FLfYIaElEsdRhdJkTawDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODM2NTlaFw0yMjExMDkxODM2NThaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRUSVNEMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwMzFLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEA2otyUR7bCPIGK2ozboGh9D700pABVHHHb%2Fa0bY%2FZ6Tpy4DGW6VTnjEvdRx2pIUHZQ2s524J17DDDsKmwKbk4KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQB74WG0xTnvIunT5aNICTQWcrMTzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwMzFLMAoGCCqGSM49BAMCA0kAMEYCIQDa%2BO8TeVTQsXv9diKJ4oconj25X35D4W68k5XcMB0dsgIhAP9%2FpQpVu01L3oUQj%2BI%2BXf%2BgPPaATbtCLGcuNOBu5XvQ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC