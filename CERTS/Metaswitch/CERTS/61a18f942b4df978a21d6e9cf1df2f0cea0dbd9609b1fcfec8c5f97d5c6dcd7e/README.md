# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Buckeye SHAKEN Cert 7608

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 129 day(s)\
Subject: CN=Buckeye SHAKEN Cert 7608, O=Buckeye Telesystem Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3c6ee1503f2b739a7b68a95572cba9a30bfad75f

[View certificate details](https://understandingwebpki.com/?cert=MIICWTCCAf6gAwIBAgIQKFML8OQtAH%2BU1ugMy%2BImGzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYyMTE1MDUzOFoXDTI0MDYyMDE1MDUzOFowUTELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkJ1Y2tleWUgVGVsZXN5c3RlbSBJbmMxITAfBgNVBAMMGEJ1Y2tleWUgU0hBS0VOIENlcnQgNzYwODBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJcZg3k8YdBoQhcnc0FlKNUuVMQ6B8xlJCR%2Fwh%2BRcGEClEx%2BADSYbZFbgMQbGhlfwoSRdPCv9xVkFRwzN3eCM3WjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc2MDgwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU4YkIR39UaK2zVovk0U1nN%2FFXlAkwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMKzPT8OPYcsVwzs4QQwO25tQF5Au3rAJ7673LbvduhGAiEA5N6jmQq%2FLVmEN4ijgI31iIh%2FbK68KedW3Gq9iszCtRE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC