# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 186K

Tested At: 28 Nov 23 10:23 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -405 day(s)\
Subject: CN=SHAKEN 186K, OU=SHAKEN, O=Go2Uno LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/92df566f-5229-4e9a-b7c7-ec203fc4196d/f6fd06cff383c1babdd4f9e1eca1b778.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQRKdQ6M44R0KbbkzpNp24iTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExNDI2MzdaFw0yMjEwMTgxNDI2MzZaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpHbzJVbm8gTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxODZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6EPczLXwmeIZ%2BYk1tNKXzYQszXtwlSONjPdhgb6l0mRYGJOHMUSg6VYaPkQq4YqpzqxQJdEckNr3Y68c1VK6RaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSW7TVk07q6lr0DePybMKpKp662JDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxODZLMAoGCCqGSM49BAMCA0kAMEYCIQDvtAfaGgfDB1Dnwaxu3mqsnte6i%2B%2FoTu9WiEr%2Bci%2Bo8gIhAJC33OMOYlfnXHWy4NhwJ1%2BMOvsSKC%2F4ZdlctJjyytx0)

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


Generated: 28 Nov 23 10:53 UTC