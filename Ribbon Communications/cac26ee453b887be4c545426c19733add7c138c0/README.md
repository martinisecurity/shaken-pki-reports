# STIR/SHAKEN CA Ecosystem Compliance
## Ribbon Communications

### Certificate cac26ee453b887be4c545426c19733add7c138c0
Tested At: 2022-10-26 20:21:02 +0000 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3852 day(s)\
Subject: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Issuer: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US

Link: https://prod001-cr.rbbnidhub.com/Jx6yR-yMgz/620J202204-2c7d5c55a20834b031681dbd3e2eb9f0

View: [Click to view](https://understandingwebpki.com/?cert=MIICpTCCAkugAwIBAgIQW5Vnb2lq3Sln3c6nhnqzXjAKBggqhkjOPQQDAjBvMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSAwHgYDVQQLExdDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEeMBwGA1UEAxMVU0hBS0VOIFJpYmJvbiBSb290IENBMB4XDTIxMDUxMzAwMDAwMFoXDTMzMDUxMjIzNTk1OVowdDELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFVJpYmJvbiBDb21tdW5pY2F0aW9uczEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEhMB8GA1UEAxMYU0hBS0VOIFJpYmJvbiBJc3N1aW5nIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEW0X8zKK6TZCrJua%2FzicaGFtz9Vz3x3bXKHntBSKIsL3ZJOVAOMaKXTWeZalehDddsJCkieNDvg5vZMX2zftADaOBwzCBwDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH%2FBAUwAwEB%2FzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMEQGCCsGAQUFBwEBBDgwNjA0BggrBgEFBQcwAoYoaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9yb290LmNydDAfBgNVHSMEGDAWgBS7bhaSFwYNkmN%2BwZQyNeVu0dUUZDAKBggqhkjOPQQDAgNIADBFAiEA8hUUx6fyh4cfzHOTuZPyslWA2TNf5t530%2BTPwqZNL9UCIArsu2c7gMwj29hOM%2BtzpDptbOI872LIhwhzVBToc81k)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

150 tests were ran and no warning or error level issues were found

### Not Effective

- e_sti_ca_signature_algorithm
- e_sti_ca_subject
- e_sti_ca_extension_unknown
- e_sti_ca_subject_public_key
- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_issuer_dn
- e_sti_ca_subject_key_identifier
- e_sti_ca_serial_number
- e_sti_ca_key_usage
- e_sti_ca_certificate_policies
- e_sti_basic_constraints
- e_sti_ca_authority_key_identifier
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_version
- e_sti_ca_crl_distribution
- n_sti_ca_certificate_policy_critical
- e_sti_ca_subject_cn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 20:21:30