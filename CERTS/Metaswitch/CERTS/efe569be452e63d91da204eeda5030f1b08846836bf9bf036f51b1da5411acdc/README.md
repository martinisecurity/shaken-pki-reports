# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nex-Tech Wireless SHAKEN Cert 122D

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 166 day(s)\
Subject: CN=Nex-Tech Wireless SHAKEN Cert 122D, O=Nex-Tech Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/211845364ec11f9349b1d056e2f7bdfc490fefe0

[View certificate details](https://understandingwebpki.com/?cert=MIICXDCCAgOgAwIBAgIQAx49OZyCBUvwgzslGTTuSzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQzMDE2NTkzN1oXDTI0MDQyOTE2NTkzN1owVjELMAkGA1UEBhMCVVMxGjAYBgNVBAoMEU5leC1UZWNoIFdpcmVsZXNzMSswKQYDVQQDDCJOZXgtVGVjaCBXaXJlbGVzcyBTSEFLRU4gQ2VydCAxMjJEMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEMTCalobTB6oS58qT%2FAJt6W0PbdUpmt76qymLdamNGhS%2BwC6pzvGCayrI17kcNPgdsn%2FXV8FFzALkWgOsQhjpBKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTIyRDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTujlGx6wthl129vHnZFXHx7rMxqjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiAxKu5qms%2Bdc5ig3rZZ1i7%2BchsfcRLb%2FsIoJCWLLg%2BgNwIgUkZ3tdgx3FGQ%2FuvFNDjzRPgNkhmwYTx1ijKwZmP1zPw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 15 Nov 23 17:17 UTC