# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 674J

Tested At: 27 Nov 23 23:11 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -393 day(s)\
Subject: CN=SHAKEN 674J, OU=SHAKEN, O=Panterra Networks Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0b98ecf4-0a43-4b89-8fc0-8e029d8258fa/c965b971d4b3e1371cd7155cfb0b07d5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQVo5RBQP6AI7e6JcVe7eeaDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDE3MjVaFw0yMjEwMzAyMDE3MjRaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZQYW50ZXJyYSBOZXR3b3JrcyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2Bdc5i5WnarW8mFITuHFZkFUFebdA38roNgXT2J8UmzDqiM00TNAsa6E15IL7TEKQMIPkccdUy7SfQBXtC5aQvKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQJt8CV36hdV2VVzYoNamtMJdtIATAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzRKMAoGCCqGSM49BAMCA0gAMEUCIE495wcqiSXdhHns9wUuDjaGRWtr1tp2w%2Fawca5Eut%2F8AiEA1ZlHzc62BaCcv0M6lEyc8%2FYkY8%2B3UsAB9mEGrW3obxQ%3D)

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


Generated: 27 Nov 23 23:28 UTC