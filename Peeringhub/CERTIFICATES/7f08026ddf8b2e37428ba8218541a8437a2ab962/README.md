# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate 7f08026ddf8b2e37428ba8218541a8437a2ab962
Tested At: 2022-10-27 18:24:45 +0000 UTC\
Initial Validity Period: 3650 day(s)\
Remaining Validity Period: 3524 day(s)\
Subject: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Issuer: CN=Peeringhub Inc Root CA, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/240K/240K.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDDzCCArWgAwIBAgIRALZNFq96KG4nRDKyzVhPcaUwCgYIKoZIzj0EAwIwazELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMR8wHQYDVQQDDBZQZWVyaW5naHViIEluYyBSb290IENBMB4XDTIyMDYyMjIyNDUwMloXDTMyMDYxOTIyNDUwMlowfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXEEZfOD4d6FnCYEcvzXAMpFyyP8F1fGefqfxy%2BR22euER5DvlXYg0vRiyF1hDKC%2FhvM7eC74evVeiNgpWB3j%2Bo4IBJzCCASMwDgYDVR0PAQH%2FBAQDAgGGMA8GA1UdEwEB%2FwQFMAMBAf8wHQYDVR0OBBYEFK6hc1GIKVcRygyp9LEKbk64S00HMB8GA1UdIwQYMBaAFBzpokb%2F6xMa1JLngvRijv4Fon%2FPMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTkoxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwCgYIKoZIzj0EAwIDSAAwRQIgGkVAngGoauAcC66weTZ39bchjkWjkJKdZifjcqqM%2Fy0CIQCmG8NEcrd6aC92TjHpN2%2Fd38qvM8uAbTECpsy%2Bsmqhcg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| e_sti_ca_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_ca_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- w_cp1_3_ca_subject_rdn_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 18:24:52