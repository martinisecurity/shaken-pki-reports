# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Highland Telephone Cooperative SHAKEN Cert 0565

Tested At: 21 Aug 23 20:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 950 day(s)\
Subject: CN=Highland Telephone Cooperative SHAKEN Cert 0565, O=Highland Telephone Cooperative, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3e401c7a84f75f433a566157ca877a2a3143ba4d

[View certificate details](https://understandingwebpki.com/?cert=MIICdzCCAh2gAwIBAgIQepNNtqTB453JZJYwYT9B8zAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMyOTE4NDMzNFoXDTI2MDMyODE4NDMzNFowcDELMAkGA1UEBhMCVVMxJzAlBgNVBAoMHkhpZ2hsYW5kIFRlbGVwaG9uZSBDb29wZXJhdGl2ZTE4MDYGA1UEAwwvSGlnaGxhbmQgVGVsZXBob25lIENvb3BlcmF0aXZlIFNIQUtFTiBDZXJ0IDA1NjUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARmVv0eMxWl7P%2BtZd0iYVcFwvUtQ3fvYg3jtBuDcAbuHvkuk0l5ubDseXoEfXX467nB%2BbQxqra3I2XeQKLv9XyKo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNTY1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFBExB01F8ohnt0lBW%2B4F%2BSD%2BgkmbMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIH49ZcR7AHAYTEiG%2BbOf%2FTpfArkPBF22wA7UuDwoo6SIAiEA1aboFpliZP9n5eQaOMDtyvxSvhb88EsK15f12oRlIRo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0565' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 21 Aug 23 20:18 UTC