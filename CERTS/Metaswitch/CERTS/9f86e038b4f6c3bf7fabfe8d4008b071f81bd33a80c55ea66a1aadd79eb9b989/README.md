# STIR/SHAKEN CA Ecosystem Compliance

## Certificate WOW Internet Cable and Phone SHAKEN Cert 665E

Tested At: 23 Nov 22 18:08 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 549 day(s)\
Subject: CN=WOW Internet Cable and Phone SHAKEN Cert 665E, O=WOW Internet Cable and Phone, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/786c2c44c5fcfd1ed973607e3ac560ee3425d030

[View certificate details](https://understandingwebpki.com/?cert=MIICczCCAhmgAwIBAgIQEqy50c5Qk3H9FMJsgHYCKDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyNjE1Mzc0MVoXDTI0MDUyNTE1Mzc0MVowbDELMAkGA1UEBhMCVVMxJTAjBgNVBAoMHFdPVyBJbnRlcm5ldCBDYWJsZSBhbmQgUGhvbmUxNjA0BgNVBAMMLVdPVyBJbnRlcm5ldCBDYWJsZSBhbmQgUGhvbmUgU0hBS0VOIENlcnQgNjY1RTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJsCWJUT1NeLu%2B7jFplPYChdbOxoYA9WTHxYh%2BmCpUm1fUrRuM4rdjnvnPvGTzZMP%2BtIlw28hdskuw%2FqSLaDED6jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDY2NUUwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUz8RvWWhWv3epXqmwEdbU8PgS05gwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgOw2oREqPaQxbsfsn4kYls7IZxYoik3Z4prOhT8EBI4gCIQD%2B0A8dPQz4UdqSATiDPTafBm1EFSaHYWHcaBARqLgcjQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 23 Nov 22 18:09 UTC