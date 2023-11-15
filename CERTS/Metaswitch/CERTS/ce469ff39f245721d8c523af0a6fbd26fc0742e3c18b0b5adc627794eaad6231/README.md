# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ben Lomand SHAKEN Cert 0553

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 947 day(s)\
Subject: CN=Ben Lomand SHAKEN Cert 0553, O=Ben Lomand Rural Telephone COOP Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/fe74899c830827a37cd701d031e0a7714e0604ab

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg6gAwIBAgIQTeK6A19cuHcISx6gxUZDzDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYyMDA5MjkwNVoXDTI2MDYxOTA5MjkwNVowYTELMAkGA1UEBhMCVVMxLDAqBgNVBAoMI0JlbiBMb21hbmQgUnVyYWwgVGVsZXBob25lIENPT1AgSW5jMSQwIgYDVQQDDBtCZW4gTG9tYW5kIFNIQUtFTiBDZXJ0IDA1NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARJ3V8Ubfp6gxw1uOSSQLrollan4Qmw9e6k8b5ufCsBU%2F4QDlMUX4RdZ3HW6cJuQnp2kNreP8GFQGLSSPbzAoE4o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNTUzMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFM6a5Ihl1%2FpoxDQOa4d9pjsHDKNWMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDFxr9NuFVILVoPOhuGmWOBCmC9oIGGwOniGUb0y9IFGgIgUsZ08nFpQQLQKp8vYW1ka7Uy8z1iXisQJnpFDEiE%2F8A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0553' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 15 Nov 23 16:51 UTC