# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate eb9f3119926d3cc1bcc5ce0ef37ab3adbad9da16
Tested At: 2022-10-26 23:14:36 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN IDT America\\, Corp 363A, OU=NOC, O=IDT America\\, Corp, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/IDT_363A

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1zCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTygwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTg1NVoXDTIyMTExMDE3MTg1NVowdDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGjAYBgNVBAoMEUlEVCBBbWVyaWNhLCBDb3JwMQwwCgYDVQQLDANOT0MxJjAkBgNVBAMMHVNIQUtFTiBJRFQgQW1lcmljYSwgQ29ycCAzNjNBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9AzzI%2BlYrGbQEn6MZEj%2F%2Bf33BQM6UzSk%2BTC6JOpbLxMvM3FJrWpHuklhAPC1O7O6vsNvtfgcYHnJnHIiJdlYvqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzYzQTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFG3FRvnniG7YKOzgSDNbU5FwXNFXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBXDraMJeJTq1FNWCr%2BmYRhGM%2BqoZTE%2BWyqzVjPf2r3wQIgDqLxzafWky%2BqHn3FVLeMQPBQnD9yHd2cUKA%2B9tIrbgs%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 363A' |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 26/10/2022 at 23:14:41