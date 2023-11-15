# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Farmers Telecommunications Inc SHAKEN Cert 2188

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 908 day(s)\
Subject: CN=Farmers Telecommunications Inc SHAKEN Cert 2188, O=Farmers Telecommunications Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/7cd1ca9605d897bf552d7470db39657b5b2c843b

[View certificate details](https://understandingwebpki.com/?cert=MIICdjCCAh2gAwIBAgIQdYZRxTc2av4pKlz9jFZJ%2BjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUxMjExMTcwOFoXDTI2MDUxMTExMTcwOFowcDELMAkGA1UEBhMCVVMxJzAlBgNVBAoMHkZhcm1lcnMgVGVsZWNvbW11bmljYXRpb25zIEluYzE4MDYGA1UEAwwvRmFybWVycyBUZWxlY29tbXVuaWNhdGlvbnMgSW5jIFNIQUtFTiBDZXJ0IDIxODgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQRySeR7YrKT%2FYmlrj25UpQV3Fa3XIgbNzm%2B7Ga95eQPi9zQMG1CKAkutG4FEU97A%2FPJzMXAJJvC3Vu3BBi7ctto4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMTg4MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFCIygzlJnpYhxXIwE7KiU5bDxAK%2FMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIH0cMxEPrSIY8Qi%2BpeofTvo6pPbHNzyLB6qEs7iii2fkAiBe1MUGoG6DykBA3vGX5aDKSGF76ntp8sUmI25yI15YHw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2188' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 15 Nov 23 18:10 UTC