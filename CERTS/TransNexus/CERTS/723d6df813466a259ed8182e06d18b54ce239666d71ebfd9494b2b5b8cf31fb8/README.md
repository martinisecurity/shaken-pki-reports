# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 790J

Tested At: 12 Feb 24 16:29 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -508 day(s)\
Subject: CN=SHAKEN 790J, OU=SHAKEN, O=Viirtue, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a429df3a-62d2-4851-90c8-fc446859fb08/df09a6def78f6ee9cf90e390af802924.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIQY17HZXrQ3G2CrLNWJlCSBDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE3MzdaFw0yMjA5MjEyMDE3MzZaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdWaWlydHVlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3OTBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4GLYX4%2BCT1xbOwI5Z3cN%2B6Qfv9gLY95VecXqjUrmwbcefLl6t9e%2BgsjgEKx%2FW79J%2BaphFQAZntQaPiypi5LDqqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBT0m9CbD%2FOOXXAJZOcbu6TPvg%2BUAjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3OTBKMAoGCCqGSM49BAMCA0kAMEYCIQC64ipRfW3KP9EHTQUgaB68GJsrXKVDcAeVQqelKRiQOgIhAPy3hKNOD0%2FOP8kmg9pH09z6qUcgfnMiFYWnwfkEva12)

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


Generated: 12 Feb 24 17:02 UTC