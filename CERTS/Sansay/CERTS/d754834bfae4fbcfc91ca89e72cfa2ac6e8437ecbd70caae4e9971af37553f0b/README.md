# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IDT America, Corp 363A

Tested At: 21 Nov 22 20:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN IDT America\\, Corp 363A, OU=NOC, O=IDT America\\, Corp, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/IDT_363A

[View certificate details](https://understandingwebpki.com/?cert=MIIC2DCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVYgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMDYyMFoXDTIyMTIxNTIxMDYyMFowdDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGjAYBgNVBAoMEUlEVCBBbWVyaWNhLCBDb3JwMQwwCgYDVQQLDANOT0MxJjAkBgNVBAMMHVNIQUtFTiBJRFQgQW1lcmljYSwgQ29ycCAzNjNBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9AzzI%2BlYrGbQEn6MZEj%2F%2Bf33BQM6UzSk%2BTC6JOpbLxMvM3FJrWpHuklhAPC1O7O6vsNvtfgcYHnJnHIiJdlYvqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzYzQTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFG3FRvnniG7YKOzgSDNbU5FwXNFXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBF0K2bnH2IwGLRyUtnRcZKS85HbGWuKD7LJ%2Fj8EbIjjAIhAOSPVrZZWfI72%2BaImr7Vjzek3HBICu20wHZs%2BjlAH951)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 363A' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Nov 22 20:55 UTC