# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate a754e630241fb966de4cac2cd9eb9db90c9d029e
Tested At: 2022-10-28 18:22:41 +0000 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 2903 day(s)\
Subject: CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN Root CA, OU=cms-ca.ccid.neustar, O=Neustar Information Services Inc, C=CA

Link: https://stir.fibernetics.ca/prod-cert2022.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDYzCCAwmgAwIBAgIUf6BMyINxdCzqcpPgxEDs9Zf%2Fth0wCgYIKoZIzj0EAwIwgZIxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTY21zLWNhLmNjaWQubmV1c3RhcjE6MDgGA1UEAwwxTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gUm9vdCBDQTAeFw0yMDEwMDgxNDUyMDZaFw0zMDEwMDkxNDUyMDZaMIGPMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxHDAaBgNVBAsME3d3dy5jYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIENBLTEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ59cw2g99pUFDICOymrr7k%2FM%2FRxFHQpvcBIlUj7OKxFDurd2C%2F1rpEXvySZVNGheIl88FkkTozNp90RQnhNCC5o4IBPDCCATgwDwYDVR0TAQH%2FBAUwAwEB%2FzAfBgNVHSMEGDAWgBRIA3%2FvHO2MLj%2BLxARyFhAD%2BtrZLjBdBggrBgEFBQcBAQRRME8wTQYIKwYBBQUHMAKGQWh0dHA6Ly9jYWNlcnRzLWNhLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRSb290Q0EuY3J0MBUGA1UdIAQOMAwwCgYIKwYBBAGDtx8wDwYDVR0lBAgwBgYEVR0lADBOBgNVHR8ERzBFMEOgQaA%2Fhj1odHRwOi8vY3JsLWNhLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRSb290Q0EuY3JsMB0GA1UdDgQWBBRT7lV9CanUzWgOsZmr0gA1fwhKuDAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSAAwRQIgad871RcpG74akDMq%2BW76nX8OzvhiKMDeAX1gKZlhzrQCIQC2yC6omJIuiZvQfqIhZZEtX8Z3ipDgim%2Bdeo98t2Av%2BA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| e_sti_ca_crl_distribution_not_reachable | error | ATIS-1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_sti_ca_subject_cn
- e_sti_ca_extension_unknown
- e_sti_ca_crl_distribution
- n_sti_ca_certificate_policy_critical
- e_sti_ca_serial_number
- e_sti_ca_issuer_dn
- e_sti_ca_certificate_policies
- e_cp1_3_ca_key_usage_crl_sign

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:22:55