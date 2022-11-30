# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H

Tested At: 30 Nov 22 17:24 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 1047 day(s)\
Subject: CN=Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H, O=Northeast Oklahoma Electric Cooperative\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/235213bf0fd6f738e4ac1c4dfc9ccfd304428ace

[View certificate details](https://understandingwebpki.com/?cert=MIICjzCCAjWgAwIBAgIQXt3ENQ27HH3UKTOWikObsDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTAxMjE3NTI1MFoXDTI1MTAxMTE3NTI1MFowgYcxCzAJBgNVBAYTAlVTMTUwMwYDVQQKDCxOb3J0aGVhc3QgT2tsYWhvbWEgRWxlY3RyaWMgQ29vcGVyYXRpdmUsIEluYzFBMD8GA1UEAww4Tm9ydGhlYXN0IE9rbGFob21hIEVsZWN0cmljIENvb3BlcmF0aXZlIFNIQUtFTiBDZXJ0IDk0NUgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQGOcF4BJppPhOty0%2B%2BlrbiTOGrtMuF2QIaTC%2F%2FtbABY61BTIxSf0yJuPYWs2Bg14xxZLJoqp%2BV6CcR5zF9r0evo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5NDVIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFOJcLKdo7aW%2Fd5o2teT6Vt%2BiyNr3MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIBUZhRl9qZCc0qUdqOgZeyT7GMEuMnoXwIOt6LvxwJfGAiEAocAxUU1Ty6Y%2FJu7gJgT2L%2B%2BU8gRlGV2advGqGkqatTI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 945H' |


Generated: 30 Nov 22 17:24 UTC