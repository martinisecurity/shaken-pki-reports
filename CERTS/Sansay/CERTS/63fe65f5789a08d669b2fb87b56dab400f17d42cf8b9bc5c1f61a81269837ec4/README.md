# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IDT America, Corp 363A

Tested At: 31 Oct 22 20:25 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 26 day(s)\
Subject: CN=SHAKEN IDT America\\, Corp 363A, OU=NOC, O=IDT America\\, Corp, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/IDT_363A

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2TCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU38wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzg1NFoXDTIyMTEyNjAwMzg1NFowdDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGjAYBgNVBAoMEUlEVCBBbWVyaWNhLCBDb3JwMQwwCgYDVQQLDANOT0MxJjAkBgNVBAMMHVNIQUtFTiBJRFQgQW1lcmljYSwgQ29ycCAzNjNBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9AzzI%2BlYrGbQEn6MZEj%2F%2Bf33BQM6UzSk%2BTC6JOpbLxMvM3FJrWpHuklhAPC1O7O6vsNvtfgcYHnJnHIiJdlYvqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzYzQTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFG3FRvnniG7YKOzgSDNbU5FwXNFXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAqUHVIbZajAzWI4NwYobmX1ljL2YdWrzbhyj7RmlJvwgCIQCYKj3mwK9wgRo8LuOGPASNaFkG964usnN6lWiiDYQLwA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 363A' |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 20:32:42