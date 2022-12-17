# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Viaero Wireless SHAKEN Cert 6874

Tested At: 17 Dec 22 17:03 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 517 day(s)\
Subject: CN=Viaero Wireless SHAKEN Cert 6874, O=Viaero Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/836dbe7b4ab9ad475219697d1553bed65e3d3463

[View certificate details](https://understandingwebpki.com/?cert=MIICWTCCAf%2BgAwIBAgIQCxBHoZcYCvB85F8RsDgV%2FzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUxODEzMjI1NFoXDTI0MDUxNzEzMjI1NFowUjELMAkGA1UEBhMCVVMxGDAWBgNVBAoMD1ZpYWVybyBXaXJlbGVzczEpMCcGA1UEAwwgVmlhZXJvIFdpcmVsZXNzIFNIQUtFTiBDZXJ0IDY4NzQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATBaR2A7GY4WG1gtAQyPLyWjHw4rMEAiTAVbyFnQqpSq1WM6jimOvc%2FHcEorS4iLAjG9dAUrjrntZU%2Bk960nc5Ao4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODc0MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFBIoYxrBvxS3FG6QBApcIK0i4fc5MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCsA4857W3gO4TYyyzI3Y8%2Fl1rhKxBTOisXHfQkO6vu0wIgK0ikyhvuZFZ6ZbbkmQG7wEvliUnG9XYAS0OqGtH1hyc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Dec 22 17:07 UTC