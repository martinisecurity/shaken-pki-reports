# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Farmers Telecommunications Inc SHAKEN Cert 2188

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 584 day(s)\
Subject: CN=Farmers Telecommunications Inc SHAKEN Cert 2188, O=Farmers Telecommunications Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/7cd1ca9605d897bf552d7470db39657b5b2c843b

[View certificate details](https://x509.io/?cert=MIICdjCCAh2gAwIBAgIQdYZRxTc2av4pKlz9jFZJ%2BjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUxMjExMTcwOFoXDTI2MDUxMTExMTcwOFowcDELMAkGA1UEBhMCVVMxJzAlBgNVBAoMHkZhcm1lcnMgVGVsZWNvbW11bmljYXRpb25zIEluYzE4MDYGA1UEAwwvRmFybWVycyBUZWxlY29tbXVuaWNhdGlvbnMgSW5jIFNIQUtFTiBDZXJ0IDIxODgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQRySeR7YrKT%2FYmlrj25UpQV3Fa3XIgbNzm%2B7Ga95eQPi9zQMG1CKAkutG4FEU97A%2FPJzMXAJJvC3Vu3BBi7ctto4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMTg4MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFCIygzlJnpYhxXIwE7KiU5bDxAK%2FMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIH0cMxEPrSIY8Qi%2BpeofTvo6pPbHNzyLB6qEs7iii2fkAiBe1MUGoG6DykBA3vGX5aDKSGF76ntp8sUmI25yI15YHw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2188', but common name is 'Farmers Telecommunications Inc SHAKEN Cert 2188' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC