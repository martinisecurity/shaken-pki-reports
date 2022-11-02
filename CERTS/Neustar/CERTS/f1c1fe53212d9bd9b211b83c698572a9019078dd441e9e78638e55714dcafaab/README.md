# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Inteliquent.com

Tested At: 02 Nov 22 17:23 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 292 day(s)\
Subject: C=US, ST=IL, L=Chicago, O=Inteliquent, OU=Inteliquent Voice, CN=Inteliquent.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/1.1

[View certificate details](https://understandingwebpki.com/?cert=MIIEAzCCAuugAwIBAgIUVkzUPVEKdxnbuRQR%2BVeCP3TuocYwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA4MjAwMDU0NDRaFw0yMzA4MjEwMDU0NDRaMHgxGDAWBgNVBAMMD0ludGVsaXF1ZW50LmNvbTEaMBgGA1UECwwRSW50ZWxpcXVlbnQgVm9pY2UxFDASBgNVBAoMC0ludGVsaXF1ZW50MRAwDgYDVQQHDAdDaGljYWdvMQswCQYDVQQIDAJJTDELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR1epK4YJ112hZGZeOVfqZ706xJHAgkNsAUXLt5hBSfzUZkLCu0o%2BfKkLiWymlls2O38mUUh%2FJVENmGrq%2Fpg4fCo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBTd5ZeLz%2FeRQwmdjV%2B37gOkm3g1aTAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDIyNEMwDQYJKoZIhvcNAQELBQADggEBADDynNUAZazgHdPkyZdIqHCzJd88K3wQczpVA0KyBZWYJdB3mWqQuAAofZesZEwm45EfWBKgv3TBe1jrWVse0ayoQbBTQr2Dwmi34sANexqjM0dAtJzi798bPgyJXrqj%2B6Zcgzudawlr0fK6CYIKxL%2BxlYr%2B7iFeUrF8YRtf%2B0xyBfC76NFfbnw0n6xSRodDyOue%2BBL39jj3guaH8R3lTORh94pVGHnKMdxV%2Br2oMMTtVp2KVbrfPAKoGigQ9KCXJOpdXbUHAo692AQGfM9hRWG5FrEDyb6F%2Bbk%2FxidWzXKvHlt9D%2BqxAcmcmwasWhFY2IGTvpOc%2FTXeNK10wOmyHX4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_authority_key_identifier
- e_atis_basic_constraints
- e_atis_certificate_policies
- e_atis_crl_distribution
- e_atis_extension_unknown
- e_atis_issuer_dn
- e_atis_key_usage
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject
- e_atis_subject_cn
- e_atis_subject_key_identifier
- e_atis_subject_public_key
- e_atis_tn_auth_list
- e_atis_version

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Nov 22 17:25 UTC