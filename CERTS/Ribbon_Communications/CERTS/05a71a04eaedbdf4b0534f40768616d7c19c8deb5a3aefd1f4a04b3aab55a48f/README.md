# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ribbon Issuing CA

Tested At: 08 Feb 23 19:45 UTC\
Initial Validity Period: 4383 day(s)\
Remaining Validity Period: 3747 day(s)\
Subject: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Issuer: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIICpTCCAkugAwIBAgIQW5Vnb2lq3Sln3c6nhnqzXjAKBggqhkjOPQQDAjBvMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSAwHgYDVQQLExdDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEeMBwGA1UEAxMVU0hBS0VOIFJpYmJvbiBSb290IENBMB4XDTIxMDUxMzAwMDAwMFoXDTMzMDUxMjIzNTk1OVowdDELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFVJpYmJvbiBDb21tdW5pY2F0aW9uczEiMCAGA1UECxMZQ2VydGlmaWNhdGlvbiBBdXRob3JpdGllczEhMB8GA1UEAxMYU0hBS0VOIFJpYmJvbiBJc3N1aW5nIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEW0X8zKK6TZCrJua%2FzicaGFtz9Vz3x3bXKHntBSKIsL3ZJOVAOMaKXTWeZalehDddsJCkieNDvg5vZMX2zftADaOBwzCBwDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH%2FBAUwAwEB%2FzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFI%2FflztFcnlC%2Bq8979xSNUBgjKTYMEQGCCsGAQUFBwEBBDgwNjA0BggrBgEFBQcwAoYoaHR0cDovL3N0aWNyLnJiYm5pZGh1Yi5jb20vcmJibl9yb290LmNydDAfBgNVHSMEGDAWgBS7bhaSFwYNkmN%2BwZQyNeVu0dUUZDAKBggqhkjOPQQDAgNIADBFAiEA8hUUx6fyh4cfzHOTuZPyslWA2TNf5t530%2BTPwqZNL9UCIArsu2c7gMwj29hOM%2BtzpDptbOI872LIhwhzVBToc81k)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 08 Feb 23 19:45 UTC