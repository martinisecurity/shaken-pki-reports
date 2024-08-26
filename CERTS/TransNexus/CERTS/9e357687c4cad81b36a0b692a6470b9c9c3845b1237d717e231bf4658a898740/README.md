# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 060K

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -668 day(s)\
Subject: CN=SHAKEN 060K, OU=SHAKEN, O=Telware, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0c4e235f-3e3e-4dd3-bbfc-2f15badab180/655d0a57a1487790f33100d907332e2e.pem

[View certificate details](https://x509.io/?cert=MIIC6jCCAo%2BgAwIBAgIQWjUMdL2XVzTb1XvOcWL6sTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAxODEzMjBaFw0yMjEwMjcxODEzMTlaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdUZWx3YXJlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNjBLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGSBHYJASZuC%2BRNu7PY%2F%2FBcsep%2BEZFWkvC7Mg4ADzucrdSahc5W508EqDBo59xrzsjAgi1ZCmAUq4FR7SoY4pcqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBS0tZt89D%2FweyFpY%2FCG7CR%2FMzJgrTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNjBLMAoGCCqGSM49BAMCA0kAMEYCIQCShFtNZkHkvw52ZH4Hbv2uUXg4uT%2FccK4l9nF059ePGQIhANj4DK0vRO6p91lE%2F0fr0k3Jt4rdftwolOQFu11s9xrY)

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


Generated: 26 Aug 24 18:03 UTC