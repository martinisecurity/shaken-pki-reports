# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0127

Tested At: 04 Oct 24 15:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -668 day(s)\
Subject: CN=SHAKEN 0127, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/649d9560-7197-4132-94b4-40da224ec372/f3f01a543f6e4d1c3b553723ac5c4d4a.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIQX332I8%2BAnSgBmU2TcYCdjjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMTI4MjAyNDUzWhcNMjIxMjA1MjAyNDUyWjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiAwMTI3MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMRJJffyGIM4sXcsULsJKPtuU5XbAaSLR2WuC0ALkNEK6Yplr16DIt9eVfF75gnk1zuB5nthY2hC6hOfGB%2FYED6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSfC9B%2FiizTbSEui8k9rGKGXwU83jAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwMTI3MAoGCCqGSM49BAMCA0gAMEUCIDmxcCIcrsk1YEmPBszc6BPBMk%2FeFXxlWF2Mfi9LzneaAiEAhmVhg06fstOt0xEvHruYheo389XCsfzzNIpzQ4LYDzs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC