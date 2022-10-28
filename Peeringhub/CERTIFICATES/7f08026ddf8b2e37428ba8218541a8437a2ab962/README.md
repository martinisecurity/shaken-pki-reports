# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate 7f08026ddf8b2e37428ba8218541a8437a2ab962
Tested At: 2022-10-28 19:21:56 +0000 UTC\
Initial Validity Period: 3650 day(s)\
Remaining Validity Period: 3523 day(s)\
Subject: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/240K/240K.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDDzCCArWgAwIBAgIRALZNFq96KG4nRDKyzVhPcaUwCgYIKoZIzj0EAwIwazELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMR8wHQYDVQQDDBZQZWVyaW5naHViIEluYyBSb290IENBMB4XDTIyMDYyMjIyNDUwMloXDTMyMDYxOTIyNDUwMlowfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXEEZfOD4d6FnCYEcvzXAMpFyyP8F1fGefqfxy%2BR22euER5DvlXYg0vRiyF1hDKC%2FhvM7eC74evVeiNgpWB3j%2Bo4IBJzCCASMwDgYDVR0PAQH%2FBAQDAgGGMA8GA1UdEwEB%2FwQFMAMBAf8wHQYDVR0OBBYEFK6hc1GIKVcRygyp9LEKbk64S00HMB8GA1UdIwQYMBaAFBzpokb%2F6xMa1JLngvRijv4Fon%2FPMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwCgYIKoZIzj0EAwIDSAAwRQIgGkVAngGoauAcC66weTZ39bchjkWjkJKdZifjcqqM%2Fy0CIQCmG8NEcrd6aC92TjHpN2%2Fd38qvM8uAbTECpsy%2Bsmqhcg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_ca_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ca_key_usage_crl_sign | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

Generated: 28/10/2022 at 19:22:10