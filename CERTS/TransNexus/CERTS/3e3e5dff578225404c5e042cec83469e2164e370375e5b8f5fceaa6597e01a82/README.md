# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0753

Tested At: 04 Oct 24 15:36 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -611 day(s)\
Subject: CN=SHAKEN 0753, O=Endeavor Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/426afef5-4cac-437d-817b-671ec10e028a/206ada20b0a94cbdf3912a535104aca1.pem

[View certificate details](https://x509.io/?cert=MIIC1zCCAn2gAwIBAgIQadLTHsc1nXF2ePdeWRmS%2BTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTI0MjAzMDE1WhcNMjMwMTMxMjAzMDE0WjBFMQswCQYDVQQGEwJVUzEgMB4GA1UEChMXRW5kZWF2b3IgQ29tbXVuaWNhdGlvbnMxFDASBgNVBAMTC1NIQUtFTiAwNzUzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsnlyS7FGCAfrga%2FdmyHC%2FYqEgvB3eZZoBcZX%2BKVCG4Dz2YMTi8ynlFeDL7U1ZwgY%2B9h3iQ5YRsOK2Jd8DC7hL6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRvFC4wwRlh1QW0WeHC1geOCvMSrTAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNzUzMAoGCCqGSM49BAMCA0gAMEUCIFesEZhgHQmGGH04XcugSV4bndPHXFwx2g8914y8Uh%2BLAiEAnQC%2FKN69NnyB1VUpJ%2BurkkBSk2E0fv5x0nM%2FgdLyTKY%3D)

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