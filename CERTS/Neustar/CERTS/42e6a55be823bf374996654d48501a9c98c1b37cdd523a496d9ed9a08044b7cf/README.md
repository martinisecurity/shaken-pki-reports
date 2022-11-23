# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Granite

Tested At: 23 Nov 22 18:06 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 542 day(s)\
Subject: C=US, ST=MA, L=Quincy, O=Granite Telecom, OU=Granite, CN=Granite\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11282.10179

[View certificate details](https://understandingwebpki.com/?cert=MIID9DCCAtygAwIBAgIUZ3adBkhZzYjmASzaR8i7uxEY0%2BEwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MTgxNTA3NTdaFw0yNDA1MTgxNTA3NTdaMGkxEDAOBgNVBAMMB0dyYW5pdGUxEDAOBgNVBAsMB0dyYW5pdGUxGDAWBgNVBAoMD0dyYW5pdGUgVGVsZWNvbTEPMA0GA1UEBwwGUXVpbmN5MQswCQYDVQQIDAJNQTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARqqQGioPs1kCJ1RX8yWBVr93ctdO5V4vXPMbL7Kj26A6X1jmTDV74QMLMe4qbXUJvuJLmB06Wu%2BujBUFnQ7cY9o4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBQ8QlTR9YgoF1qaJQmrilI%2FeUOSpDAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDgwNTIwDQYJKoZIhvcNAQELBQADggEBAAv3VuZ7GIWqVovRTp58gexwcVqB0CG1ewVi4ouUa0BYCOCWOladI0%2F9Wgobs3AnOc4%2BeFc33oEG6Mhps5lNKI7KfYGULSYJaaqLGBLdvwyyGCYBX87zhsgN7zkV66VK1lnqHILIkDfaCBkkxlQ4wUYY%2BvPLivT9H95s7Nwso3BZuBRmcdpyhPJT9I7a%2F%2BvRrMK5kmCaR%2FToUbsilRYslXABw3p%2BSWC8D0FcxWuuXKfbEm3flKzjYrG9fv%2B8bpqsJbsMRRK9AD21vW9mzFIvPeO26rw2cJWTRo7cIidLltL0BKXugUxYrDNKsP23%2BZ8KS7F3Z8JFhBgEH6UCV%2BuqSdc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 23 Nov 22 18:09 UTC