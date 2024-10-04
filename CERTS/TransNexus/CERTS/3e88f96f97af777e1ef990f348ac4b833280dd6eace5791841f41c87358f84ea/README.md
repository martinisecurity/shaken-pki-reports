# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 179K

Tested At: 04 Oct 24 15:34 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -716 day(s)\
Subject: CN=SHAKEN 179K, OU=SHAKEN, O=The Tech Consultants\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2d93d07b-20ff-4289-a40d-959976cc0595/0630c442a9ee31c6271432091950d170.pem

[View certificate details](https://x509.io/?cert=MIIC%2BjCCAqGgAwIBAgIQREX63H72pQDeqU1mvfgN6DAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExODQ0MDhaFw0yMjEwMTgxODQ0MDdaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlUaGUgVGVjaCBDb25zdWx0YW50cywgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1eR8WQlR5YWTYZMgsjRNHFlmVOdNUV4dg2l%2Bk5zdChi%2FHNggNh5kwh9EHY6GfAFqxe%2FTBqg4yeXayXH16c%2F56qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTNFznRI%2FUbKBrFGfW2gAdGUZhU5DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzlLMAoGCCqGSM49BAMCA0cAMEQCIHClM%2Bat8rnx4bNp1oT1wQB11pBYZCX7aAk8NusNCiMGAiASo0pkLElAuJFSDe4gTRr3nQRwQCv15jJ6e8Po%2B%2F%2F8FQ%3D%3D)

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