# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 179K

Tested At: 22 Nov 23 03:18 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -396 day(s)\
Subject: CN=SHAKEN 179K, OU=SHAKEN, O=The Tech Consultants\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2d93d07b-20ff-4289-a40d-959976cc0595/acc7328684bd54f2b672b0988753cf96.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqGgAwIBAgIQU1peHd0XP8IWH%2Fpv%2F%2BNR5TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQxODQ0MjlaFw0yMjEwMjExODQ0MjhaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlUaGUgVGVjaCBDb25zdWx0YW50cywgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkLvTGPCszgVaZr2lAV7oberJROmJ3opD%2F%2FivBDjSWU70fARH43cAOoU1eSvqbm4viQ1f9z2oy4U4o2scMqGMrKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSFJjKT0ZMgCHqekmb39DL4ZwsHNzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzlLMAoGCCqGSM49BAMCA0gAMEUCIQDm%2FtUzJvq6HoSZluaOs0Vb8bIdPN3W6ujdv6BLQW%2FD4wIgZwMoKwggFESSB1M5MYxxdYHo0E3MYuhAaEHdPgxYLYM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC