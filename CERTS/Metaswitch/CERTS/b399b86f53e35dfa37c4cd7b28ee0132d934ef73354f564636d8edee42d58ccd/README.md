# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 411 day(s)\
Subject: CN=Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H, O=Northeast Oklahoma Electric Cooperative\\, Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/766ee42919612faa4f0d943be5a055664928eb77

[View certificate details](https://x509.io/?cert=MIICjzCCAjWgAwIBAgIQXt3ENQ27HH3UKTOWikObsDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMTAxMjE3NTI1MFoXDTI1MTAxMTE3NTI1MFowgYcxCzAJBgNVBAYTAlVTMTUwMwYDVQQKDCxOb3J0aGVhc3QgT2tsYWhvbWEgRWxlY3RyaWMgQ29vcGVyYXRpdmUsIEluYzFBMD8GA1UEAww4Tm9ydGhlYXN0IE9rbGFob21hIEVsZWN0cmljIENvb3BlcmF0aXZlIFNIQUtFTiBDZXJ0IDk0NUgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQGOcF4BJppPhOty0%2B%2BlrbiTOGrtMuF2QIaTC%2F%2FtbABY61BTIxSf0yJuPYWs2Bg14xxZLJoqp%2BV6CcR5zF9r0evo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5NDVIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFOJcLKdo7aW%2Fd5o2teT6Vt%2BiyNr3MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIBUZhRl9qZCc0qUdqOgZeyT7GMEuMnoXwIOt6LvxwJfGAiEAocAxUU1Ty6Y%2FJu7gJgT2L%2B%2BU8gRlGV2advGqGkqatTI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 945H', but common name is 'Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC