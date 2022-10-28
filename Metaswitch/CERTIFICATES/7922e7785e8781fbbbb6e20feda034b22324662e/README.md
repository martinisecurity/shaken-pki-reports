# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate 7922e7785e8781fbbbb6e20feda034b22324662e
Tested At: 2022-10-28 19:21:55 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 505 day(s)\
Subject: CN=GCI SHAKEN Cert 7785, O=GCI, C=us\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://sti-cr.cgah.tnsi.com/certs/cd1856717574765eb6b4bddb7a5bc8814e1e2103

View: [Click to view](https://understandingwebpki.com/?cert=MIICQjCCAeegAwIBAgIQWVIfIucXyRwdk2r1QxfxijAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MTI0MFoXDTI0MDMxNjE3MTI0MFowOjELMAkGA1UEBhMCdXMxDDAKBgNVBAoMA0dDSTEdMBsGA1UEAwwUR0NJIFNIQUtFTiBDZXJ0IDc3ODUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARaB2tkqnNVmOGuizyDsUBBLFhwWwoHf7WBK%2B539zfhV1dT5OQrbZJOgmiMy1cW7q%2BOAd1uuyOeag%2B2jcvLXmquo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3Nzg1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNBTKUsbL5mmfkdrpXdVRnf1NlU8MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQD1plr6Dl%2FKnrzzFBGkxc17Py0p0BToNjrqn8ReVd9RWQIhAK8BHwvgTKCJ8QAi2yCwJN%2Bc%2BWK4Jd2oKPkmxZuKlwHL)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_sti_subject_cn
- e_cp1_3_subject_sn
- e_sti_serial_number
- e_cp1_3_ambiguous_identifier
- w_cp_1_3_subject_email
- e_sti_extension_unknown
- e_sti_signature_algorithm

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 19:22:10