# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate b5f5297b6f49e51369eac1089bd10d7ae0f8aa3b
Tested At: 2022-10-28 18:15:35 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 348 day(s)\
Subject: CN=SHAKEN TeleVoIPs 138K, OU=Corp, O=TeleVoIPs, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/TeleVoIPs_138K

View: [Click to view](https://understandingwebpki.com/?cert=MIICxjCCAmygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTwgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE2NDgzM1oXDTIzMTAxMTE2NDgzM1owYjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCVRlbGVWb0lQczENMAsGA1UECwwEQ29ycDEeMBwGA1UEAwwVU0hBS0VOIFRlbGVWb0lQcyAxMzhLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEi%2F5ouSX5xZkj3XnR9lxs1TdEZJU5AhaR6PgjPYxACDbSP%2B3w41mwbosq%2BsQBFuMHHOIsTRv0zRZHgstqBc8MUaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMTM4SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHq2l%2BVzU6t%2BSgU8DjMx3a0hLK9xMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBUcFKDXxrhE1BOH3B26kJhGdLgPxqAMt3mP%2BCPe8XLLgIhAMdUbGEaGSvNFQTglGipJ97OAcV3QzVo3OaZJraW449n)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 138K' |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 18:15:47