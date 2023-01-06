# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Tri-County Telephone Association Inc SHAKEN Cert 2296

Tested At: 06 Jan 23 02:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 564 day(s)\
Subject: CN=Tri-County Telephone Association Inc SHAKEN Cert 2296, O=Tri-County Telephone West Association Inc, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/260bd48b7dcfa43beee9f59287f1e4d87fe3fb5f

[View certificate details](https://understandingwebpki.com/?cert=MIICijCCAi%2BgAwIBAgIQSrit0NzBWRoXEEfryyoqpjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDcyMzE2NTkyN1oXDTI0MDcyMjE2NTkyN1owgYExCzAJBgNVBAYTAlVTMTIwMAYDVQQKDClUcmktQ291bnR5IFRlbGVwaG9uZSBXZXN0IEFzc29jaWF0aW9uIEluYzE%2BMDwGA1UEAww1VHJpLUNvdW50eSBUZWxlcGhvbmUgQXNzb2NpYXRpb24gSW5jIFNIQUtFTiBDZXJ0IDIyOTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASUm0U2109MMqKBAf9d%2FnVVMV7ilqm%2FANuto%2BVHnXSIgi6%2BTV5l3iv7Ah%2FrstAxAynox%2B3PgMBRbjocPMUAXzSPo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMjk2MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFJJ942wRQ5ojUia47HPj0ONukYObMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQCHnVuF%2BsfOlAcHWTuGv8SBRDd8k2GcI0SNgN0ZZZi63QIhAMEAf%2FZR%2F%2BfTUktypy8SykbrmnuU6zmxIq8Ia6FsKjw2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 06 Jan 23 03:03 UTC