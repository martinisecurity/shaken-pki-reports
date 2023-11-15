# STIR/SHAKEN CA Ecosystem Compliance

## Certificate 3 Rivers Communications SHAKEN Cert 2255

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 899 day(s)\
Subject: CN=3 Rivers Communications SHAKEN Cert 2255, O=3 Rivers Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3f3c33b7470f849766844b4a9a3ccc61f15c944a

[View certificate details](https://understandingwebpki.com/?cert=MIICaTCCAg%2BgAwIBAgIQBlPpeaB5ASZMLiDk2md8GzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwMzE2NTMxM1oXDTI2MDUwMjE2NTMxM1owYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMFzMgUml2ZXJzIENvbW11bmljYXRpb25zMTEwLwYDVQQDDCgzIFJpdmVycyBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAyMjU1MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhoXjsMhvC1QK7dqDUQy1dRjc15kPsiutX8Ndhexa4mI5h94TlTmvx1AudIGvzB2eKhH61JpUHOZjQxoheGp8tqOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjI1NTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTzTQHBEafzRtlyt1sNUvr3rI6rsDAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiAlPLZGxRU1ggcQwqzHKdSlL77VIVONaE8VM6zIdN7vYQIhAIWx0r32DEk7145%2B9nBKboNUw8U30gejfjSMLY06fOMr)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2255' |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


Generated: 15 Nov 23 17:17 UTC