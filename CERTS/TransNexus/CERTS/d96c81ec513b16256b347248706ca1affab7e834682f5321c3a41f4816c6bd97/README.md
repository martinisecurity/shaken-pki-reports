# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 159H

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -89 day(s)\
Subject: CN=SHAKEN 159H, OU=SHAKEN, O=Edge Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/159H/c52be509-f1a2-4b77-9300-df1f3ffb4245.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9DCCApugAwIBAgIQd2%2FyGedwhpR7UhFwYrABojAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDMwOTM0MDhaFw0yMjEwMTAwOTM0MDdaMFIxCzAJBgNVBAYTAlVTMRwwGgYDVQQKExNFZGdlIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE78Onlk2jyHL1aj%2FAKR6uGiBY16%2BjFzL%2FxfxEkevuw%2FY2Ya%2F76dWEoaWel%2FiqCXJIwcMXjdRKNVVmb7uEyPx6S6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQ1ANsNOV%2BgwhS5r14GRocHHUWV0TAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNTlIMAoGCCqGSM49BAMCA0cAMEQCIDRTHiXzXYc6L2m%2BaCMC2uRSupraUIPF0pNO96FREorWAiBmVl3YhmfxXDgnYbsNDj0nRUNwl%2BqEYOO39Yjg7XGaGw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 07 Jan 23 19:18 UTC