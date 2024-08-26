# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Verizon SHAKEN cert 5807

Tested At: 26 Aug 24 18:44 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 873 day(s)\
Subject: CN=Verizon SHAKEN cert 5807, OU=NNO CDS, O=Verizon Data Services LLC, L=Southlake, ST=Texas, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti.verizon.com/vzwcert/vzshaken-01-2027.crt

[View certificate details](https://x509.io/?cert=MIICkjCCAjigAwIBAgIQKhbfLnGA5Vw4ZBf8rw1fyjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDExNzExNTkxNFoXDTI3MDExNjExNTkxNFowgYoxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczESMBAGA1UEBwwJU291dGhsYWtlMSIwIAYDVQQKDBlWZXJpem9uIERhdGEgU2VydmljZXMgTExDMRAwDgYDVQQLDAdOTk8gQ0RTMSEwHwYDVQQDDBhWZXJpem9uIFNIQUtFTiBjZXJ0IDU4MDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASBDPoBKI2KeWsVaW9EPaArBrlbiS9%2BAhGAqCDTmkbt520L5IAtToF5rphsKp02yLvpq8Us8Ze9Zh4A%2FYCSZnAho4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODA3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFLzCDjxfuU3ObZVJBsf27qs8H5TlMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCEdIvTU2ofJfKMMP8wJxUOFapISWctkD155skaqJ8%2BrQIgJOD9UzQcVHlp9%2FiBf4le47xF0J9ViMvZWnUirikzdA0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5807', but common name is 'Verizon SHAKEN cert 5807' |


Generated: 26 Aug 24 18:49 UTC