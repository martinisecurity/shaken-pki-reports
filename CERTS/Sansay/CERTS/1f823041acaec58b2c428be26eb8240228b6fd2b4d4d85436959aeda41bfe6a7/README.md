# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Connexum LLC 203K

Tested At: 31 Oct 22 20:25 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 345 day(s)\
Subject: CN=SHAKEN Connexum LLC 203K, OU=Connexum LLC, O=Connexum LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Connexum_LLC_203K

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1zCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTyowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTkyOFoXDTIzMTAxMTE3MTkyOFowczELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFTATBgNVBAoMDENvbm5leHVtIExMQzEVMBMGA1UECwwMQ29ubmV4dW0gTExDMSEwHwYDVQQDDBhTSEFLRU4gQ29ubmV4dW0gTExDIDIwM0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASSDurgWBjek5iHEZbxOI6guiPHbiIxidfoRgS0d0T7X7tUzFBFZbyPtClNVUjtuCxE8mfb%2BrsLeiogCXpnQTbDo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMDNLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUFA265FeWEfQo8h1%2FaXjH58B2kcAwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBf5guIcZw8ncpBfwSayJ1YDGi7tLHXZHrEYj3zZyWtpAiEAyj7%2Fvl7N1rzcbPBlJG%2BBzWVrkJkbe4xDt%2FwJPj2QfIo%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 203K' |


Generated: 31/10/2022 at 20:32:42