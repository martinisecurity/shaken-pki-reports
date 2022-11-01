# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Peeringhub Inc SHAKEN Intermediate CA 2

Tested At: 01 Nov 22 17:00 UTC\
Initial Validity Period: 3650 day(s)\
Remaining Validity Period: 3519 day(s)\
Subject: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIIDDzCCArWgAwIBAgIRALZNFq96KG4nRDKyzVhPcaUwCgYIKoZIzj0EAwIwazELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMR8wHQYDVQQDDBZQZWVyaW5naHViIEluYyBSb290IENBMB4XDTIyMDYyMjIyNDUwMloXDTMyMDYxOTIyNDUwMlowfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXEEZfOD4d6FnCYEcvzXAMpFyyP8F1fGefqfxy%2BR22euER5DvlXYg0vRiyF1hDKC%2FhvM7eC74evVeiNgpWB3j%2Bo4IBJzCCASMwDgYDVR0PAQH%2FBAQDAgGGMA8GA1UdEwEB%2FwQFMAMBAf8wHQYDVR0OBBYEFK6hc1GIKVcRygyp9LEKbk64S00HMB8GA1UdIwQYMBaAFBzpokb%2F6xMa1JLngvRijv4Fon%2FPMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwCgYIKoZIzj0EAwIDSAAwRQIgGkVAngGoauAcC66weTZ39bchjkWjkJKdZifjcqqM%2Fy0CIQCmG8NEcrd6aC92TjHpN2%2Fd38qvM8uAbTECpsy%2Bsmqhcg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [e_atis_ca_certificate_policies](../../ISSUES/e_atis_ca_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 17:00:23