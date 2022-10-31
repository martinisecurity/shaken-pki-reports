# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 952J

Tested At: 31 Oct 22 20:39 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN 952J, OU=SHAKEN, O=Syndeo LLC dba Broadvoice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/952J/933ae4bb-a38f-410b-9c14-38fd2f4c7893.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BjCCAqGgAwIBAgIQWZrM%2BOzPbr%2FFIwTy2vRYEjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzExODA0MDlaFw0yMjExMzAxODA0MDhaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlTeW5kZW8gTExDIGRiYSBCcm9hZHZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA5NTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEccbtKocaRvW%2FfsY%2BINvKtGufeTV81GyTL4siYQj4CWdx%2Bd8kS1dI6LvHstTl9eFcQ1mzBnyyji2WsfDZmy4A9qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTUiQ8Q5cgjKm1UyBTVlDNeq%2FuKzzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTJKMAoGCCqGSM49BAMCA0cAMEQCIB7EenveO6b8%2Ba7pqixfTZzRN8MdMB2irQdfRhmAGizoAiBpQe47S55luZpObn%2FoRi%2FxYAGIYO3HLKifvTNuLVzFvQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 20:47:45