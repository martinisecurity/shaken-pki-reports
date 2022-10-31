# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ribbon Issuing CA

Tested At: 31 Oct 22 18:34 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3847 day(s)\
Subject: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Issuer: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIICpTCCAkugAwIBAgIQW5Vnb2lq3Sln3c6nhnqzXjAKBggqhkjOPQQDAjBvMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSAwHgYDVQQLExdDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEeMBwGA1UEAxMVU0hBS0VOIFJpYmJvbiBSb290IENBMB4XDTIxMDUxMzAwMDAwMFoXDTMzMDUxMjIzNTk1OVowdDELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFVJpYmJvbiBDb21tdW5pY2F0aW9uczEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEhMB8GA1UEAxMYU0hBS0VOIFJpYmJvbiBJc3N1aW5nIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEW0X8zKK6TZCrJua%2FzicaGFtz9Vz3x3bXKHntBSKIsL3ZJOVAOMaKXTWeZalehDddsJCkieNDvg5vZMX2zftADaOBwzCBwDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH%2FBAUwAwEB%2FzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMEQGCCsGAQUFBwEBBDgwNjA0BggrBgEFBQcwAoYoaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9yb290LmNydDAfBgNVHSMEGDAWgBS7bhaSFwYNkmN%2BwZQyNeVu0dUUZDAKBggqhkjOPQQDAgNIADBFAiEA8hUUx6fyh4cfzHOTuZPyslWA2TNf5t530%2BTPwqZNL9UCIArsu2c7gMwj29hOM%2BtzpDptbOI872LIhwhzVBToc81k)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [n_pki_ca_key_usage](../../ISSUES/n_pki_ca_key_usage/README.md) | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| [w_pki_ca_subject_rdn_unknown](../../ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_certificate_policies
- e_sti_ca_crl_distribution
- e_sti_ca_extension_unknown
- e_sti_ca_serial_number
- e_sti_ca_subject_cn
- n_sti_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 18:34:12