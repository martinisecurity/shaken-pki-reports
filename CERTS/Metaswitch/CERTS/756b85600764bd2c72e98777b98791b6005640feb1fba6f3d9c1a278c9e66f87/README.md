# STIR/SHAKEN CA Ecosystem Compliance

## Certificate GCI SHAKEN Cert 7785

Tested At: 22 Aug 24 15:19 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 894 day(s)\
Subject: CN=GCI SHAKEN Cert 7785, O=GCI, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/f4b0a304bfb6f3db28f8eab0011a9c643d91849a

[View certificate details](https://x509.io/?cert=MIICQTCCAeegAwIBAgIQXrNiQYw3j7giQbwtXzydRzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIwMjE4MjE0OVoXDTI3MDIwMTE4MjE0OVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA0dDSTEdMBsGA1UEAwwUR0NJIFNIQUtFTiBDZXJ0IDc3ODUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASXI1LjZ8BtikaYO1FZLXstORSPsp9bW6D%2FXJXsVPfkFqg47jv3JKjokhOFKrUo4vNIFI64OniF9GRRAou8%2BxBoo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3Nzg1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFNj7Vg5aule3pG4iXJAm1LmxKMgcMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCtHHPYCg99Hy%2BzzdNk3DPj55kY3EIt8y2TvPguioKGogIgZBL4cwdII7Q1beyzDkgDWJHKD7vsML7FPmxXFQ78lIk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7785', but common name is 'GCI SHAKEN Cert 7785' |


Generated: 22 Aug 24 16:06 UTC