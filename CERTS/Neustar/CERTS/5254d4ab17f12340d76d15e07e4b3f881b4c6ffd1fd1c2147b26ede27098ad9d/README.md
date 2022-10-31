# STIR/SHAKEN CA Ecosystem Compliance

## Certificate alticeusa.com

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 49 day(s)\
Subject: C=US, ST=NY, L=Hicksville, O=alticeusa, OU=alticeusa, CN=alticeusa.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.alticeusa.com/ccid/authn/v2/certs/11013.10003

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2BjCCAuKgAwIBAgIUCRPClb2wIhPjn1FZy%2BeFdmywPXwwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0xOTEyMTgyMDAxMjlaFw0yMjEyMTgyMDAxMjlaMG8xFjAUBgNVBAMMDWFsdGljZXVzYS5jb20xEjAQBgNVBAsMCWFsdGljZXVzYTESMBAGA1UECgwJYWx0aWNldXNhMRMwEQYDVQQHDApIaWNrc3ZpbGxlMQswCQYDVQQIDAJOWTELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2FDctE3GhH7PpUPVrYh5R7bG%2Bt51ULaxYYfAqrNE95EmDUhwJtVF3RfqETJb8aAG%2FyceUveHrR3qiVAdBgCToqo4IBSDCCAUQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBT2DWn6soookIKcaAsZCxFXaC7HLTAOBgNVHQ8BAf8EBAMCB4AwFgYIKwYBBQUHARoECjAIoAYWBDcxMjYwDQYJKoZIhvcNAQELBQADggEBAH5rPUnW%2FPSuWsXfRNSvqDZCBWxxUtiqOz7vI%2BjXhoaAONvabz79PiR%2FM%2BIRB7y2aCF9v67nPYoL1wsGnt8v3IRwc4K3x3MiLqnZRLucAgfscP8AsF0%2BljpeHsY0Cou6PjsdDIB9UNbsJJHxeoNRvVIlzrIlTYv%2FasAkyThzc62BxD6wy4flP74JHDAfxTfSYdi47ZWEI5twD7zENrJhiOtd5pQa9p%2F3T7F2YQPJPhWseKyxhGXzTV7vBWAtRoISUOKyuw820IZc2gLICj4P6ij9OdRsL1rJeI42NvSYEdtEJwfZ8niV57DLqjBlfRpr2OCZbU9snmjsJCw7R2eKrPM%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_authority_key_identifier
- e_sti_basic_constraints
- e_sti_certificate_policies
- e_sti_crl_distribution
- e_sti_crl_distribution_not_reachable
- e_sti_extension_unknown
- e_sti_issuer_dn
- e_sti_key_usage
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject
- e_sti_subject_cn
- e_sti_subject_key_identifier
- e_sti_subject_public_key
- e_sti_tn_auth_list
- e_sti_version
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22