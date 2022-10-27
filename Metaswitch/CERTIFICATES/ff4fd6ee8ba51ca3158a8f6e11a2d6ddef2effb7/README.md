# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate ff4fd6ee8ba51ca3158a8f6e11a2d6ddef2effb7
Tested At: 2022-10-27 21:42:52 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 499 day(s)\
Subject: CN=Allstream SHAKEN cert 4130, O=Allstream Business US\\, LLC, L=Vancouver, ST=WA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://cert-stir-us.allstream.com/certs/allstreamcertchain.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICgDCCAiWgAwIBAgIQCqHY7R%2BQKo8XOGwx9tpRmTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxMDIwNDAwMVoXDTI0MDMwOTIwNDAwMVoweDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldBMRIwEAYDVQQHDAlWYW5jb3V2ZXIxIzAhBgNVBAoMGkFsbHN0cmVhbSBCdXNpbmVzcyBVUywgTExDMSMwIQYDVQQDDBpBbGxzdHJlYW0gU0hBS0VOIGNlcnQgNDEzMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPipWA%2BL%2B%2Bk2L8jvUDv%2BCE7eVsyWkAHLdyHVqge0DnMeEoBCSfxq5QmLeItuzSKNkM0vbdJvVF42MYkW%2By%2FYV0ajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDQxMzAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUs21bCRoUNlKajnDxQjAEaZlsWtwwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMklAtNDoG%2F3ZZ06g9r%2BRLSWfY4egtkXkvOf3L2KXlZuAiEA63u1XxXh2w8ktJ9yuPjAThXk2CbGEK1c2o%2B2Dvo7fZ4%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_sti_subject_cn
- w_cp_1_3_subject_email
- e_cp1_3_subject_sn
- e_sti_signature_algorithm
- w_cp1_3_subject_rdn_unknown
- e_sti_serial_number
- e_sti_extension_unknown
- e_cp1_3_ambiguous_identifier

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 21:42:52