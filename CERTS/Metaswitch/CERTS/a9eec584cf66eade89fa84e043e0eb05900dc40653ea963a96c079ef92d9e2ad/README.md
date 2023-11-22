# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Allstream SHAKEN cert 4130

Tested At: 22 Nov 23 03:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 109 day(s)\
Subject: CN=Allstream SHAKEN cert 4130, O=Allstream Business US\\, LLC, L=Vancouver, ST=WA, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cert-stir-us.allstream.com/certs/allstreamcertchain.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICgDCCAiWgAwIBAgIQCqHY7R%2BQKo8XOGwx9tpRmTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxMDIwNDAwMVoXDTI0MDMwOTIwNDAwMVoweDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldBMRIwEAYDVQQHDAlWYW5jb3V2ZXIxIzAhBgNVBAoMGkFsbHN0cmVhbSBCdXNpbmVzcyBVUywgTExDMSMwIQYDVQQDDBpBbGxzdHJlYW0gU0hBS0VOIGNlcnQgNDEzMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPipWA%2BL%2B%2Bk2L8jvUDv%2BCE7eVsyWkAHLdyHVqge0DnMeEoBCSfxq5QmLeItuzSKNkM0vbdJvVF42MYkW%2By%2FYV0ajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDQxMzAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUs21bCRoUNlKajnDxQjAEaZlsWtwwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMklAtNDoG%2F3ZZ06g9r%2BRLSWfY4egtkXkvOf3L2KXlZuAiEA63u1XxXh2w8ktJ9yuPjAThXk2CbGEK1c2o%2B2Dvo7fZ4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC