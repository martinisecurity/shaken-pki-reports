# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 3201

Tested At: 28 Nov 23 10:45 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 176 day(s)\
Subject: C=US, ST=PR, L=Caparra, O=PRTC, OU=VOIP, CN=SHAKEN 3201\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://prod001-cr.rbbnidhub.com/9JhfutuGgz/0

[View certificate details](https://understandingwebpki.com/?cert=MIID6zCCAtOgAwIBAgIUUibL4HkveCE%2BPqEKsPfH%2FhkG5qwwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MjIwNzM1NDFaFw0yNDA1MjIwNzM1NDFaMGAxFDASBgNVBAMMC1NIQUtFTiAzMjAxMQ0wCwYDVQQLDARWT0lQMQ0wCwYDVQQKDARQUlRDMRAwDgYDVQQHDAdDYXBhcnJhMQswCQYDVQQIDAJQUjELMAkGA1UEBhMCVVMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT6VHDf0WOZ6kpfgC6CxB2KMV%2Fa3DOyqHt1BXBG9YvYOuNse74vZnMs2lOgvcK9%2BPek6wM%2BR5sp2XvWks6sZwgno4IBSDCCAUQwFgYIKwYBBQUHARoECjAIoAYWBDMyMDEwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQ7uVzLMRecSM%2FU2mbrjbao1eEJ9zCBgQYIKwYBBQUHAQEEdTBzMEcGCCsGAQUFBzAChjtodHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydDAoBggrBgEFBQcwAYYcaHR0cDovL29jc3AtY2ExLmNjaWQubmV1c3RhcjBIBgNVHR8EQTA%2FMD2gO6A5hjdodHRwOi8vY3JsLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3JsMB0GA1UdDgQWBBTrBvTrc%2BCrYO2bVWwXIlU5dk8GPzAOBgNVHQ8BAf8EBAMCB4AwDQYJKoZIhvcNAQELBQADggEBAIMlMNA1QLXNwhIQszWdARzbyK2h4nB3qF9IW7nmJInCnvitvQ3bPDVQ7j6iTp9xv2fx2JI7kkmV6ADotcGjVlgkbooB1ej8Y%2Fx3w7vhX30IFLFIblRC6HqVvs0uDj46DzkaStf60GNL1MoP%2BewEo4fIVKafBvL%2Fmr%2F0dcgKPvTTCirsHgq7%2BSjqUOIgeEZsnuT8E6z6UNu6EXRbYhWdWcafIDILhui6XkUvHOvUlsycP8%2F%2FTZ%2BINLAx5%2Bf9JV7otk7o41bCNm8ygNuWAFSRVEW52ITqZ46MWju7mitiGnA1gdetNjTfO%2FrMguIeJFct9RxZR7Mo%2BISh2PuOkYyg8Cw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC