# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Buckeye SHAKEN Cert 7608

Tested At: 06 Jul 23 13:53 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 351 day(s)\
Subject: CN=Buckeye SHAKEN Cert 7608, O=Buckeye Telesystem Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/ccf242c7f37d7d7d0af1091ba5d7894e312ccedd

[View certificate details](https://understandingwebpki.com/?cert=MIICWTCCAf6gAwIBAgIQKFML8OQtAH%2BU1ugMy%2BImGzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYyMTE1MDUzOFoXDTI0MDYyMDE1MDUzOFowUTELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkJ1Y2tleWUgVGVsZXN5c3RlbSBJbmMxITAfBgNVBAMMGEJ1Y2tleWUgU0hBS0VOIENlcnQgNzYwODBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJcZg3k8YdBoQhcnc0FlKNUuVMQ6B8xlJCR%2Fwh%2BRcGEClEx%2BADSYbZFbgMQbGhlfwoSRdPCv9xVkFRwzN3eCM3WjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc2MDgwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU4YkIR39UaK2zVovk0U1nN%2FFXlAkwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAMKzPT8OPYcsVwzs4QQwO25tQF5Au3rAJ7673LbvduhGAiEA5N6jmQq%2FLVmEN4ijgI31iIh%2FbK68KedW3Gq9iszCtRE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 06 Jul 23 14:08 UTC