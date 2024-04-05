# STIR/SHAKEN CA Ecosystem Compliance

## Certificate 3 Rivers Communications SHAKEN Cert 2255

Tested At: 05 Apr 24 18:39 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 757 day(s)\
Subject: CN=3 Rivers Communications SHAKEN Cert 2255, O=3 Rivers Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3f3c33b7470f849766844b4a9a3ccc61f15c944a

[View certificate details](https://x509.io/?cert=MIICaTCCAg%2BgAwIBAgIQBlPpeaB5ASZMLiDk2md8GzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwMzE2NTMxM1oXDTI2MDUwMjE2NTMxM1owYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMFzMgUml2ZXJzIENvbW11bmljYXRpb25zMTEwLwYDVQQDDCgzIFJpdmVycyBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAyMjU1MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhoXjsMhvC1QK7dqDUQy1dRjc15kPsiutX8Ndhexa4mI5h94TlTmvx1AudIGvzB2eKhH61JpUHOZjQxoheGp8tqOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjI1NTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTzTQHBEafzRtlyt1sNUvr3rI6rsDAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiAlPLZGxRU1ggcQwqzHKdSlL77VIVONaE8VM6zIdN7vYQIhAIWx0r32DEk7145%2B9nBKboNUw8U30gejfjSMLY06fOMr)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2255', but common name is '3 Rivers Communications SHAKEN Cert 2255' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 05 Apr 24 19:04 UTC