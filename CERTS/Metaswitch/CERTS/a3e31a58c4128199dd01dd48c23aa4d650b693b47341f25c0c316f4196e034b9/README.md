# STIR/SHAKEN CA Ecosystem Compliance

## Certificate RCN SHAKEN Cert 7615

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 875 day(s)\
Subject: CN=RCN SHAKEN Cert 7615, O=RCN, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/a23201bf22dbf58314c267f8d2dc304ff278000d

[View certificate details](https://x509.io/?cert=MIICQjCCAeegAwIBAgIQI17to%2BebM%2BMcKINbuQAUyDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIyNzExMjYzMVoXDTI3MDIyNjExMjYzMVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA1JDTjEdMBsGA1UEAwwUUkNOIFNIQUtFTiBDZXJ0IDc2MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARRs97%2BOz2dQ93Gi449eDY93bZpXwKHeBC8FuUu8n9qGIPezOzb7GkMZ0p0ipxW2hgyEYN0T%2FaVp3WbEVEsir8Yo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFFwcIn%2B4wAGrKXmAkdkWRgJv0ug4MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQC9CHF0KQ8K7c%2BR7nutp8PxGK9Btfno9uOOTtFs%2BT7D9gIhAKK92CfrYII8L5oyO2jdP%2BwJzTEqIKf%2BItS5%2B17MC9l5)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7615', but common name is 'RCN SHAKEN Cert 7615' |


Generated: 04 Oct 24 16:29 UTC