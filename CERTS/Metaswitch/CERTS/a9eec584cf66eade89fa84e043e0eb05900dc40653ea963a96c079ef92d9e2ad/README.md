# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Allstream SHAKEN cert 4130

Tested At: 31 Oct 22 16:39 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 496 day(s)\
Subject: CN=Allstream SHAKEN cert 4130, O=Allstream Business US\\, LLC, L=Vancouver, ST=WA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cert-stir-us.allstream.com/certs/allstreamcertchain.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICgDCCAiWgAwIBAgIQCqHY7R%2BQKo8XOGwx9tpRmTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxMDIwNDAwMVoXDTI0MDMwOTIwNDAwMVoweDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldBMRIwEAYDVQQHDAlWYW5jb3V2ZXIxIzAhBgNVBAoMGkFsbHN0cmVhbSBCdXNpbmVzcyBVUywgTExDMSMwIQYDVQQDDBpBbGxzdHJlYW0gU0hBS0VOIGNlcnQgNDEzMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPipWA%2BL%2B%2Bk2L8jvUDv%2BCE7eVsyWkAHLdyHVqge0DnMeEoBCSfxq5QmLeItuzSKNkM0vbdJvVF42MYkW%2By%2FYV0ajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDQxMzAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUs21bCRoUNlKajnDxQjAEaZlsWtwwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMklAtNDoG%2F3ZZ06g9r%2BRLSWfY4egtkXkvOf3L2KXlZuAiEA63u1XxXh2w8ktJ9yuPjAThXk2CbGEK1c2o%2B2Dvo7fZ4%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22