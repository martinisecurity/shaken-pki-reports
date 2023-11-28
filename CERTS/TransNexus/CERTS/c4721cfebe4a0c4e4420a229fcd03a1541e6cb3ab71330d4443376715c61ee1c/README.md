# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 28 Nov 23 19:54 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -406 day(s)\
Subject: CN=SHAKEN 849J, OU=SHAKEN, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/bd852bf061c1db8dbe457102f941cfe4.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApKgAwIBAgIQa%2Bl1rlXfDqEPIkU0UsRJ6jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIwMjVaFw0yMjEwMTgyMDIwMjRaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpGdXNlLkNsb3VkMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NDlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqsr0gYGeYaRSFOxfG05497l3Y2msC9sGC0tD996exj6H6wz4Cr3eixJcHu6ed8H5Z6%2Fmv%2BbSuC0%2BtirXZRaqXaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQqdeVbORHLAJkkKtWXoq0JRH5j%2BzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlKMAoGCCqGSM49BAMCA0cAMEQCIDv9q0tgcx%2Blct2gzqTC6uQutDj1yocrfQ9GgXlcHN73AiBHeIBFUi5wWDGj%2BTHOYXlUIjdMR4PXd7Oz1Z9CRvaQWw%3D%3D)

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


Generated: 28 Nov 23 20:21 UTC