# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Project Mutual SHAKEN Cert 2231

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 253 day(s)\
Subject: CN=Project Mutual SHAKEN Cert 2231, O=Project Mutual, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/d4e575ce3348a47b80616aae48c240808772df2d

[View certificate details](https://x509.io/?cert=MIICVjCCAf2gAwIBAgIQTW1%2FNQhoxOBBOLrEb1A9%2BjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQyOTExMTE1M1oXDTI2MDQyODExMTE1M1owUDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlByb2plY3QgTXV0dWFsMSgwJgYDVQQDDB9Qcm9qZWN0IE11dHVhbCBTSEFLRU4gQ2VydCAyMjMxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJ%2BefblrwkP1gW8UAt15owlqGKTTy3fvGOYDMdJ7HPtETsajDEQ9HGVI7hAZ5CLK6WzrSZAZMrL%2BhYNUgn0QY16OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMjIzMTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTVWA%2BQtUQ8RWkQqv6U9uc%2F2yNDrjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBoc3eQzHDzEZqDiuGCoznVgp9ZHMU8qX6npHACUwuaZAIgX4Jt5Q3tuJVWihl%2FTDEDmNq%2BkJoFDEmbwSv9%2FTr3V0Y%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2231', but common name is 'Project Mutual SHAKEN Cert 2231' |


Generated: 18 Aug 25 21:13 UTC