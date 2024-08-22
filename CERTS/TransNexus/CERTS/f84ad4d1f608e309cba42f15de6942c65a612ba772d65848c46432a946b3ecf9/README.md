# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 819H

Tested At: 22 Aug 24 15:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -640 day(s)\
Subject: CN=SHAKEN 819H, OU=SHAKEN, O=BluIP Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3c4ce448-386b-4d47-a276-7fe32e380a83/918c386fd44ecab91d5ded655b2c20b0.pem

[View certificate details](https://x509.io/?cert=MIIC6zCCApKgAwIBAgIQRhqabdAg5YjMxtbfBrMjUzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMyMDIxNDJaFw0yMjExMjAyMDIxNDFaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpCbHVJUCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4MTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYbN17sgMj1QnaeG0vfOcm4y3RAAWzccUNR%2BPsbslZfOijdWSfCmYU0b75cCNtkKONqNo83qkDrPcEXWnBU3m9KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSB%2BY0ESE8S%2B%2B77sXOB9RfLjhb2AzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTlIMAoGCCqGSM49BAMCA0cAMEQCIG1ZgRwAw5MFZ3NXw6NPsNE1iHzpLS4LKwPhT8wPLnrOAiBCxKbkq5lwn0YOl9wtyW8hhGY5CsmtIc8N4orzyQfy%2BQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC