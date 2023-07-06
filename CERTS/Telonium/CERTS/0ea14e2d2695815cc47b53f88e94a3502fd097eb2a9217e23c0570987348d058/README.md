# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 084E

Tested At: 06 Jul 23 13:57 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 359 day(s)\
Subject: serialNumber=a201c9f6-41db-4c41-947e-112fa91fa0f8, CN=SHAKEN 084E, O=ChoiceTEL, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://choicetel.net/084E.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC4jCCAomgAwIBAgIQZlTnoY02tMwk2L2HfmTeAzAKBggqhkjOPQQDAjBEMRgwFgYDVQQKEw9UZWxvbml1bSBTVEktQ0ExKDAmBgNVBAMTH1RlbG9uaXVtIFNUSS1DQSBJbnRlcm1lZGlhdGUgQ0EwHhcNMjMwNjI4MTUyNzUyWhcNMjQwNjI4MTUyODUyWjBmMQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2hvaWNlVEVMMRQwEgYDVQQDEwtTSEFLRU4gMDg0RTEtMCsGA1UEBRMkYTIwMWM5ZjYtNDFkYi00YzQxLTk0N2UtMTEyZmE5MWZhMGY4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE99F8BBHGsnSjEN27h2Qq6fSChbtFV7xHZYR2qwTysClshplQFPYbeiiP1z9ChDbljxeqD6j1Uls3dJXF0WeMc6OCATkwggE1MA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRd3B8xB4bpD0GxhTgFclMf34WcpTAfBgNVHSMEGDAWgBSqJLv%2FFHVAeS2Hb%2BgNQXfKu82IsDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBQGCCsGAQUFBwEaBAgwBhMEMDg0RTBdBgwrBgEEAYKkZMYoQAEETTBLAgEBBBlUZWxvbml1bSBTVEktQ0EgU1BDIFRva2VuBCtYVkNWYnNYNmlrd0psaXZCdnlZaGV4SmFiMjNiVHFmcjJhdGxkM2JVLXZFMAoGCCqGSM49BAMCA0cAMEQCIFl4ZG9c%2FDqZG8ZSvbxuJ%2BYgWLFjnW7hXHxUydgGU8%2BoAiA%2FO4xR6CL9SHdwG3i8TPg%2FLGBFHVM5EBvAQ1hp29x9kA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | TNAuthorizationList shall have only one TN Entry |


Generated: 06 Jul 23 14:08 UTC