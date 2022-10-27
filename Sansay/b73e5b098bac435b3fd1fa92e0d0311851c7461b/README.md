# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate b73e5b098bac435b3fd1fa92e0d0311851c7461b
Tested At: 2022-10-27 00:06:30 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN Connexum LLC 203K, OU=Connexum LLC, O=Connexum LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Connexum_LLC_203K

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1zCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTyowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTkyOFoXDTIzMTAxMTE3MTkyOFowczELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFTATBgNVBAoMDENvbm5leHVtIExMQzEVMBMGA1UECwwMQ29ubmV4dW0gTExDMSEwHwYDVQQDDBhTSEFLRU4gQ29ubmV4dW0gTExDIDIwM0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASSDurgWBjek5iHEZbxOI6guiPHbiIxidfoRgS0d0T7X7tUzFBFZbyPtClNVUjtuCxE8mfb%2BrsLeiogCXpnQTbDo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMDNLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUFA265FeWEfQo8h1%2FaXjH58B2kcAwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBf5guIcZw8ncpBfwSayJ1YDGi7tLHXZHrEYj3zZyWtpAiEAyj7%2Fvl7N1rzcbPBlJG%2BBzWVrkJkbe4xDt%2FwJPj2QfIo%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 203K' |

Generated: 27/10/2022 at 00:07:07