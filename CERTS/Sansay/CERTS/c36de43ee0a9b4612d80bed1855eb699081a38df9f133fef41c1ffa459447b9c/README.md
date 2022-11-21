# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Magna5, LLC 3849

Tested At: 21 Nov 22 20:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN Magna5\\, LLC 3849, OU=Operations, O=Magna5\\, LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Magna5_3849.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDfDCCAyKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVZgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMTkwM1oXDTIyMTIxNTIxMTkwM1owajELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAtNYWduYTUsIExMQzETMBEGA1UECwwKT3BlcmF0aW9uczEgMB4GA1UEAwwXU0hBS0VOIE1hZ25hNSwgTExDIDM4NDkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR0phL5NKHG781QqEg9ZLVQsnMgBZhvqhWDIboajsBYFH7o4nQcN5Xi%2F7vDaerLQs5t%2Fs0SKYvQIVNuAYLaJELho4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDM4NDkwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTUNSS%2F%2BskE2xBGU9i%2Frkve0RREQDCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQC%2FGmfFE2zYZX86S0bWPeB5sFtU5gy4eJ6nhfDwW75GrQIgdhq5yrgTNOtw17bOF1pWBQVD0zi2W%2B9qFkD0zMmZ7io%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3849' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Nov 22 20:55 UTC