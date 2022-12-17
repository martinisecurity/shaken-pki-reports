# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ConvergeTel LLC 388K

Tested At: 17 Dec 22 12:15 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: 177 day(s)\
Subject: CN=SHAKEN ConvergeTel LLC 388K, OU=Voice, O=ConvergeTel LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ConvergeTel_LLC_388K

[View certificate details](https://understandingwebpki.com/?cert=MIIDgjCCAyigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWHEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIxMzE5MzgxNFoXDTIzMDYxMTE5MzgxNFowcDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRgwFgYDVQQKDA9Db252ZXJnZVRlbCBMTEMxDjAMBgNVBAsMBVZvaWNlMSQwIgYDVQQDDBtTSEFLRU4gQ29udmVyZ2VUZWwgTExDIDM4OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATssxwTUb6pO%2FIYAdGV1vS00MjDuKx6607CG2HdjJu0PO8gkCWZZjHVOKx9u1z%2F7%2B3e7VQjcPsz0lM1zM9Vtoaso4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDM4OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBREoflHIRQLSFn9f43xXgookT3WazCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCxDZ1qpbQ762qIso5sCXzE53BBixNZknyJcMINySKAMwIgONP%2FZpAV%2FqSm1EypacWsPBnPSesuvkvOiqkH0nZYaw0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 388K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 17 Dec 22 12:22 UTC