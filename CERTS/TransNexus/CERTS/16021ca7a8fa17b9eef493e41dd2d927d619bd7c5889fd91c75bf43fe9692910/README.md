# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 159H

Tested At: 02 Dec 22 07:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -47 day(s)\
Subject: CN=SHAKEN 159H, OU=SHAKEN, O=Edge Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/159H/d63a047f-9c5a-44e6-b60a-35f1c738015e.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApugAwIBAgIQbjn1q%2B9KkxVPQGfR40qPADAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgxNTM1MDdaFw0yMjEwMTUxNTM1MDZaMFIxCzAJBgNVBAYTAlVTMRwwGgYDVQQKExNFZGdlIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEgJ%2FCxUnXulydg7IL4bmPnz9hz5ECaVrL%2Bw%2BTxuWP%2FJVsY8Ep%2BPIBU5l7duv9qeBYTY6jP0BmQLEpQv9ejZ7WLaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRUuootYlIKeUcpNW9N2%2B9adlqt8DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNTlIMAoGCCqGSM49BAMCA0kAMEYCIQDQT6fQ%2B81vM8BABRuU0U0fYl417b%2FCs%2Fl3g7mANuo6rQIhAKWOBkU80m%2BrLqGJfJSab5S5mSe%2FG0wb%2FwG6W4JOEb46)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 02 Dec 22 07:46 UTC