# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate ea1f353b3f34c6823727490c5dd0c497bfdf07fd
Tested At: 2022-10-28 18:22:51 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN IPitomy 652J, OU=Operations, O=IPitomy, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/IPitomy_652J

View: [Click to view](https://understandingwebpki.com/?cert=MIICyTCCAm6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzcyNVoXDTIyMTEyNjAwMzcyNVowZDELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEDAOBgNVBAoMB0lQaXRvbXkxEzARBgNVBAsMCk9wZXJhdGlvbnMxHDAaBgNVBAMME1NIQUtFTiBJUGl0b215IDY1MkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATya7C0e2%2F7028Zf07DXg%2BtZHQ5IwPYAr92QpZf1GcdmUZcSKaB3Su%2BTsDXww8AHKPLLNUXB6u5V1709i9AgLo9o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NTJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU%2Fc6R%2B8KytDCVC%2BZXONwWCd6j3gIwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQClNb4cv%2F5%2BqX9n%2BIuzlyjN%2BSDJpCj7jHEtL1kWtbQiOwIhALCbAcgH4JF1nrhdNMJsSa4uHVgdJdAUYHY%2BwhOuaLvl)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 652J' |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 28/10/2022 at 18:22:55