# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 15 Nov 23 16:14 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -371 day(s)\
Subject: CN=SHAKEN 672B, OU=SHAKEN, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/fd670ee0966d161a5cab78616bbb5670.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQZEdO%2F4FTaydORCKhiLlXOzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDI0MzhaFw0yMjExMDgyMDI0MzdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpDbGVhciBSYXRlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzJCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZgvQvGFMZb%2BiLUYLJYSGJOPhaT4KmZzljE63MWyfHE0nWj6tgVm001qJ3TkgWuV8E%2BIrcUQB%2FSDjKSU7uiMZkaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBR697Js2u%2Ba2vWEGGWTkIRKkGu5OTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJCMAoGCCqGSM49BAMCA0gAMEUCIBrNkRZaDbDcmHafzpRFgV5B6jI5i85yCZIldhYGPHSVAiEAoIJVtZa%2BfNnMHjhX4BBflq0jnsyXdeHTCgcOzDUb%2FSg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC