# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 469A

Tested At: 04 Oct 24 15:40 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -743 day(s)\
Subject: CN=SHAKEN 469A, OU=SHAKEN, O=T3 Communications Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88814589-ad88-4c53-b3c3-47f8334afb98/65b9c8389b135d2db8bd95e6d1e8c36e.pem

[View certificate details](https://x509.io/?cert=MIIC%2BDCCAp6gAwIBAgIQXDrNiTm47b17muX%2FqotOYTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE3MzBaFw0yMjA5MjEyMDE3MjlaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZUMyBDb21tdW5pY2F0aW9ucyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0NjlBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4oFNOkF1tQoP4zcU4NAZ5nD8feAHZIVdJ7%2F2KAjqkW3Nv7X3EXWqtJ59WP%2FjXajD2V8awcKuLrkZ4fqE1senbKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTiU2E2tasOIeagxI3F%2FHC%2BWnHwbTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjlBMAoGCCqGSM49BAMCA0gAMEUCIDU7CzioPgMqFR87bgBCgHdEVodAKtlGiLvQB8VEkqDRAiEA3EDI4vxSa5mol%2B9vckdubxfttCkHEnkjeAguynPeF90%3D)

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


Generated: 04 Oct 24 16:29 UTC