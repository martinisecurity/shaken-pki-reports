# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 159H

Tested At: 26 Aug 24 17:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -686 day(s)\
Subject: CN=SHAKEN 159H, OU=SHAKEN, O=Edge Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/159H/c52be509-f1a2-4b77-9300-df1f3ffb4245.pem

[View certificate details](https://x509.io/?cert=MIIC9DCCApugAwIBAgIQd2%2FyGedwhpR7UhFwYrABojAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDMwOTM0MDhaFw0yMjEwMTAwOTM0MDdaMFIxCzAJBgNVBAYTAlVTMRwwGgYDVQQKExNFZGdlIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE78Onlk2jyHL1aj%2FAKR6uGiBY16%2BjFzL%2FxfxEkevuw%2FY2Ya%2F76dWEoaWel%2FiqCXJIwcMXjdRKNVVmb7uEyPx6S6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQ1ANsNOV%2BgwhS5r14GRocHHUWV0TAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNTlIMAoGCCqGSM49BAMCA0cAMEQCIDRTHiXzXYc6L2m%2BaCMC2uRSupraUIPF0pNO96FREorWAiBmVl3YhmfxXDgnYbsNDj0nRUNwl%2BqEYOO39Yjg7XGaGw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC