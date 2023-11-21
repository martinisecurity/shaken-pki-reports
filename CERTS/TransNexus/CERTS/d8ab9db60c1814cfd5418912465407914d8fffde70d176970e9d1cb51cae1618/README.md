# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 21 Nov 23 17:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -425 day(s)\
Subject: CN=SHAKEN 849J, OU=SHAKEN, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/5d3642ee14c8c74eed126e51171527ee.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApKgAwIBAgIQfn5ErpQxztWS%2FxDODYiCBTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTQyMDE3MzBaFw0yMjA5MjEyMDE3MjlaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpGdXNlLkNsb3VkMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NDlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMznWXLQk6Zi4nlyn8LIzYX3WYAupAIUlbn8mYDDdN2HgOQ5dp%2FkifkRFCzWHXdkQY%2FidZnbk7w1fNSed2zfIA6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQhjTQwXajl6rQgqEIDYdvccGbKgTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlKMAoGCCqGSM49BAMCA0cAMEQCIAGjFwSj9wtYm7KVoCZLfEFnjdnR22OQ2I%2FGSxDFASK3AiB1pvEUVpZQgOuDTw%2Fm%2BKDVgrM%2BoM1PjY9W9PWbpMRVqw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC