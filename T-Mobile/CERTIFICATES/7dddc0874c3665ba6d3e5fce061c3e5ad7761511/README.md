# STIR/SHAKEN CA Ecosystem Compliance
## T-Mobile

### Certificate 7dddc0874c3665ba6d3e5fce061c3e5ad7761511
Tested At: 2022-10-27 22:13:25 +0000 UTC\
Initial Validity Period: 9131 day(s)\
Remaining Validity Period: 7997 day(s)\
Subject: CN=TMOBILE-PROD-ROOT-STIRSHAKEN-EC, O=TMOBILE-USA, C=US\
Issuer: CN=TMOBILE-PROD-ROOT-STIRSHAKEN-EC, O=TMOBILE-USA, C=US

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIICBjCCAaugAwIBAgICDh8wDAYIKoZIzj0EAwIFADBNMQswCQYDVQQGEwJVUzEUMBIGA1UEChMLVE1PQklMRS1VU0ExKDAmBgNVBAMTH1RNT0JJTEUtUFJPRC1ST09ULVNUSVJTSEFLRU4tRUMwHhcNMTkwOTE5MjAxMjAyWhcNNDQwOTE4MjAxMjAyWjBNMQswCQYDVQQGEwJVUzEUMBIGA1UEChMLVE1PQklMRS1VU0ExKDAmBgNVBAMTH1RNT0JJTEUtUFJPRC1ST09ULVNUSVJTSEFLRU4tRUMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASkBvB4%2Fv9K6e2gJESiPLFTW0HW9u1tJVaD%2BAF7SLkFUpHj89lqT5m%2Bv3PoBbt6NjEacsl6n7q37pGfbrA8Jzxbo3kwdzAfBgNVHSMEGDAWgBSD6W6bpOOt86MfODQwnn6hS7sJ1zAdBgNVHQ4EFgQUg%2Blum6TjrfOjHzg0MJ5%2BoUu7CdcwDgYDVR0PAQH%2FBAQDAgGmMBEGA1UdIAQKMAgwBgYEVR0gADASBgNVHRMBAf8ECDAGAQH%2FAgEBMAwGCCqGSM49BAMCBQADRwAwRAIgTimHZEmDJ1qY%2BNoOE0vWjbFGC%2FABB1vSkIFnAMu%2FxaUCIEF9yPw3mFC0TA4DiPtR4md516yZ82mLnW%2F%2BekdU1bGh)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

152 tests were ran and no warning or error level issues were found

### Not Effective

- e_sti_root_certificate_policies
- e_sti_ca_signature_algorithm
- e_sti_basic_constraints
- e_sti_root_authority_key_identifier
- e_sti_root_extension_unknown
- e_sti_ca_serial_number
- e_sti_ca_subject_key_identifier
- e_sti_ca_version
- e_sti_ca_key_usage
- e_sti_ca_subject
- e_sti_ca_subject_public_key
- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_subject_cn
- e_sti_ca_issuer_dn
- e_cp1_3_ca_key_usage_crl_sign

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:13:25