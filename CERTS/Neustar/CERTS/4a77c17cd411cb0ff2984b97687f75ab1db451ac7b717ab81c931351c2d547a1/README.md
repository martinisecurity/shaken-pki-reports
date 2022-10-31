# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID Root CA

Tested At: 31 Oct 22 16:43 UTC\
Initial Validity Period: 7305 day(s)\
Remaining Validity Period: 6171 day(s)\
Subject: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID Root CA\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID Root CA

View: [Click to view](https://understandingwebpki.com/?cert=MIIEBjCCAu6gAwIBAgIUfzd00EOROqoBaE%2FZjx%2BKO8oDcZgwDQYJKoZIhvcNAQELBQAwgYExLDAqBgNVBAMMI05ldXN0YXIgQ2VydGlmaWVkIENhbGxlciBJRCBSb290IENBMRkwFwYDVQQLDBB3d3cuY2NpZC5uZXVzdGFyMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzELMAkGA1UEBhMCVVMwHhcNMTkwOTIzMTMyNjAzWhcNMzkwOTIzMTMyNjAzWjCBgTEsMCoGA1UEAwwjTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFJvb3QgQ0ExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJJ21OlKIeUlZhhiKKHDyM4ExD88e%2Bogs2dysFNV0xnyQdVLC4%2BrrjJgLKFMbyMR%2FK6IFhYorEjAU2YCJGeOtFK6iYovJMqfJzfZ9qFrHn61KTg7N1yCmFcIdddr6GJiclPwwQAIGoMG4N%2F3oSFGTihfek%2FmrUoHsTNp8%2BQ2%2Fv97mSTVHvw%2FWWXjXvUkY4epWxSYebEwmJrmEBmCT%2FK2ZJn6rhzZD5cgQJiq4G6tfWx6qoXCiaSuGgSKxdhCBFFQd3dKWVgD%2FT6DqOq1keUwfLVAZzJeDRLXXq09YPbHKaRrBN2yOp8fZUNUonK7yUViSA8m%2B7v3T5PUPcDH%2BL%2BnA6sCAwEAAaN0MHIwDwYDVR0TAQH%2FBAUwAwEB%2FzAfBgNVHSMEGDAWgBQDxFFTZk7gLnzaaVrRHhiUhp4JGzAPBgNVHSUECDAGBgRVHSUAMB0GA1UdDgQWBBQDxFFTZk7gLnzaaVrRHhiUhp4JGzAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQELBQADggEBAEIHuI04hnW1Eur3sqc2FRbYcCZC4uhcGEGD0tnUyiy7IksPPumEfuwfwUJGV5g73SIWBTFjkzaE3UQzfoIBVrtcG0oWCVSpPnvrEyiDkVAMbQ31lsjQ3Pi7pXse6jOCDJ91NY1rcJhJ0lNfHR2vzHAPw7J8opa05t6B%2FcKTWMVBr6M0xOmdkKtpPNsLnkQc7Q36HrsU2o7q6crRJg8LgUyi7vB16KDTAjlISRcoJDFAAqZ6%2FPXeXDOwzinmOPQn8iCS9JUPWdd1%2F%2BZAfz%2B7clFz%2Bhje90PXj66Zg520KyTdqchC09SnJCwvB0YGsUz4yOdbn86pC1BRVJZSLGHgjAQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_ca_subject_rdn_unknown](../../ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [n_pki_ca_key_usage](../../ISSUES/n_pki_ca_key_usage/README.md) | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |


### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- e_sti_basic_constraints
- e_sti_ca_issuer_dn
- e_sti_ca_key_usage
- e_sti_ca_serial_number
- e_sti_ca_signature_algorithm
- e_sti_ca_subject
- e_sti_ca_subject_cn
- e_sti_ca_subject_key_identifier
- e_sti_ca_subject_public_key
- e_sti_ca_version
- e_sti_root_authority_key_identifier
- e_sti_root_certificate_policies
- e_sti_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22