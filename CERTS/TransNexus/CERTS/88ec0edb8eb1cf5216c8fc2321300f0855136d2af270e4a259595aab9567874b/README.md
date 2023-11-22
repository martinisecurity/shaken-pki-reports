# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 177K

Tested At: 22 Nov 23 03:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -412 day(s)\
Subject: CN=SHAKEN 177K, OU=SHAKEN, O=Cytracom\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/8d118994-4994-4735-ac71-42c0bbb7848f/6c0aaa065be50fd3054801aedd8e8406.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApWgAwIBAgIQeiP9j6peR3eq3qKtgTfUNTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjgxNzM2MTlaFw0yMjEwMDUxNzM2MThaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DeXRyYWNvbSwgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKuhToAgDgxg1sjlK6XmzDUsG62g3KMmAz%2Fk%2FXHO%2FvdqaBHd%2FTXM%2Fk2Hkud9zXOOBsvX904%2FsExMSewMUXdSGH6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQSI8fCcpfqljaSwwuQyvBDKnknUjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzdLMAoGCCqGSM49BAMCA0cAMEQCIFHCOqrRfoBWHqACjPbp6J1G1eISMmGPbU436ccjfYZSAiB9V%2Bir9xeBFmiLcWEupk1eVvQJPnJYPq1ArJ2fJfTZ7Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC