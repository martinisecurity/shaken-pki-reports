# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 469A

Tested At: 31 Oct 22 16:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 469A, OU=SHAKEN, O=T3 Communications Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88814589-ad88-4c53-b3c3-47f8334afb98/61db9dd902bd71a164aa5c3770731009.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9zCCAp6gAwIBAgIQWgYGT22D4vySy1pybDUadzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIyNDRaFw0yMjExMDIyMDIyNDNaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZUMyBDb21tdW5pY2F0aW9ucyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0NjlBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE58cq1EJ2rnLAUeDl98Ohg63RRj6R3FkZjUW8RT%2B4Fnat3Hs%2FD92joH%2FFfnyyjiHlDshuCR49dJmOxVMuXKh8R6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQlEjZGqYMicl2coANrsPDL7gU9VzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjlBMAoGCCqGSM49BAMCA0cAMEQCH3pNSy6gSxpjft7GOBn%2BTGqnZIj5PtI%2BMP9OFVUt1fMCIQDoEI3Q8oKm7Qr9CutmwsPfGeeSxTA241KZQUddHAh8jg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 16:43:22