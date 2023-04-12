# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Canada Certified Caller ID SHAKEN CA-1

Tested At: 12 Apr 23 22:02 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 2737 day(s)\
Subject: CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN Root CA, OU=cms-ca.ccid.neustar, O=Neustar Information Services Inc, C=CA

[View certificate details](https://understandingwebpki.com/?cert=MIIDYzCCAwmgAwIBAgIUf6BMyINxdCzqcpPgxEDs9Zf%2Fth0wCgYIKoZIzj0EAwIwgZIxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTY21zLWNhLmNjaWQubmV1c3RhcjE6MDgGA1UEAwwxTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gUm9vdCBDQTAeFw0yMDEwMDgxNDUyMDZaFw0zMDEwMDkxNDUyMDZaMIGPMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxHDAaBgNVBAsME3d3dy5jYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIENBLTEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ59cw2g99pUFDICOymrr7k%2FM%2FRxFHQpvcBIlUj7OKxFDurd2C%2F1rpEXvySZVNGheIl88FkkTozNp90RQnhNCC5o4IBPDCCATgwDwYDVR0TAQH%2FBAUwAwEB%2FzAfBgNVHSMEGDAWgBRIA3%2FvHO2MLj%2BLxARyFhAD%2BtrZLjBdBggrBgEFBQcBAQRRME8wTQYIKwYBBQUHMAKGQWh0dHA6Ly9jYWNlcnRzLWNhLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRSb290Q0EuY3J0MBUGA1UdIAQOMAwwCgYIKwYBBAGDtx8wDwYDVR0lBAgwBgYEVR0lADBOBgNVHR8ERzBFMEOgQaA%2Fhj1odHRwOi8vY3JsLWNhLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRSb290Q0EuY3JsMB0GA1UdDgQWBBRT7lV9CanUzWgOsZmr0gA1fwhKuDAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSAAwRQIgad871RcpG74akDMq%2BW76nX8OzvhiKMDeAX1gKZlhzrQCIQC2yC6omJIuiZvQfqIhZZEtX8Z3ipDgim%2Bdeo98t2Av%2BA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_issuer_dn
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 22:02 UTC