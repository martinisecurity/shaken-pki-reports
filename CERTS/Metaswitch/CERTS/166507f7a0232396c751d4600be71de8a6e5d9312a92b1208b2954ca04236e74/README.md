# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Duo Broadband SHAKEN Cert 0401

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 124 day(s)\
Subject: CN=Duo Broadband SHAKEN Cert 0401, O=Duo Broadband, C=us\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/54a54551e139f777dcad69a463d282d8b6b9dade

[View certificate details](https://understandingwebpki.com/?cert=MIICVTCCAfugAwIBAgIQHuB6t83p1oJ9g7qgOH5fhjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMyNDE3NTcxMVoXDTI0MDMyMzE3NTcxMVowTjELMAkGA1UEBhMCdXMxFjAUBgNVBAoMDUR1byBCcm9hZGJhbmQxJzAlBgNVBAMMHkR1byBCcm9hZGJhbmQgU0hBS0VOIENlcnQgMDQwMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABO%2FolkJrJhsHFUMs%2Ft2YSMHuQHDOWKj2KRxdJN6LoWVs2cpvrrmjlD28fr6MSQhzAlhvTYQTT2neKbs0jxSKo2mjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA0MDEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUcy1hpia45zj6pmShTvMO1cvXuHIwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgUQqHT9gIc%2B%2BQ9dDVchpPUx%2BqT7iFOfofo1fZV7E3qPoCIQDGf0VHW0QwT43BVabSsVIblCFcf%2BmFUyQ%2FaWoY8spSbw%3D%3D)

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


Generated: 21 Nov 23 01:55 UTC