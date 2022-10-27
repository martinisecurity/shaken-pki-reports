# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate bf14dbda90b2ffc796c34770c105c1ca625de56b
Tested At: 2022-10-27 21:42:42 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SIP.US LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/2185f3c0-09fe-4171-ac47-5f28eb0ef48e/a2dd72ab2b20edd2643533d9a1b707c7.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQWmwEBkLRGm%2BrXcGB7ExLCDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDE4MzdaFw0yMjEwMzAyMDE4MzZaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpTSVAuVVMgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2BJU7jetTx4wQU4MTCn9mxOaRm4HW%2BCH8j3bRQdrIPD%2FNgvxQSCaWooY%2FPJl1WNbwKLipD%2FqisaipUxkDcApnkaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRdTw25lIKnO9fMMqXvfQjlyeytDDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzhKMAoGCCqGSM49BAMCA0kAMEYCIQCZuMtseBQmSh142m%2BpuyAX%2FP7ybWuXMCIuW7sZPOyr5QIhAOie3jb5I1DnkTFL9EDWh4qnhKNdH%2B2NBiVcf9nbVNry)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 21:42:52