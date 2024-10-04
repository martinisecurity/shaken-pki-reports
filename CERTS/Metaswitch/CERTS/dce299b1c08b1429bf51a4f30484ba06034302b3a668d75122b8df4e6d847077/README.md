# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TDS Telecom SHAKEN Cert 7804

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 793 day(s)\
Subject: CN=TDS Telecom SHAKEN Cert 7804, O=TDS Telecom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/35db85c681cd72f18e564492a27fd418bc58be28

[View certificate details](https://x509.io/?cert=MIICUDCCAfegAwIBAgIQQDJOIe8dTeVVDcVSKXVf8zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMTIwNjE3MzAwNVoXDTI2MTIwNTE3MzAwNVowSjELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1REUyBUZWxlY29tMSUwIwYDVQQDDBxURFMgVGVsZWNvbSBTSEFLRU4gQ2VydCA3ODA0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBmgzCXpPpKHS34pkC0H89BqbrxHxQJfnAl1XDrMBZrexGsmF53TTUOtm%2B8PQzJ9sPiz3tUKWybs8RYvxEWM6e6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENzgwNDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRcqC1m8YF54asOR03KoOc%2FaQSVVDAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiAjZ0YiqgOReyGvFJtBa3ZXzq2cJpazdPRimX7pwwit6wIgHQIyvl1%2BKsvwqt33Ko5I%2Fro5Ll6lvC4RN0tNCNZSqHk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7804', but common name is 'TDS Telecom SHAKEN Cert 7804' |


Generated: 04 Oct 24 16:29 UTC