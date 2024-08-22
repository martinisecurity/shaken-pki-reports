# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 952J

Tested At: 22 Aug 24 15:28 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -648 day(s)\
Subject: CN=SHAKEN 952J, OU=SHAKEN, O=Syndeo LLC dba Broadvoice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/952J/b81c5f3e-5577-44d2-a22e-5bcfb144b13e.pem

[View certificate details](https://x509.io/?cert=MIIC%2FDCCAqGgAwIBAgIQQaPxdUWJnCfxUe8cP%2FV%2BizAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTMxODAyMzFaFw0yMjExMTIxODAyMzBaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlTeW5kZW8gTExDIGRiYSBCcm9hZHZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA5NTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEebmXtP%2FEfuPT%2FhMJDYWpnJXT%2FaoaL7ayWNh2B6WF2T3IQeVq4woR92DZyTghehNheWCq720r3nwSnSq97fjuFKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSN3k5gvLleS37Kd%2FO2fdv3B4vA0jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTJKMAoGCCqGSM49BAMCA0kAMEYCIQChl4Rqx%2F4SOKQT8IkMCwI7rNtBYz%2BDdhXmvDvnAVAl6QIhAIbIUkF%2FcEJWgJObrZuuxkGHhYAITlHsnO88nrv%2FQYRM)

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


Generated: 22 Aug 24 15:44 UTC