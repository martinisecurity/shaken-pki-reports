# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Magna5, LLC 8249

Tested At: 23 Nov 22 18:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 23 day(s)\
Subject: CN=SHAKEN Magna5\\, LLC 8249, OU=Operations, O=Magna5\\, LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Magna5_8249.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDezCCAyKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVZkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMjA0NloXDTIyMTIxNTIxMjA0NlowajELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAtNYWduYTUsIExMQzETMBEGA1UECwwKT3BlcmF0aW9uczEgMB4GA1UEAwwXU0hBS0VOIE1hZ25hNSwgTExDIDgyNDkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQkoTMj2M%2B8DOeKQeTjDAPSm3hP13VARS11vJZSHL7HbZJu1qSeASLGRzlUU1%2B9sXc%2FtyXhdlpeMry%2F%2FIVp%2F0DFo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDgyNDkwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRyJDOKmOPqe%2FwWtpNxFSCBZfYfBTCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIAkELwAk5048T1rRNbMr8sC7uFVAsil%2B7IWgVfj11HPHAiAM1abdFzmBudatEw%2F6I3VSAT0z6u5q%2BytyaNToKx7yqQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 8249' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 23 Nov 22 18:09 UTC