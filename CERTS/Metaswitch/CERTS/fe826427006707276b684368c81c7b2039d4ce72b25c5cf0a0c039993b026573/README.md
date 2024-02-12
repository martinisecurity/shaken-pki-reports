# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Viaero Wireless SHAKEN Cert 6874

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 95 day(s)\
Subject: CN=Viaero Wireless SHAKEN Cert 6874, O=Viaero Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/5da2f338f4df53a4a7212bae4465f45063b263b1

[View certificate details](https://understandingwebpki.com/?cert=MIICWTCCAf%2BgAwIBAgIQCxBHoZcYCvB85F8RsDgV%2FzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUxODEzMjI1NFoXDTI0MDUxNzEzMjI1NFowUjELMAkGA1UEBhMCVVMxGDAWBgNVBAoMD1ZpYWVybyBXaXJlbGVzczEpMCcGA1UEAwwgVmlhZXJvIFdpcmVsZXNzIFNIQUtFTiBDZXJ0IDY4NzQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATBaR2A7GY4WG1gtAQyPLyWjHw4rMEAiTAVbyFnQqpSq1WM6jimOvc%2FHcEorS4iLAjG9dAUrjrntZU%2Bk960nc5Ao4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODc0MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFBIoYxrBvxS3FG6QBApcIK0i4fc5MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCsA4857W3gO4TYyyzI3Y8%2Fl1rhKxBTOisXHfQkO6vu0wIgK0ikyhvuZFZ6ZbbkmQG7wEvliUnG9XYAS0OqGtH1hyc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC