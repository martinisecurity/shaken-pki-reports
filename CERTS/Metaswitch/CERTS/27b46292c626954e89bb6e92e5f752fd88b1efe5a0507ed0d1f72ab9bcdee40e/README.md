# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Point Broadband Inc Bristol SHAKEN Cert 9809

Tested At: 21 Nov 23 18:47 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 887 day(s)\
Subject: CN=Point Broadband Inc Bristol SHAKEN Cert 9809, O=Point Broadband Inc Bristol, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/2aeba0ef7bc96ec8521d519c00854a8f850ddb3b

[View certificate details](https://understandingwebpki.com/?cert=MIICcjCCAhegAwIBAgIQEF9YgADD6qZhK3dI9V5N8DAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDQyNjIzNDQzMVoXDTI2MDQyNTIzNDQzMVowajELMAkGA1UEBhMCVVMxJDAiBgNVBAoMG1BvaW50IEJyb2FkYmFuZCBJbmMgQnJpc3RvbDE1MDMGA1UEAwwsUG9pbnQgQnJvYWRiYW5kIEluYyBCcmlzdG9sIFNIQUtFTiBDZXJ0IDk4MDkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR0%2Fr%2FMXmDuzSlqOZo0PoRhBZ6nlOKnq8Yz1PQOYNVE1YeTRArjMEHZMdRfJUlNKiXcqamYafjxaXu%2FvfOL8H%2Boo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5ODA5MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFAcitlzEfOq8Yj71Gt95vBbqo4jFMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQDf2SScGV%2Fo6hGB4seRmN3ILftQIk%2BEjcSKWP%2FT2lCkjgIhALQ%2BtmOuoRARqfr%2FnyHMRT%2BlqaZYpwP%2BAV49Dj90yqDr)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9809' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 21 Nov 23 19:18 UTC