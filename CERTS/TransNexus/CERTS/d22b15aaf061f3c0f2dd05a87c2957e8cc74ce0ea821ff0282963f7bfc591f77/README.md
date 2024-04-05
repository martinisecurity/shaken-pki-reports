# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 297K

Tested At: 05 Apr 24 18:42 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -540 day(s)\
Subject: CN=SHAKEN 297K, OU=SHAKEN, O=Clarity Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/859299ca-4c86-4dfd-9785-6c758bee1b37/6b7e2daa26347b1369c76e30cb8994a2.pem

[View certificate details](https://x509.io/?cert=MIIC8DCCApWgAwIBAgIQR1DP3wkrHI%2BO7BY13SIFYjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUxOTU0MDJaFw0yMjEwMTIxOTU0MDFaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DbGFyaXR5IFZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAyOTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkzRX%2FO2xOI1WWfke%2BLBxEXty1Yv2dIFOziw4A6EO3mMsMwGueANC2v17LpXWBxmnxh94O3IPAPxaKG63979y4KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQoAy%2BLat%2B2oWESPk113vcHS2hIXDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQyOTdLMAoGCCqGSM49BAMCA0kAMEYCIQDKp6HUPWzCbyyr9dCoPldO4eQzsDylLwSpxXcaubfzPgIhAKgaPmwta77Nz%2FU7eM7PffLvR3YKi%2BMnSpz10Acd4zKf)

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


Generated: 05 Apr 24 19:04 UTC