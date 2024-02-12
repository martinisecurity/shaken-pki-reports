# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 495J

Tested At: 12 Feb 24 16:29 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -484 day(s)\
Subject: CN=SHAKEN 495J, OU=SHAKEN, O=Stratus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/94c136bd-6438-40a6-b740-7765d5ead597/c554fd909a16573071e40fbac14f35bd.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQXknWFmw2dlRl8TCQV2b5lDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgyMDIwMTJaFw0yMjEwMTUyMDIwMTFaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdTdHJhdHVzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0OTVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKuD0pUcCH%2FUGw7DfG2VDJysdrNK%2BaDTJMwnEQUAPq%2FTsJgBfkICUXqdORRp4uBeaXZUKtMC7yIhV288cwvE7hqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBT%2FoUVKMhRX8R4gZDcvBRIqQ3CnHzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTVKMAoGCCqGSM49BAMCA0gAMEUCIHoRCsc37vizOdlA57QjFh%2B5IqMBrBE%2BWXmTexFalchJAiEAxvbZTIvWihFY9eFIbGtxNQGbH9plgu8%2FfM0CzguTWXs%3D)

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