# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Terra Nova Telecom 382G

Tested At: 26 Aug 24 18:23 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -199 day(s)\
Subject: CN=SHAKEN Terra Nova Telecom 382G, OU=NOC, O=Terra Nova Telecom, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/382G_TERRANOVA_STIR_SHAKEN.crt

[View certificate details](https://x509.io/?cert=MIIDhDCCAyugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjVYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEwOTIxMDE1MFoXDTI0MDIwODIxMDE1MFowczELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGzAZBgNVBAoMElRlcnJhIE5vdmEgVGVsZWNvbTEMMAoGA1UECwwDTk9DMScwJQYDVQQDDB5TSEFLRU4gVGVycmEgTm92YSBUZWxlY29tIDM4MkcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQUZcBpTXzq%2F4l2CuZ90OBlLvIaqacu478cFveXc0%2Bt6sPgddYepZCyT9JcZG0NYLjUHUjCQ9WZslg8AbwbS0Qeo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDM4MkcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQZ3CGz9usyXL1owkd%2FQ8Oe1gPMCjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBxsQiCBIwj9R1jZF9AA%2BRCri2N%2BVwD9nEoK7BlIUWMVAiAZ6w5MDkGMkas0Ar3I2L14hmC%2FUj7sybUOCplZ03dQtQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 382G', but common name is 'SHAKEN Terra Nova Telecom 382G' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 26 Aug 24 18:49 UTC