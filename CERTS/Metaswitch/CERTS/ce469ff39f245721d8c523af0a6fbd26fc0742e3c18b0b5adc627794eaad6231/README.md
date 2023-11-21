# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ben Lomand SHAKEN Cert 0553

Tested At: 21 Nov 23 17:37 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 941 day(s)\
Subject: CN=Ben Lomand SHAKEN Cert 0553, O=Ben Lomand Rural Telephone COOP Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/fe74899c830827a37cd701d031e0a7714e0604ab

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg6gAwIBAgIQTeK6A19cuHcISx6gxUZDzDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYyMDA5MjkwNVoXDTI2MDYxOTA5MjkwNVowYTELMAkGA1UEBhMCVVMxLDAqBgNVBAoMI0JlbiBMb21hbmQgUnVyYWwgVGVsZXBob25lIENPT1AgSW5jMSQwIgYDVQQDDBtCZW4gTG9tYW5kIFNIQUtFTiBDZXJ0IDA1NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARJ3V8Ubfp6gxw1uOSSQLrollan4Qmw9e6k8b5ufCsBU%2F4QDlMUX4RdZ3HW6cJuQnp2kNreP8GFQGLSSPbzAoE4o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNTUzMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFM6a5Ihl1%2FpoxDQOa4d9pjsHDKNWMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQDFxr9NuFVILVoPOhuGmWOBCmC9oIGGwOniGUb0y9IFGgIgUsZ08nFpQQLQKp8vYW1ka7Uy8z1iXisQJnpFDEiE%2F8A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0553' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 21 Nov 23 17:53 UTC