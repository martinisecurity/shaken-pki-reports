# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Vexus Fiber SHAKEN Cert 4913

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 632 day(s)\
Subject: CN=Vexus Fiber SHAKEN Cert 4913, O=Vexus Fiber, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/97be1f553890f85d8a8bf06683588ce5061742f3

[View certificate details](https://x509.io/?cert=MIICUjCCAfegAwIBAgIQESygwDtsrbcSv4znfpMc%2FzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYyOTExMDQxNloXDTI2MDYyODExMDQxNlowSjELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1ZleHVzIEZpYmVyMSUwIwYDVQQDDBxWZXh1cyBGaWJlciBTSEFLRU4gQ2VydCA0OTEzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7r8%2FqawNeSkZ0r4IAu1gHZAJ7bgguS%2FEpQLGAYzrcW0jmmCc7rXCB8IwPFbqLSpf6o92G2k9x0Wozwi6AQ7L3qOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENDkxMzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSIoOvkLdhBpDH%2BYwXkfeo9rwev7TAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA4OjPDOnyvRUZP8LvcNqwOUlVXqWmhkmaQrAfX1JJM1gCIQDI6eapP15vxuBhRQccS0O%2FUE59x6XhhcHPmXvNbPhzXg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 4913', but common name is 'Vexus Fiber SHAKEN Cert 4913' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 04 Oct 24 16:29 UTC