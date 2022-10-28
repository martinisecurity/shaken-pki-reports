# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 824b3f223bd6cd5ec9f7748528dcc903eb8cd2db
Tested At: 2022-10-28 10:32:38 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 29 day(s)\
Subject: CN=SHAKEN Apeiron Systems 012J, OU=NOC, O=Apeiron Systems, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Apeiron_012J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1DCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzYwNFoXDTIyMTEyNjAwMzYwNFowcDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGDAWBgNVBAoMD0FwZWlyb24gU3lzdGVtczEMMAoGA1UECwwDTk9DMSQwIgYDVQQDDBtTSEFLRU4gQXBlaXJvbiBTeXN0ZW1zIDAxMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvM7voHFBJOHVADYo1gKHK2wU0lgJXrz%2FJPCTX0AKFzXtzFfZF0xEGkCL70IlycCO4vTtcloIGOIJTSN8ZSPxko4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMTJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUfHBaMnt53HaA60SExkm4zJE0Wl4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCJ48s6pkt1jyA2Ry28Me2szNB%2F3sJbsQR%2Fv763wYGqYgIgHL33tGiIHAo1GQeKkVN1INHUpRXZkm11J30UvmjdqFw%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 012J' |

Generated: 28/10/2022 at 10:33:25